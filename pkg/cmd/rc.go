package cmd

import (
	"github.com/appscode/go-term"
	otx "github.com/appscode/osm/pkg/context"
	"github.com/spf13/cobra"
)

type bucketRemoveRequest struct {
	context string
	bucket  string
	force   bool
}

func NewCmdRemoveContainer() *cobra.Command {
	req := &bucketRemoveRequest{}
	cmd := &cobra.Command{
		Use:     "rc <name>",
		Short:   "Remove container",
		Example: "osm rc mybucket",
		Run: func(cmd *cobra.Command, args []string) {
			req.bucket = args[0]
			removeContainer(req)
		},
	}

	cmd.Flags().StringVarP(&req.context, "context", "", "", "Name of osmconfig context to use")
	cmd.Flags().BoolVarP(&req.force, "force", "f", false, "Force delete any files inside the container")
	return cmd
}

func removeContainer(req *bucketRemoveRequest) {
	cfg, err := otx.LoadConfig()
	term.ExitOnError(err)

	loc, err := cfg.Dial(req.context)
	term.ExitOnError(err)

	if req.force {
		c, err := loc.Container(req.bucket)
		term.ExitOnError(err)

		cursor := ""
		for {
			items, next, err := c.Items("", cursor, 50)
			term.ExitOnError(err)
			for _, item := range items {
				c.RemoveItem(item.ID())
			}
			if len(items) == 0 {
				break
			}
			cursor = next
		}
	}

	err = loc.RemoveContainer(req.bucket)
	term.ExitOnError(err)
	term.Successln("Successfully removed container " + req.bucket)
}

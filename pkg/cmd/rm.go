package cmd

import (
	"github.com/appscode/go-term"
	otx "github.com/appscode/osm/pkg/context"
	"github.com/spf13/cobra"
)

type itemRemoveRequest struct {
	context   string
	container string
	itemID    string
}

func NewCmdRemove() *cobra.Command {
	req := &itemRemoveRequest{}
	cmd := &cobra.Command{
		Use:     "rm <id>",
		Short:   "Remove item from container",
		Example: "osm rm -c mybucket f1.txt",
		Run: func(cmd *cobra.Command, args []string) {
			req.itemID = args[0]
			removeItem(req)
		},
	}

	cmd.Flags().StringVar(&req.context, "context", "", "Name of osmconfig context to use")
	cmd.Flags().StringVarP(&req.container, "container", "c", "", "Name of container")
	return cmd
}

func removeItem(req *itemRemoveRequest) {
	cfg, err := otx.LoadConfig()
	term.ExitOnError(err)

	loc, err := cfg.Dial(req.context)
	term.ExitOnError(err)

	c, err := loc.Container(req.container)
	term.ExitOnError(err)

	err = c.RemoveItem(req.itemID)
	term.ExitOnError(err)
	term.Successln("Successfully removed item " + req.itemID)
}

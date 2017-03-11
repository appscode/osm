package cmd

import (
	"github.com/appscode/go-term"
	otx "github.com/appscode/osm/pkg/context"
	"github.com/spf13/cobra"
)

type containerCreateRequest struct {
	context string
	bucket  string
}

func NewCmdMakeContainer() *cobra.Command {
	req := &containerCreateRequest{}
	cmd := &cobra.Command{
		Use:     "mc <name>",
		Short:   "Create container",
		Example: "osm mc mybucket",
		Run: func(cmd *cobra.Command, args []string) {
			req.bucket = args[0]
			createContainer(req)
		},
	}

	cmd.Flags().StringVarP(&req.context, "context", "", "", "Name of osmconfig context to use")
	return cmd
}

func createContainer(req *containerCreateRequest) {
	cfg, err := otx.LoadConfig()
	term.ExitOnError(err)

	loc, err := cfg.Dial(req.context)
	term.ExitOnError(err)

	_, err = loc.CreateContainer(req.bucket)
	term.ExitOnError(err)
	term.Successln("Successfully created container " + req.bucket)
}

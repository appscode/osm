package cmd

import (
	"github.com/appscode/go-term"
	"github.com/appscode/go/io"
	otx "github.com/appscode/osm/pkg/context"
	"github.com/spf13/cobra"
)

type itemPullRequest struct {
	context   string
	container string
	srcID     string
	destPath  string
}

func NewCmdPull() *cobra.Command {
	req := &itemPullRequest{}
	cmd := &cobra.Command{
		Use:     "pull <src> <dest>",
		Short:   "Pull item from container",
		Example: "osm pull -c mybucket f1.txt /tmp/f1.txt",
		Run: func(cmd *cobra.Command, args []string) {
			req.srcID = args[0]
			req.destPath = args[1]
			pullItem(req)
		},
	}

	cmd.Flags().StringVar(&req.context, "context", "", "Name of osmconfig context to use")
	cmd.Flags().StringVarP(&req.container, "container", "c", "", "Name of container")
	return cmd
}

func pullItem(req *itemPullRequest) {
	cfg, err := otx.LoadConfig()
	term.ExitOnError(err)

	loc, err := cfg.Dial(req.context)
	term.ExitOnError(err)

	c, err := loc.Container(req.container)
	term.ExitOnError(err)

	item, err := c.Item(req.srcID)
	term.ExitOnError(err)

	rd, err := item.Open()
	term.ExitOnError(err)
	defer rd.Close()

	err = io.WriteFile(req.destPath, rd, 0640)
	term.ExitOnError(err)
	term.Successln("Successfully pulled item " + req.srcID)
}

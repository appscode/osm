package cmd

import (
	"os"

	"github.com/appscode/go-term"
	otx "github.com/appscode/osm/pkg/context"
	"github.com/spf13/cobra"
)

type itemPushRequest struct {
	context   string
	container string
	srcPath   string
	destID    string
}

func NewCmdPush() *cobra.Command {
	req := &itemPushRequest{}
	cmd := &cobra.Command{
		Use:     "push <src> <dest>",
		Short:   "Push item from container",
		Example: "osm push -c mybucket f1.txt /tmp/f1.txt",
		Run: func(cmd *cobra.Command, args []string) {
			req.srcPath = args[0]
			req.destID = args[1]
			pushItem(req)
		},
	}

	cmd.Flags().StringVar(&req.context, "context", "", "Name of osmconfig context to use")
	cmd.Flags().StringVarP(&req.container, "container", "c", "", "Name of container")
	return cmd
}

func pushItem(req *itemPushRequest) {
	cfg, err := otx.LoadConfig()
	term.ExitOnError(err)

	loc, err := cfg.Dial(req.context)
	term.ExitOnError(err)

	c, err := loc.Container(req.container)
	term.ExitOnError(err)

	in, err := os.Open(req.srcPath)
	if err != nil {
		return
	}
	defer in.Close()

	fi, err := in.Stat()
	term.ExitOnError(err)

	item, err := c.Put(req.destID, in, fi.Size(), nil)
	term.ExitOnError(err)
	term.Successln("Successfully pushed item " + item.ID())
}

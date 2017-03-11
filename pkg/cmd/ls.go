package cmd

import (
	"fmt"

	"github.com/appscode/go-term"
	otx "github.com/appscode/osm/pkg/context"
	"github.com/graymeta/stow"
	"github.com/spf13/cobra"
)

type containerListRequest struct {
	context   string
	container string
	prefix    string
	delimiter string
}

func NewCmdListContainer() *cobra.Command {
	req := &containerListRequest{}
	cmd := &cobra.Command{
		Use:     "ls <name>",
		Short:   "List container",
		Example: "osm ls mybucket",
		Run: func(cmd *cobra.Command, args []string) {
			req.container = args[0]
			listContainer(req)
		},
	}

	cmd.Flags().StringVarP(&req.context, "context", "", "", "Name of osmconfig context to use")
	cmd.Flags().StringVarP(&req.prefix, "prefix", "", stow.NoPrefix, "Prefix of container")
	cmd.Flags().StringVarP(&req.delimiter, "delimiter", "", "", "Delimiter for path (optional)")
	return cmd
}

func listContainer(req *containerListRequest) {
	cfg, err := otx.LoadConfig()
	term.ExitOnError(err)

	loc, err := cfg.Dial(req.context)
	term.ExitOnError(err)

	c, err := loc.Container(req.container)
	term.ExitOnError(err)

	cursor := stow.CursorStart
	n := 0
	for {
		_, items, next, err := c.Browse(req.prefix, req.delimiter, cursor, 50)
		term.ExitOnError(err)
		for _, item := range items {
			n++
			term.Infoln(item.ID())
		}
		if len(items) == 0 {
			break
		}
		cursor = next
	}
	term.Successln(fmt.Sprintf("Found %v items in container %v", n, req.container))
}
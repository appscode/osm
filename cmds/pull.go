package cmds

import (
	"os"
	"path/filepath"

	"github.com/appscode/go-term"
	"github.com/appscode/go/io"
	otx "github.com/appscode/osm/context"
	"github.com/graymeta/stow"
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
		Use:               "pull <src> <dest>",
		Short:             "Pull item from container",
		Example:           "osm pull -c mybucket f1.txt /tmp/f1.txt",
		DisableAutoGenTag: true,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 2 {
				term.Errorln("Provide source item and destination path as argument. See examples:")
				cmd.Help()
				os.Exit(1)
			} else if len(args) > 2 {
				cmd.Help()
				os.Exit(1)
			}

			req.srcID = args[0]
			req.destPath = args[1]
			pull(req, otx.GetConfigPath(cmd))
		},
	}

	cmd.Flags().StringVar(&req.context, "context", "", "Name of osmconfig context to use")
	cmd.Flags().StringVarP(&req.container, "container", "c", "", "Name of container")
	return cmd
}

func pull(req *itemPullRequest, configPath string) {
	cfg, err := otx.LoadConfig(configPath)
	term.ExitOnError(err)

	loc, err := cfg.Dial(req.context)
	term.ExitOnError(err)

	c, err := loc.Container(req.container)
	term.ExitOnError(err)

	item, err := c.Item(req.srcID)
	if err != nil {
		cursor := stow.CursorStart
		for {
			items, next, err := c.Items(req.srcID, cursor, 50)
			term.ExitOnError(err)
			for _, item := range items {
				r, err := filepath.Rel(req.srcID, item.ID())
				term.ExitOnError(err)

				f := filepath.Join(req.destPath, r)
				os.MkdirAll(filepath.Dir(f), 0755)
				pullItem(item, f, item.ID())
			}
			cursor = next
			if stow.IsCursorEnd(cursor) {
				break
			}
		}
		term.Successln("Successfully pulled folder " + req.srcID)
	} else {
		pullItem(item, req.destPath, req.srcID)
	}
}

func pullItem(item stow.Item, destPath, srcID string) {
	rd, err := item.Open()
	term.ExitOnError(err)
	defer rd.Close()

	err = io.WriteFile(destPath, rd, 0640)
	term.ExitOnError(err)
	term.Successln("Successfully pulled item " + srcID)
}

/*
Copyright The osm Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmds

import (
	"io"
	"os"
	"path/filepath"

	otx "github.com/appscode/osm/context"
	"github.com/spf13/cobra"
	"gomodules.xyz/stow"
	"gomodules.xyz/stow/local"
	"gomodules.xyz/x/term"
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
				_ = cmd.Help()
				os.Exit(1)
			} else if len(args) > 2 {
				_ = cmd.Help()
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

	ctx, err := cfg.Context(req.context)
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
				src := req.srcID
				if ctx.Provider == local.Kind {
					src = filepath.Join(c.ID(), src)
				}
				r, err := filepath.Rel(src, item.ID())
				term.ExitOnError(err)

				f := filepath.Join(req.destPath, r)
				err = os.MkdirAll(filepath.Dir(f), 0755)
				term.ExitOnError(err)
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

func pullItem(item stow.Item, dstPath, srcID string) {
	rd, err := item.Open()
	term.ExitOnError(err)
	defer rd.Close()

	out, err := os.OpenFile(dstPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0640)
	if err != nil {
		return
	}
	defer func() {
		if e := out.Close(); e != nil {
			err = e
		}
	}()
	_, err = io.Copy(out, rd)
	term.ExitOnError(err)
	term.Successln("Successfully pulled item " + srcID)
}

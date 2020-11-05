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
	"os"
	"path/filepath"

	otx "github.com/appscode/osm/context"
	"github.com/spf13/cobra"
	"gomodules.xyz/stow"
	"gomodules.xyz/x/term"
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
		Use:               "push <src> <dest>",
		Short:             "Push item to container",
		Example:           "osm push -c mybucket f1.txt /tmp/f1.txt",
		DisableAutoGenTag: true,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 2 {
				term.Errorln("Provide source path and destination item as argument. See examples:")
				_ = cmd.Help()
				os.Exit(1)
			} else if len(args) > 2 {
				_ = cmd.Help()
				os.Exit(1)
			}

			req.srcPath = args[0]
			req.destID = args[1]
			push(req, otx.GetConfigPath(cmd))
		},
	}

	cmd.Flags().StringVar(&req.context, "context", "", "Name of osmconfig context to use")
	cmd.Flags().StringVarP(&req.container, "container", "c", "", "Name of container")
	return cmd
}

func push(req *itemPushRequest, configPath string) {
	cfg, err := otx.LoadConfig(configPath)
	term.ExitOnError(err)

	loc, err := cfg.Dial(req.context)
	term.ExitOnError(err)

	c, err := loc.Container(req.container)
	term.ExitOnError(err)

	si, err := os.Stat(req.srcPath)
	term.ExitOnError(err)
	if si.IsDir() {
		err := filepath.Walk(req.srcPath, func(path string, fi os.FileInfo, err error) error {
			if !fi.IsDir() {
				r, err := filepath.Rel(req.srcPath, path)
				term.ExitOnError(err)
				pushItem(c, filepath.Join(req.destID, r), path, fi)
			}
			return nil
		})
		term.ExitOnError(err)
		term.Successln("Successfully pushed folder " + req.srcPath)
	} else {
		pushItem(c, req.destID, req.srcPath, si)
	}
}

func pushItem(c stow.Container, destID, srcPath string, fi os.FileInfo) {
	in, err := os.Open(srcPath)
	if err != nil {
		return
	}
	defer in.Close()

	item, err := c.Put(destID, in, fi.Size(), nil)
	term.ExitOnError(err)
	term.Successln("Successfully pushed item " + item.ID())
}

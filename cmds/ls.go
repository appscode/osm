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
	"fmt"
	"os"

	"github.com/appscode/go/term"
	otx "github.com/appscode/osm/context"
	"github.com/spf13/cobra"
	"gomodules.xyz/stow"
)

type itemListRequest struct {
	context   string
	container string
	prefix    string
	delimiter string
}

func NewCmdListIetms() *cobra.Command {
	req := &itemListRequest{}
	cmd := &cobra.Command{
		Use:               "ls <name>",
		Short:             "List items in a container",
		Example:           "osm ls mybucket",
		DisableAutoGenTag: true,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				term.Errorln("Provide container name as argument. See examples:")
				_ = cmd.Help()
				os.Exit(1)
			} else if len(args) > 1 {
				_ = cmd.Help()
				os.Exit(1)
			}

			req.container = args[0]
			listItems(req, otx.GetConfigPath(cmd))
		},
	}

	cmd.Flags().StringVarP(&req.context, "context", "", "", "Name of osmconfig context to use")
	cmd.Flags().StringVarP(&req.prefix, "prefix", "", stow.NoPrefix, "Prefix of container")
	cmd.Flags().StringVarP(&req.delimiter, "delimiter", "", "", "Delimiter for path (optional)")
	return cmd
}

func listItems(req *itemListRequest, configPath string) {
	cfg, err := otx.LoadConfig(configPath)
	term.ExitOnError(err)

	loc, err := cfg.Dial(req.context)
	term.ExitOnError(err)

	c, err := loc.Container(req.container)
	term.ExitOnError(err)

	cursor := stow.CursorStart
	n := 0
	for {
		page, err := c.Browse(req.prefix, req.delimiter, cursor, 50)
		term.ExitOnError(err)
		for _, item := range page.Items {
			n++
			term.Infoln(item.ID())
		}
		cursor = page.Cursor
		if stow.IsCursorEnd(cursor) {
			break
		}
	}
	cnt := fmt.Sprintf("%v items", n)
	if n <= 1 {
		cnt = fmt.Sprintf("%v item", n)
	}
	term.Successln(fmt.Sprintf("Found %v in container %v", cnt, req.container))
}

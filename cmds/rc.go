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

	"github.com/appscode/go/term"
	otx "github.com/appscode/osm/context"
	"github.com/spf13/cobra"
	"gomodules.xyz/stow"
)

type containerRemoveRequest struct {
	context   string
	container string
	force     bool
}

func NewCmdRemoveContainer() *cobra.Command {
	req := &containerRemoveRequest{}
	cmd := &cobra.Command{
		Use:               "rc <name>",
		Short:             "Remove container",
		Example:           "osm rc mybucket",
		DisableAutoGenTag: true,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				term.Errorln("Provide container name as argument. See examples:")
				cmd.Help()
				os.Exit(1)
			} else if len(args) > 1 {
				cmd.Help()
				os.Exit(1)
			}

			req.container = args[0]
			removeContainer(req, otx.GetConfigPath(cmd))
		},
	}

	cmd.Flags().StringVarP(&req.context, "context", "", "", "Name of osmconfig context to use")
	cmd.Flags().BoolVarP(&req.force, "force", "f", false, "Force delete any files inside the container")
	return cmd
}

func removeContainer(req *containerRemoveRequest, configPath string) {
	cfg, err := otx.LoadConfig(configPath)
	term.ExitOnError(err)

	loc, err := cfg.Dial(req.context)
	term.ExitOnError(err)

	if req.force {
		c, err := loc.Container(req.container)
		term.ExitOnError(err)

		cursor := stow.CursorStart
		for {
			items, next, err := c.Items(stow.NoPrefix, cursor, 50)
			term.ExitOnError(err)
			for _, item := range items {
				term.Warningln("Removing item: " + item.ID())
				c.RemoveItem(item.ID())
			}
			cursor = next
			if stow.IsCursorEnd(cursor) {
				break
			}
		}
	}

	err = loc.RemoveContainer(req.container)
	term.ExitOnError(err)
	term.Successln("Successfully removed container " + req.container)
}

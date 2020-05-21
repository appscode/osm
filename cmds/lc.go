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

type containerListRequest struct {
	context string
}

func NewCmdListContainers() *cobra.Command {
	req := &containerListRequest{}
	cmd := &cobra.Command{
		Use:               "lc",
		Short:             "List containers",
		Example:           "osm lc",
		DisableAutoGenTag: true,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 1 {
				_ = cmd.Help()
				os.Exit(1)
			}

			listContainers(req, otx.GetConfigPath(cmd))
		},
	}

	cmd.Flags().StringVarP(&req.context, "context", "", "", "Name of osmconfig context to use")
	return cmd
}

func listContainers(req *containerListRequest, configPath string) {
	cfg, err := otx.LoadConfig(configPath)
	term.ExitOnError(err)

	loc, err := cfg.Dial(req.context)
	term.ExitOnError(err)

	cursor := stow.CursorStart
	n := 0
	for {
		containers, next, err := loc.Containers(stow.NoPrefix, cursor, 10)
		term.ExitOnError(err)
		for _, c := range containers {
			n++
			term.Infoln(c.ID())
		}
		cursor = next
		if stow.IsCursorEnd(cursor) {
			break
		}
	}
	cnt := fmt.Sprintf("%v containers", n)
	if n <= 1 {
		cnt = fmt.Sprintf("%v container", n)
	}
	term.Successln(fmt.Sprintf("Found %v in %v", cnt, req.context))
}

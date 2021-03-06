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

	otx "github.com/appscode/osm/context"
	"github.com/spf13/cobra"
	"gomodules.xyz/x/term"
)

type containerMakeRequest struct {
	context   string
	container string
}

func NewCmdMakeContainer() *cobra.Command {
	req := &containerMakeRequest{}
	cmd := &cobra.Command{
		Use:               "mc <name>",
		Short:             "Make container",
		Example:           "osm mc mybucket",
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
			makeContainer(req, otx.GetConfigPath(cmd))
		},
	}

	cmd.Flags().StringVarP(&req.context, "context", "", "", "Name of osmconfig context to use")
	return cmd
}

func makeContainer(req *containerMakeRequest, configPath string) {
	cfg, err := otx.LoadConfig(configPath)
	term.ExitOnError(err)

	loc, err := cfg.Dial(req.context)
	term.ExitOnError(err)

	_, err = loc.Container(req.container)
	if err != nil {
		_, err = loc.CreateContainer(req.container)
		term.ExitOnError(err)
		term.Successln("Successfully created container " + req.container)
	} else {
		term.Infoln("Container " + req.container + " already exists!")
	}
}

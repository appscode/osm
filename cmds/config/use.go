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

package config

import (
	"os"

	otx "github.com/appscode/osm/context"
	"github.com/spf13/cobra"
	"gomodules.xyz/x/term"
)

func newCmdUse() *cobra.Command {
	setCmd := &cobra.Command{
		Use:               "use-context <name>",
		Short:             "Use context",
		Example:           "osm config use-context <name>",
		DisableAutoGenTag: true,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				term.Errorln("Provide context name as argument. See examples:")
				_ = cmd.Help()
				os.Exit(1)
			} else if len(args) > 1 {
				_ = cmd.Help()
				os.Exit(1)
			}

			name := args[0]
			useContex(name, otx.GetConfigPath(cmd))
		},
	}
	return setCmd
}

func useContex(name, configPath string) {
	config, err := otx.LoadConfig(configPath)
	term.ExitOnError(err)

	if config.CurrentContext == name {
		return
	}

	found := false
	for i := range config.Contexts {
		if config.Contexts[i].Name == name {
			found = true
			break
		}
	}
	if !found {
		term.Fatalln("Invalid context name")
	}

	config.CurrentContext = name
	err = config.Save(configPath)
	term.ExitOnError(err)
}

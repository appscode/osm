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

	"github.com/appscode/go/term"
	otx "github.com/appscode/osm/context"
	"github.com/ghodss/yaml"
	"github.com/spf13/cobra"
)

func newCmdView() *cobra.Command {
	setCmd := &cobra.Command{
		Use:               "view",
		Short:             "Print osm config",
		Example:           "osm config view",
		DisableAutoGenTag: true,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 {
				cmd.Help()
				os.Exit(1)
			}
			viewContext(otx.GetConfigPath(cmd))
		},
	}
	return setCmd
}

func viewContext(configPath string) {
	config, err := otx.LoadConfig(configPath)
	term.ExitOnError(err)

	data, err := yaml.Marshal(config)
	term.ExitOnError(err)

	term.Infoln(string(data))
}

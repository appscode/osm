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

import "github.com/spf13/cobra"

func NewCmdConfig() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "config",
		Short:             "OSM configuration",
		Example:           "osm config view",
		DisableAutoGenTag: true,
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}

	cmd.AddCommand(newCmdSet())
	cmd.AddCommand(newCmdUse())
	cmd.AddCommand(newCmdView())
	cmd.AddCommand(newCmdGet())
	cmd.AddCommand(newCmdCurrent())
	return cmd
}

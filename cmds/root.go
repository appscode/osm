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
	"flag"
	"path/filepath"

	v "github.com/appscode/go/version"
	cfgCmd "github.com/appscode/osm/cmds/config"
	"github.com/spf13/cobra"
	"k8s.io/client-go/util/homedir"
	"kmodules.xyz/client-go/tools/cli"
)

func NewCmdOsm() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:               "osm [command]",
		Short:             `Object Store Manipulator by AppsCode`,
		DisableAutoGenTag: true,
		PersistentPreRun: func(c *cobra.Command, args []string) {
			cli.SendAnalytics(c, v.Version.Version)
		},
		Run: func(c *cobra.Command, args []string) {
			_ = c.Help()
		},
	}
	rootCmd.PersistentFlags().AddGoFlagSet(flag.CommandLine)
	rootCmd.PersistentFlags().String("osmconfig", filepath.Join(homedir.HomeDir(), ".osm", "config"), "Path to osm config")
	rootCmd.PersistentFlags().BoolVar(&cli.EnableAnalytics, "enable-analytics", cli.EnableAnalytics, "Send usage events to Google Analytics")

	rootCmd.PersistentFlags().BoolVar(&cli.EnableAnalytics, "analytics", cli.EnableAnalytics, "Send usage events to Google Analytics")
	_ = rootCmd.PersistentFlags().MarkDeprecated("analytics", "use --enable-analytics")

	rootCmd.AddCommand(cfgCmd.NewCmdConfig())

	rootCmd.AddCommand(NewCmdListContainers())
	rootCmd.AddCommand(NewCmdMakeContainer())
	rootCmd.AddCommand(NewCmdRemoveContainer())

	rootCmd.AddCommand(NewCmdListIetms())
	rootCmd.AddCommand(NewCmdPush())
	rootCmd.AddCommand(NewCmdPull())
	rootCmd.AddCommand(NewCmdStat())
	rootCmd.AddCommand(NewCmdRemove())

	rootCmd.AddCommand(v.NewCmdVersion())

	return rootCmd
}

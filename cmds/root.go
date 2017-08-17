package cmds

import (
	"flag"

	v "github.com/appscode/go/version"
	cfgCmd "github.com/appscode/osm/cmds/config"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

func NewCmdOsm() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:               "osm [command]",
		Short:             `Object Store Manipulator by AppsCode`,
		DisableAutoGenTag: true,
		Run: func(c *cobra.Command, args []string) {
			c.Help()
		},
	}
	rootCmd.PersistentFlags().AddGoFlagSet(flag.CommandLine)
	home, _ := homedir.Dir()
	rootCmd.PersistentFlags().String("osmconfig", home+"/.osm/config", "Path to osm config")

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

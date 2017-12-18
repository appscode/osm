package cmds

import (
	"flag"
	"strings"

	v "github.com/appscode/go/version"
	"github.com/appscode/kutil/tools/analytics"
	cfgCmd "github.com/appscode/osm/cmds/config"
	"github.com/jpillora/go-ogle-analytics"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

const (
	gaTrackingCode = "UA-62096468-20"
)

func NewCmdOsm() *cobra.Command {
	var (
		enableAnalytics = true
	)
	rootCmd := &cobra.Command{
		Use:               "osm [command]",
		Short:             `Object Store Manipulator by AppsCode`,
		DisableAutoGenTag: true,
		PersistentPreRun: func(c *cobra.Command, args []string) {
			if enableAnalytics && gaTrackingCode != "" {
				if client, err := ga.NewClient(gaTrackingCode); err == nil {
					client.ClientID(analytics.ClientID())
					parts := strings.Split(c.CommandPath(), " ")
					client.Send(ga.NewEvent(parts[0], strings.Join(parts[1:], "/")).Label(v.Version.Version))
				}
			}
		},
		Run: func(c *cobra.Command, args []string) {
			c.Help()
		},
	}
	rootCmd.PersistentFlags().AddGoFlagSet(flag.CommandLine)
	home, _ := homedir.Dir()
	rootCmd.PersistentFlags().String("osmconfig", home+"/.osm/config", "Path to osm config")
	rootCmd.PersistentFlags().BoolVar(&enableAnalytics, "analytics", enableAnalytics, "Send analytical events to Google Analytics")

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

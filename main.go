package main

import (
	"flag"
	"log"

	v "github.com/appscode/go/version"
	osmCmd "github.com/appscode/osm/pkg/cmd"
	cfgCmd "github.com/appscode/osm/pkg/cmd/config"
	"github.com/spf13/cobra"
)

var (
	Version         string
	VersionStrategy string
	Os              string
	Arch            string
	CommitHash      string
	GitBranch       string
	GitTag          string
	CommitTimestamp string
	BuildTimestamp  string
	BuildHost       string
	BuildHostOs     string
	BuildHostArch   string
)

func init() {
	v.Version.Version = Version
	v.Version.VersionStrategy = VersionStrategy
	v.Version.Os = Os
	v.Version.Arch = Arch
	v.Version.CommitHash = CommitHash
	v.Version.GitBranch = GitBranch
	v.Version.GitTag = GitTag
	v.Version.CommitTimestamp = CommitTimestamp
	v.Version.BuildTimestamp = BuildTimestamp
	v.Version.BuildHost = BuildHost
	v.Version.BuildHostOs = BuildHostOs
	v.Version.BuildHostArch = BuildHostArch
}

func main() {
	rootCmd := &cobra.Command{
		Use:   "osm [command]",
		Short: `AppsCode Object Store Manipulator`,
		Run: func(c *cobra.Command, args []string) {
			c.Help()
		},
	}
	rootCmd.PersistentFlags().AddGoFlagSet(flag.CommandLine)

	rootCmd.AddCommand(cfgCmd.NewCmdConfig())

	rootCmd.AddCommand(osmCmd.NewCmdMakeContainer())
	rootCmd.AddCommand(osmCmd.NewCmdListContainer())
	rootCmd.AddCommand(osmCmd.NewCmdRemoveContainer())

	rootCmd.AddCommand(osmCmd.NewCmdPush())
	rootCmd.AddCommand(osmCmd.NewCmdPull())
	rootCmd.AddCommand(osmCmd.NewCmdStat())
	rootCmd.AddCommand(osmCmd.NewCmdRemove())

	rootCmd.AddCommand(v.NewCmdVersion())
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

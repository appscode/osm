package cmd

import (
	"github.com/appscode/osm/pkg/cmd/config"
	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command {
	c := &cobra.Command{
		Use:     "osm",
		Short:   "AppsCode Object Store Manipulator",
		Example: "osm version",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	// add all sub commands to root command.
	c.AddCommand(config.NewCmdConfig())
	c.AddCommand(newCmdCreate())
	c.AddCommand(newCmdVersion())
	return c
}

package cmd

import (
	"github.com/appscode/osm/pkg/printer"
	"github.com/spf13/cobra"
)

var logo = `
 ▄██████▄     ▄████████   ▄▄▄▄███▄▄▄▄
███    ███   ███    ███ ▄██▀▀▀███▀▀▀██▄
███    ███   ███    █▀  ███   ███   ███
███    ███   ███        ███   ███   ███
███    ███ ▀███████████ ███   ███   ███
███    ███          ███ ███   ███   ███
███    ███    ▄█    ███ ███   ███   ███
 ▀██████▀   ▄████████▀   ▀█   ███   █▀

`

func newCmdVersion() *cobra.Command {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Print osm version information",
		Run: func(cmd *cobra.Command, args []string) {
			printer.Version(logo)
			printer.Info("Version: 0.1")
		},
	}
	return versionCmd
}

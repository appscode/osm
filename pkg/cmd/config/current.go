package config

import (
	"os"

	"github.com/appscode/go-term"
	otx "github.com/appscode/osm/pkg/context"
	"github.com/spf13/cobra"
)

func newCmdCurrent() *cobra.Command {
	setCmd := &cobra.Command{
		Use:     "current-context",
		Short:   "Print current context",
		Example: "osm config current-context",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 {
				cmd.Help()
				os.Exit(1)
			}
			currentContext()
		},
	}
	return setCmd
}

func currentContext() {
	config, err := otx.LoadConfig()
	term.ExitOnError(err)

	term.Infoln(config.CurrentContext)
}

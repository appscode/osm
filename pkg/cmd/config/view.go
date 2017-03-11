package config

import (
	"os"

	"github.com/appscode/go-term"
	otx "github.com/appscode/osm/pkg/context"
	"github.com/ghodss/yaml"
	"github.com/spf13/cobra"
)

func newCmdView() *cobra.Command {
	setCmd := &cobra.Command{
		Use:     "view",
		Short:   "View config",
		Example: "osm config view",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 1 {
				cmd.Help()
				os.Exit(1)
			}
			viewContex()
		},
	}
	return setCmd
}

func viewContex() {
	config, err := otx.LoadConfig()
	term.ExitOnError(err)

	data, err := yaml.Marshal(config)
	term.ExitOnError(err)

	term.Infoln(string(data))
}

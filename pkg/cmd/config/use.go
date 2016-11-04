package config

import (
	"os"

	osm_context "github.com/appscode/osm/pkg/context"
	"github.com/appscode/osm/pkg/printer"
	"github.com/spf13/cobra"
)

func newCmdUse() *cobra.Command {
	setCmd := &cobra.Command{
		Use:     "use-context",
		Short:   "Use context",
		Example: "osm config use-context",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				printer.Error("Provide context name as argument. See examples")
				cmd.Help()
				os.Exit(1)
			} else if len(args) > 1 {
				cmd.Help()
				os.Exit(1)
			}

			name := args[0]
			useContex(name)
		},
	}
	return setCmd
}

func useContex(name string) {
	config, err := osm_context.GetConfigData()
	if err != nil {
		printer.Error(err)
		return
	}
	var contextData *osm_context.Context
	for _, osmCtx := range config.Contexts {
		if osmCtx.Name == name {
			contextData = osmCtx
		}
	}

	if contextData != nil {
		config.CurrentContext = contextData.Name
	} else {
		printer.Error("Invalid context name")

	}

	if err := osm_context.SetConfigData(config); err != nil {
		printer.Error(err)
		return
	}
}

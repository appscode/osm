package config

import (
	osm_context "github.com/appscode/osm/pkg/context"
	"github.com/appscode/osm/pkg/printer"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

func newCmdView() *cobra.Command {
	viewCmd := &cobra.Command{
		Use:     "view",
		Short:   "View OSM configuration file",
		Example: "osm config view",
		Run: func(cmd *cobra.Command, args []string) {
			viewConfig()
		},
	}
	return viewCmd
}

func viewConfig() {
	data, err := osm_context.GetConfigData()
	if err != nil {
		printer.Error(err)
		return
	}
	d, err := yaml.Marshal(&data)
	if err != nil {
		printer.Error(err)
		return
	}
	printer.Config(string(d))
}

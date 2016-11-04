package config

import (
	"os"

	osm_context "github.com/appscode/osm/pkg/context"
	"github.com/appscode/osm/pkg/printer"
	"github.com/spf13/cobra"
)

type setContextRequest struct {
	CredentialDir string
	Provider      string
	Name          string
}

func newCmdSet() *cobra.Command {
	req := &setContextRequest{}
	setCmd := &cobra.Command{
		Use:     "set-context",
		Short:   "Set context",
		Example: "osm config set-context",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				printer.Error("Provide context name as argument. See examples")
				cmd.Help()
				os.Exit(1)
			} else if len(args) > 1 {
				cmd.Help()
				os.Exit(1)
			}

			req.Name = args[0]
			setContex(req)
		},
	}
	setCmd.Flags().StringVarP(&req.Provider, "provider", "p", "", "The name of the provider")
	setCmd.Flags().StringVarP(&req.CredentialDir, "credential-dir", "c", "", "The directory of credential file")
	return setCmd
}

func setContex(req *setContextRequest) {
	config, err := osm_context.GetConfigData()
	if err != nil {
		printer.Error(err)
		os.Exit(0)
	}
	var contextData *osm_context.Context
	if config != nil {
		for _, osmCtx := range config.Contexts {
			if osmCtx.Name == req.Name {
				contextData = osmCtx
			}
		}
	} else {
		config = &osm_context.ConfigData{
			Contexts: make([]*osm_context.Context, 0),
		}
	}

	if contextData != nil {
		if req.Provider != "" {
			contextData.ContextData.Provider = req.Provider
		}
		if req.CredentialDir != "" {
			contextData.ContextData.CredentialDir = req.CredentialDir
		}
	} else {
		if req.Name != "" {
			ctx := &osm_context.Context{
				Name: req.Name,
			}
			ctx.ContextData.Provider = req.Provider
			ctx.ContextData.CredentialDir = req.CredentialDir

			config.Contexts = append(config.Contexts, ctx)
		}
	}

	if err := osm_context.SetConfigData(config); err != nil {
		printer.Error(err)
		os.Exit(0)
	}
}

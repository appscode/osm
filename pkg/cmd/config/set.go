package config

import (
	"io/ioutil"
	"os"

	"github.com/appscode/go-term"
	otx "github.com/appscode/osm/pkg/context"
	"github.com/graymeta/stow"
	"github.com/graymeta/stow/azure"
	gcs "github.com/graymeta/stow/google"
	"github.com/graymeta/stow/local"
	"github.com/graymeta/stow/s3"
	"github.com/graymeta/stow/swift"
	"github.com/spf13/cobra"
)

type setContextRequest struct {
	Name                     string
	Provider                 string
	s3ConfigAccessKeyID      string
	s3ConfigEndpoint         string
	s3ConfigRegion           string
	s3ConfigSecretKey        string
	gcsConfigJSONKeyPath     string
	gcsConfigProjectId       string
	gcsConfigScopes          string
	azureConfigAccount       string
	azureConfigKey           string
	localConfigKeyPath       string
	swiftConfigKey           string
	swiftConfigTenantAuthURL string
	swiftConfigTenantName    string
	swiftConfigUsername      string
}

func newCmdSet() *cobra.Command {
	req := &setContextRequest{}
	setCmd := &cobra.Command{
		Use:     "set-context <name>",
		Short:   "Set context",
		Example: "osm config set-context <name>",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				term.Errorln("Provide context name as argument. See examples:")
				cmd.Help()
				os.Exit(1)
			} else if len(args) > 1 {
				cmd.Help()
				os.Exit(1)
			}

			req.Name = args[0]
			setContext(req)
		},
	}
	setCmd.Flags().StringVar(&req.Provider, "provider", "", "Cloud storage provider")

	setCmd.Flags().StringVar(&req.s3ConfigAccessKeyID, s3.Kind+"."+s3.ConfigAccessKeyID, "", "S3 config access key id")
	setCmd.Flags().StringVar(&req.s3ConfigEndpoint, s3.Kind+"."+s3.ConfigEndpoint, "", "S3 config endpoint")
	setCmd.Flags().StringVar(&req.s3ConfigRegion, s3.Kind+"."+s3.ConfigRegion, "", "S3 config region")
	setCmd.Flags().StringVar(&req.s3ConfigSecretKey, s3.Kind+"."+s3.ConfigSecretKey, "", "S3 config secret key")

	setCmd.Flags().StringVar(&req.gcsConfigJSONKeyPath, gcs.Kind+".json_key_path", "", "GCS config json key path")
	setCmd.Flags().StringVar(&req.gcsConfigProjectId, gcs.Kind+"."+gcs.ConfigProjectId, "", "GCS config project id")
	setCmd.Flags().StringVar(&req.gcsConfigScopes, gcs.Kind+"."+gcs.ConfigScopes, "", "GCS config scopes")

	setCmd.Flags().StringVar(&req.azureConfigAccount, azure.Kind+"."+azure.ConfigAccount, "", "Azure config account")
	setCmd.Flags().StringVar(&req.azureConfigKey, azure.Kind+"."+azure.ConfigKey, "", "Azure config key")

	setCmd.Flags().StringVar(&req.localConfigKeyPath, local.Kind+"."+local.ConfigKeyPath, "", "Local config key path")

	setCmd.Flags().StringVar(&req.swiftConfigKey, swift.Kind+"."+swift.ConfigKey, "", "Swift config key")
	setCmd.Flags().StringVar(&req.swiftConfigTenantAuthURL, swift.Kind+"."+swift.ConfigTenantAuthURL, "", "Swift teanant auth url")
	setCmd.Flags().StringVar(&req.swiftConfigTenantName, swift.Kind+"."+swift.ConfigTenantName, "", "Swift tenant name")
	setCmd.Flags().StringVar(&req.swiftConfigUsername, swift.Kind+"."+swift.ConfigUsername, "", "Swift username")

	return setCmd
}

func setContext(req *setContextRequest) {
	nc := &otx.Context{
		Name:     req.Name,
		Provider: req.Provider,
		Config:   stow.ConfigMap{},
	}
	switch req.Provider {
	case s3.Kind, "aws":
		nc.Provider = s3.Kind
		if req.s3ConfigAccessKeyID != "" {
			nc.Config[s3.ConfigAccessKeyID] = req.s3ConfigAccessKeyID
		}
		if req.s3ConfigEndpoint != "" {
			nc.Config[s3.ConfigEndpoint] = req.s3ConfigEndpoint
		}
		if req.s3ConfigRegion != "" {
			nc.Config[s3.ConfigRegion] = req.s3ConfigRegion
		}
		if req.s3ConfigSecretKey != "" {
			nc.Config[s3.ConfigSecretKey] = req.s3ConfigSecretKey
		}
	case gcs.Kind, "gs":
		nc.Provider = gcs.Kind
		if req.gcsConfigJSONKeyPath != "" {
			jsonKey, err := ioutil.ReadFile(req.gcsConfigJSONKeyPath)
			term.ExitOnError(err)
			nc.Config[gcs.ConfigJSON] = string(jsonKey)
		}
		if req.gcsConfigProjectId != "" {
			nc.Config[gcs.ConfigProjectId] = req.gcsConfigProjectId
		}
		if req.gcsConfigScopes != "" {
			nc.Config[gcs.ConfigScopes] = req.gcsConfigScopes
		}
	case azure.Kind:
		if req.azureConfigAccount != "" {
			nc.Config[azure.ConfigAccount] = req.azureConfigAccount
		}
		if req.azureConfigKey != "" {
			nc.Config[azure.ConfigKey] = req.azureConfigKey
		}
	case local.Kind:
		if req.localConfigKeyPath != "" {
			nc.Config[local.ConfigKeyPath] = req.localConfigKeyPath
		}
	case swift.Kind:
		if req.swiftConfigKey != "" {
			nc.Config[swift.ConfigKey] = req.swiftConfigKey
		}
		if req.swiftConfigTenantAuthURL != "" {
			nc.Config[swift.ConfigTenantAuthURL] = req.swiftConfigTenantAuthURL
		}
		if req.swiftConfigTenantName != "" {
			nc.Config[swift.ConfigTenantName] = req.swiftConfigTenantName
		}
		if req.swiftConfigUsername != "" {
			nc.Config[swift.ConfigUsername] = req.swiftConfigUsername
		}
	default:
		term.Fatalln("Unknown provider:" + req.Provider)
	}

	config, _ := otx.LoadConfig()
	if config == nil {
		config = &otx.OSMConfig{
			Contexts: make([]*otx.Context, 0),
		}
	}

	found := false
	for i := range config.Contexts {
		if config.Contexts[i].Name == req.Name {
			config.Contexts[i] = nc
			found = true
			break
		}
	}
	if !found {
		config.Contexts = append(config.Contexts, nc)
	}
	config.CurrentContext = req.Name
	err := config.Save()
	term.ExitOnError(err)
}

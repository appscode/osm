package context

import (
	"errors"
	"io/ioutil"
	"os"

	"github.com/graymeta/stow"
	homeDir "github.com/mitchellh/go-homedir"
	"gopkg.in/yaml.v2"
)

var (
	home, _    = homeDir.Dir()
	configPath = home + "/.osm/config"
)

type Context struct {
	Name     string         `json:"name"`
	Provider string         `json:"provider"`
	Config   stow.ConfigMap `json:"config"`
}

type OSMConfig struct {
	Contexts       []*Context `json:"contexts"`
	CurrentContext string     `json:"current-context"`
}

func LoadConfig() (*OSMConfig, error) {
	if _, err := os.Stat(configPath); err != nil {
		return nil, err
	}
	os.Chmod(configPath, 0600)

	config := &OSMConfig{}
	bytes, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(bytes, config)
	return config, err
}

func (config *OSMConfig) Save() error {
	data, err := yaml.Marshal(config)
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(configPath, data, 0600); err != nil {
		return err
	}
	return nil
}

func (config *OSMConfig) Dial(cliCtx string) (stow.Location, error) {
	ctx := config.CurrentContext
	if cliCtx != "" {
		ctx = cliCtx
	}
	for _, osmCtx := range config.Contexts {
		if osmCtx.Name == ctx {
			return stow.Dial(osmCtx.Provider, osmCtx.Config)
		}
	}
	return nil, errors.New("Failed to determine context.")
}
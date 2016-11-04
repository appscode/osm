package context

import (
	"io/ioutil"
	"os"

	homeDir "github.com/mitchellh/go-homedir"
	"gopkg.in/yaml.v2"
)

var (
	home, _ = homeDir.Dir()
)

func Home() string {
	return home
}

type Attrs struct {
	CredentialDir string `yaml:"credential_dir"`
	Provider      string `yaml:"provider"`
}

type Context struct {
	ContextData struct {
		CredentialDir string `yaml:"credential_dir"`
		Provider      string `yaml:"provider"`
	} `yaml:"context"`
	Name string `yaml:"name"`
}

type ConfigData struct {
	Contexts       []*Context `yaml:"contexts"`
	CurrentContext string     `yaml:"current-context"`
}

func GetConfigData() (*ConfigData, error) {
	config := &ConfigData{}
	path := Home() + "/.osm/config"
	d, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, nil
	}
	err = yaml.Unmarshal(d, config)
	return config, nil
}

func SetConfigData(config *ConfigData) error {
	path := Home() + "/.osm/config"
	data, err := yaml.Marshal(config)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(path, data, os.ModePerm); err != nil {
		return err
	}

	return nil
}

package config

import (
	"io/ioutil"
	"os"

	"github.com/kobsio/kobs/cmd/kobs/plugins"
	"github.com/kobsio/kobs/pkg/api/clusters"

	"sigs.k8s.io/yaml"
)

// Config is the complete configuration for kobs.
type Config struct {
	Clusters clusters.Config `json:"clusters"`
	Plugins  plugins.Config  `json:"plugins"`
}

// Load the configuration for kobs. Most of the configuration options are available as command-line flag, but we also
// need some more complex configuration options, which can be set via a config file in yaml format. The configuration
// file can contain environment variables in the following format: "${NAME_OF_THE_ENVIRONMENT_VARIABLE}".
func Load(file string) (*Config, error) {
	configContent, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	configContent = []byte(os.ExpandEnv(string(configContent)))
	cfg := &Config{}
	if err := yaml.Unmarshal(configContent, cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

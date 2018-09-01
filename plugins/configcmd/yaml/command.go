package yaml

import (
	"github.com/MrLYC/gokit/config"
	"github.com/MrLYC/gokit/plugins"
	"github.com/MrLYC/gokit/plugins/configcmd"
	"gopkg.in/yaml.v2"
)

func dumps(conf config.Configuration) (string, error) {
	data, err := yaml.Marshal(conf.AllSettings())
	return string(data), err
}

// New :
func New() plugins.Plugin {
	return configcmd.New(dumps)
}

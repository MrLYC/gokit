package yaml

import (
	"github.com/MrLYC/gokit/plugins"
	"github.com/MrLYC/gokit/plugins/config"
	"gopkg.in/yaml.v2"
)

// Plugin :
type Plugin struct {
	*config.Plugin
}

// Dumps :
func (p *Plugin) Dumps() (string, error) {
	data, err := yaml.Marshal(p.Configuration)
	return string(data), err
}

// New :
func New() plugins.Plugin {
	return &Plugin{
		Plugin: config.New().(*config.Plugin),
	}
}

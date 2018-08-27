package plugins

import (
	"github.com/MrLYC/gokit/config"
)

// Plugin :
type Plugin interface {
	Start(config.Configuration) error
	Stop() error
}

// BasePlugin :
type BasePlugin struct {
	Configuration config.Configuration
}

// Start :
func (p *BasePlugin) Start(conf config.Configuration) error {
	p.Configuration = conf
	return nil
}

// Stop :
func (p *BasePlugin) Stop() error {
	return nil
}

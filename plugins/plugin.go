package plugins

import (
	"flag"
	"fmt"

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

// BaseCommandPlugin :
type BaseCommandPlugin struct {
	BasePlugin
	name     string
	group    string
	synopsis string
}

// Start :
func (p *BaseCommandPlugin) Start(conf config.Configuration) error {
	err := p.BasePlugin.Start(conf)
	return err
}

// Stop :
func (p *BaseCommandPlugin) Stop() error {
	return nil
}

// Name :
func (p *BaseCommandPlugin) Name() string {
	return p.name
}

// Group :
func (p *BaseCommandPlugin) Group() string {
	return p.group
}

// Synopsis :
func (p *BaseCommandPlugin) Synopsis() string {
	return p.synopsis
}

// Usage :
func (p *BaseCommandPlugin) Usage() string {
	return fmt.Sprintf(`%s
  %s
`, p.name, p.synopsis)
}

// SetFlags :
func (*BaseCommandPlugin) SetFlags(f *flag.FlagSet) {

}

// NewBaseCommandPlugin :
func NewBaseCommandPlugin(name string, group string, synopsis string) *BaseCommandPlugin {
	return &BaseCommandPlugin{
		name:     name,
		group:    group,
		synopsis: synopsis,
	}
}

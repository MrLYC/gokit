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
	name  string
	group string
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
func (*BaseCommandPlugin) Synopsis() string {
	return ""
}

// Usage :
func (p *BaseCommandPlugin) Usage() string {
	return fmt.Sprintf(`%s
  %s
`, p.name, p.Synopsis())
}

// SetFlags :
func (*BaseCommandPlugin) SetFlags(f *flag.FlagSet) {

}

// NewBaseCommandPlugin :
func NewBaseCommandPlugin(name string, group string) *BaseCommandPlugin {
	return &BaseCommandPlugin{
		name:  name,
		group: group,
	}
}

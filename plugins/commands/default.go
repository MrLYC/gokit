package commands

import (
	"github.com/MrLYC/gokit/config"
	"github.com/MrLYC/gokit/plugins"
	"github.com/google/subcommands"
)

// DefaultPlugin :
type DefaultPlugin struct {
	plugins.BasePlugin
	group string
}

// Start :
func (p *DefaultPlugin) Start(conf config.Configuration) error {
	err := p.BasePlugin.Start(conf)
	if err != nil {
		return err
	}
	subcommands.Register(subcommands.HelpCommand(), p.group)
	subcommands.Register(subcommands.FlagsCommand(), p.group)
	subcommands.Register(subcommands.CommandsCommand(), p.group)
	return nil
}

// NewDefaultPlugin :
func NewDefaultPlugin(group string) plugins.Plugin {
	return &DefaultPlugin{
		group: group,
	}
}

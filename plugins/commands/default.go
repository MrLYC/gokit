package commands

import (
	"github.com/MrLYC/gokit/config"
	"github.com/MrLYC/gokit/plugins"
	"github.com/google/subcommands"
)

// DefaultPlugin :
type DefaultPlugin struct {
	plugins.BasePlugin
}

// Start :
func (p *DefaultPlugin) Start(conf config.Configuration) error {
	err := p.BasePlugin.Start(conf)
	if err != nil {
		return err
	}
	subcommands.Register(subcommands.HelpCommand(), config.Attrs.GetString("plugins.commands.group.help"))
	subcommands.Register(subcommands.FlagsCommand(), config.Attrs.GetString("plugins.commands.group.flags"))
	subcommands.Register(subcommands.CommandsCommand(), config.Attrs.GetString("plugins.commands.group.commands"))
	return nil
}

// NewDefaultPlugin :
func NewDefaultPlugin() plugins.Plugin {
	return &DefaultPlugin{}
}

func init() {
	config.Attrs.SetDefault("plugins.commands.group.help", "")
	config.Attrs.SetDefault("plugins.commands.group.flags", "")
	config.Attrs.SetDefault("plugins.commands.group.commands", "")
}

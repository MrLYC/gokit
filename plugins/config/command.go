package config

import (
	"context"
	"flag"
	"fmt"

	"encoding/json"

	"github.com/MrLYC/gokit/config"
	"github.com/MrLYC/gokit/plugins"
	"github.com/google/subcommands"
)

// Plugin :
type Plugin struct {
	*plugins.BaseCommandPlugin
}

// Start :
func (p *Plugin) Start(conf config.Configuration) error {
	err := p.BasePlugin.Start(conf)
	if err != nil {
		return err
	}
	subcommands.Register(p, p.Group())
	return nil
}

// Dumps :
func (p *Plugin) Dumps() (string, error) {
	data, err := json.Marshal(p.Configuration)
	return string(data), err
}

// Execute :
func (p *Plugin) Execute(cxt context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	data, err := p.Dumps()
	if err != nil {
		panic(err)
	}
	fmt.Printf(data)
	return subcommands.ExitSuccess
}

// New :
func New() plugins.Plugin {
	return &Plugin{
		BaseCommandPlugin: plugins.NewBaseCommandPlugin("confinfo", "config", "print configurations"),
	}
}

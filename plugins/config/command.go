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
	plugins.BasePlugin
	group string
}

// Start :
func (p *Plugin) Start(conf config.Configuration) error {
	err := p.BasePlugin.Start(conf)
	subcommands.Register(p, p.group)
	return err
}

// Name :
func (*Plugin) Name() string {
	return "config"
}

// Synopsis :
func (*Plugin) Synopsis() string {
	return "Print configurations"
}

// Usage :
func (*Plugin) Usage() string {
	return `config
  Print configurations.
`
}

// SetFlags :
func (*Plugin) SetFlags(f *flag.FlagSet) {
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

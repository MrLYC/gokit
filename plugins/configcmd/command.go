package configcmd

import (
	"context"
	"flag"
	"fmt"

	"github.com/MrLYC/gokit/config"
	"github.com/MrLYC/gokit/plugins"
	"github.com/google/subcommands"
)

// Plugin :
type Plugin struct {
	*plugins.BaseCommandPlugin
	dumpFunc func(config.Configuration) (string, error)
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

// Execute :
func (p *Plugin) Execute(cxt context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	data, err := p.dumpFunc(p.Configuration)
	if err != nil {
		panic(err)
	}
	fmt.Printf(data)
	return subcommands.ExitSuccess
}

// New :
func New(dumpFunc func(config.Configuration) (string, error)) plugins.Plugin {
	return &Plugin{
		BaseCommandPlugin: plugins.NewBaseCommandPlugin(
			config.Attrs.GetString("plugins.configcmd.name"),
			config.Attrs.GetString("plugins.configcmd.group"),
			"print configurations",
		),
		dumpFunc: dumpFunc,
	}
}

func init() {
	config.Attrs.SetDefault("plugins.configcmd.name", "confinfo")
	config.Attrs.SetDefault("plugins.configcmd.group", "config")
}

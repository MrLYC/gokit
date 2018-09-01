package version

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

	short bool

	version   string
	buildHash string
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

// SetFlags :
func (p *Plugin) SetFlags(f *flag.FlagSet) {
	f.BoolVar(&p.short, "s", false, "print version only")
}

// Execute :
func (p *Plugin) Execute(cxt context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if p.short {
		fmt.Printf("%v\n", p.version)
	} else {
		fmt.Printf(
			"Version: %v, BuildHash: %v\n", p.version, p.buildHash,
		)
	}
	return subcommands.ExitSuccess
}

// NewVersionPlugin :
func NewVersionPlugin(version string, buildHash string) plugins.Plugin {
	return &Plugin{
		BaseCommandPlugin: plugins.NewBaseCommandPlugin(
			config.Attrs.GetString("plugins.version.name"),
			config.Attrs.GetString("plugins.version.group"),
			"print version infomations",
		),
		version:   version,
		buildHash: buildHash,
	}
}

func init() {
	config.Attrs.SetDefault("plugins.version.name", "version")
	config.Attrs.SetDefault("plugins.version.group", "version")
}

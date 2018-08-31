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

// Execute :
func (p *Plugin) Execute(cxt context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	fmt.Printf(
		"Version: %v, BuildHash: %v\n", p.version, p.buildHash,
	)
	return subcommands.ExitSuccess
}

// NewVersionPlugin :
func NewVersionPlugin(version string, buildHash string) plugins.Plugin {
	return &Plugin{
		BaseCommandPlugin: plugins.NewBaseCommandPlugin("version", "version", "print version infomations"),
		version:           version,
		buildHash:         buildHash,
	}
}

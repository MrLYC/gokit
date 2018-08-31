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
	plugins.BasePlugin

	group     string
	version   string
	buildHash string
}

// Start :
func (p *Plugin) Start(conf config.Configuration) error {
	err := p.BasePlugin.Start(conf)
	subcommands.Register(p, p.group)
	return err
}

// Name :
func (p *Plugin) Name() string {
	return "version"
}

// Synopsis :
func (p *Plugin) Synopsis() string {
	return "Print version infomations"
}

// Usage :
func (p *Plugin) Usage() string {
	return `version
  Print version infomations.
`
}

// SetFlags :
func (p *Plugin) SetFlags(f *flag.FlagSet) {

}

// Execute :
func (p *Plugin) Execute(cxt context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	fmt.Printf(
		"Version: %v, BuildHash: %v\n", p.version, p.buildHash,
	)
	return subcommands.ExitSuccess
}

// NewVersionPlugin :
func NewVersionPlugin(version string, buildHash string, group string) plugins.Plugin {
	return &Plugin{
		group:     group,
		version:   version,
		buildHash: buildHash,
	}
}

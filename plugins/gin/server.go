package gin

import (
	"context"
	"flag"
	"fmt"

	"github.com/MrLYC/gokit/config"
	"github.com/MrLYC/gokit/logging"
	"github.com/MrLYC/gokit/plugins"
	"github.com/gin-gonic/gin"
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

// Route :
func (p *Plugin) Route(engine *gin.Engine) {

}

// Execute :
func (p *Plugin) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	address := fmt.Sprintf("%v:%v", p.Configuration.GetString("http.host"), p.Configuration.GetInt("http.port"))
	if !p.Configuration.GetBool("debug") {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.Default()
	p.Route(engine)
	logging.Infof("http server listening on: %s", address)
	engine.Run(address)
	return subcommands.ExitSuccess
}

// New :
func New() plugins.Plugin {
	return &Plugin{
		BaseCommandPlugin: plugins.NewBaseCommandPlugin("server", "http", "run gin server"),
	}
}

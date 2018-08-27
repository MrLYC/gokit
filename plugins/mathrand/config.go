package mathrand

import (
	"math/rand"
	"time"

	"github.com/MrLYC/gokit/config"
	"github.com/MrLYC/gokit/plugins"
)

// Plugin :
type Plugin struct{}

// Start :
func (p *Plugin) Start(conf config.Configuration) error {
	rand.Seed(time.Now().UnixNano())
	return nil
}

// Stop :
func (p *Plugin) Stop() error {
	return nil
}

// InitSeed : init rand seed
func InitSeed() plugins.Plugin {
	return &Plugin{}
}

package plugins_test

import (
	"testing"

	"github.com/MrLYC/gokit/config"
	"github.com/MrLYC/gokit/plugins"
)

type mockPlugin struct {
	plugins.BasePlugin
	value string
}

// Start :
func (p *mockPlugin) Start(conf config.Configuration) error {
	p.BasePlugin.Start(conf)
	p.value = "started"
	return nil
}

// Stop :
func (p *mockPlugin) Stop() error {
	p.value = "stoped"
	return nil
}

func TestPluginManager(t *testing.T) {
	p := &mockPlugin{}
	pm := plugins.NewPluginManager(p)
	var err error

	err = pm.Start(config.NewMapConfiguration())
	if err != nil || p.value != "started" {
		t.Errorf("plugin manager start failed: %s", err)
	}

	err = pm.Stop()
	if err != nil || p.value != "stoped" {
		t.Errorf("plugin manager stop failed: %s", err)
	}
}

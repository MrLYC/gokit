package plugins

import "github.com/MrLYC/gokit/config"

// PluginManager :
type PluginManager struct {
	BasePlugin
	plugins []Plugin
}

// Start :
func (p *PluginManager) Start(conf config.Configuration) error {
	err := p.BasePlugin.Start(conf)
	for _, plugin := range p.plugins {
		err = plugin.Start(conf)
		if err != nil {
			return err
		}
	}
	return nil
}

// Stop :
func (p *PluginManager) Stop() error {
	err := p.BasePlugin.Stop()
	for _, plugin := range p.plugins {
		err = plugin.Stop()
		if err != nil {
			return err
		}
	}
	return nil
}

// NewPluginManager :
func NewPluginManager(plugins ...Plugin) Plugin {
	p := PluginManager{
		plugins: plugins,
	}
	return &p
}

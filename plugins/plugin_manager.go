package plugins

// PluginManager :
type PluginManager struct {
	BasePlugin
	plugins []Plugin
}

// NewPluginManager :
func NewPluginManager(plugins ...Plugin) Plugin {
	p := PluginManager{
		plugins: plugins,
	}
	return &p
}

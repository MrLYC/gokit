package json

import (
	"encoding/json"

	"github.com/MrLYC/gokit/config"
	"github.com/MrLYC/gokit/plugins"
	"github.com/MrLYC/gokit/plugins/configcmd"
)

func dumps(conf config.Configuration) (string, error) {
	data, err := json.Marshal(conf.AllSettings())
	return string(data), err
}

// New :
func New() plugins.Plugin {
	return configcmd.New(dumps)
}

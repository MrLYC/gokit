package logrus

import (
	"github.com/MrLYC/gokit/config"
	"github.com/MrLYC/gokit/logging"
	"github.com/MrLYC/gokit/plugins"
	log "github.com/sirupsen/logrus"
)

// StdConfigure :
type StdConfigure struct {
	*Configure
}

// Start :
func (c *StdConfigure) Start(conf config.Configuration) error {
	err := c.Configure.Start(conf)
	if err != nil {
		return err
	}
	logging.SetLogger(c.logger)
	return nil
}

// NewStdConfigure :
func NewStdConfigure(logger *log.Logger) plugins.Plugin {
	return &StdConfigure{
		Configure: New(logger).(*Configure),
	}
}

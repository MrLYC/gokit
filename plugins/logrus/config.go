package logrus

import (
	"io"
	"os"

	"github.com/MrLYC/gokit/config"
	"github.com/MrLYC/gokit/plugins"
	log "github.com/sirupsen/logrus"
)

// Configure :
type Configure struct {
	plugins.BasePlugin
	logger  *log.Logger
	writers []io.WriteCloser
}

// Start :
func (c *Configure) Start(conf config.Configuration) error {
	err := c.BasePlugin.Start(conf)
	if err != nil {
		return err
	}

	conf.SetDefault("logging.level", "info")
	conf.SetDefault("logging.to_stdout", false)
	conf.SetDefault("logging.to_stderr", true)
	conf.SetDefault("logging.to_file", []string{})

	logger := c.logger

	level, err := log.ParseLevel(conf.GetString("logging.level"))
	if err != nil {
		panic(err)
	}
	logger.SetLevel(level)

	writers := make([]io.Writer, 0)
	if conf.GetBool("logging.to_stdout") {
		writers = append(writers, os.Stdout)
	}
	if conf.GetBool("logging.to_stderr") {
		writers = append(writers, os.Stderr)
	}
	c.writers = make([]io.WriteCloser, 0)
	for _, path := range conf.GetStringSlice("logging.to_file") {
		file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			panic(err)
		}
		writers = append(writers, file)
		c.writers = append(c.writers, file)
	}
	if len(writers) > 1 {
		writer := io.MultiWriter(writers...)
		logger.SetOutput(writer)
	}

	return nil
}

// Stop :
func (c *Configure) Stop() error {
	err := c.BasePlugin.Stop()
	if err != nil {
		return err
	}
	for _, writer := range c.writers {
		err := writer.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

// New :
func New(logger *log.Logger) plugins.Plugin {
	return &Configure{
		logger: logger,
	}
}

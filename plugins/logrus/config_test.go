package logrus_test

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/mrlyc/gokit/plugins/logrus"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func TestLogrusConfigureStart(t *testing.T) {
	conf := viper.New()
	logger := log.New()
	plugin := logrus.New(logger)

	conf.Set("logging.level", "DEBUG")

	err := plugin.Start(conf)
	if err != nil {
		t.Error(err)
	}

	if logger.Level != log.DebugLevel {
		t.Errorf("logger level not set: %v", logger.Level)
	}

	err = plugin.Stop()
	if err != nil {
		t.Error(err)
	}
}

func TestLogrusConfigureDefault(t *testing.T) {
	conf := viper.New()
	logger := log.New()
	plugin := logrus.New(logger)

	err := plugin.Start(conf)
	if err != nil {
		t.Error(err)
	}

	if logger.Level != log.InfoLevel {
		t.Errorf("logger level not set: %v", logger.Level)
	}

	err = plugin.Stop()
	if err != nil {
		t.Error(err)
	}
}

func TestLogrusConfigureStop(t *testing.T) {
	conf := viper.New()
	logger := log.New()
	plugin := logrus.New(logger)

	file, err := ioutil.TempFile("", "gokil_test")
	if err != nil {
		t.Error(err)
	}
	defer file.Close()

	conf.Set("logging.to_file", []string{file.Name()})

	err = plugin.Start(conf)
	if err != nil {
		t.Error(err)
	}

	logger.Infof("ok")

	err = plugin.Stop()
	if err != nil {
		t.Error(err)
	}

	logger.Infof("failed")

	data, err := ioutil.ReadAll(file)
	if err != nil {
		t.Error(err)
	}

	logs := string(data)
	if strings.Count(logs, "ok") != 1 || strings.Count(logs, "failed") != 0 {
		t.Errorf("write log error: %s", logs)
	}
}

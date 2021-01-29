package logger

import (
	"os"

	"github.com/frauniki/ping-mesh/pkg/config"
	log "github.com/sirupsen/logrus"
)

func Init() error {
	level, err := log.ParseLevel(config.Config.LogLevel)
	if err != nil {
		return err
	}

	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(level)

	return nil
}

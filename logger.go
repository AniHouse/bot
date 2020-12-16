package bot

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/anihouse/bot/config"
	"github.com/sirupsen/logrus"
)

// Errors
var (
	ErrNoLogger        = errors.New("Logger config doesn't exist")
	ErrUnknowFormatter = errors.New("Unknow formatter")
)

func Logger(cfg *config.Log) (*logrus.Logger, error) {
	if cfg == nil {
		return nil, ErrNoLogger
	}

	err := os.MkdirAll(filepath.Dir(cfg.File), 0775)
	if err != nil {
		return nil, err
	}

	logger := logrus.New()
	file, err := os.OpenFile(cfg.File, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}
	logger.SetOutput(file)

	level, err := logrus.ParseLevel(cfg.Level)
	if err != nil {
		return nil, err
	}
	logger.SetLevel(level)

	var formatter logrus.Formatter
	switch cfg.Formatter {
	case "text":
		formatter = &logrus.TextFormatter{}
	case "json":
		formatter = &logrus.JSONFormatter{}
	default:
		return nil, ErrUnknowFormatter
	}
	logger.SetFormatter(formatter)

	return logger, nil
}

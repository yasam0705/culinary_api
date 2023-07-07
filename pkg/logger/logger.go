package logger

import (
	"github/culinary_api/config"

	"go.uber.org/zap"
)

func New(cfg *config.Config) (*logger, error) {

	zapCfg := zap.NewProductionConfig()
	zapCfg.Development = true
	zapCfg.Level = getLevel(cfg.LogLevel)
	zapCfg.OutputPaths = []string{cfg.App + ".log"}

	log, err := zapCfg.Build()
	if err != nil {
		return nil, err
	}
	return &logger{log: log}, nil
}

func getLevel(level string) zap.AtomicLevel {
	switch level {
	case "debug":
		return zap.NewAtomicLevelAt(zap.DebugLevel)
	case "info":
		return zap.NewAtomicLevelAt(zap.InfoLevel)
	case "warn":
		return zap.NewAtomicLevelAt(zap.WarnLevel)
	case "error":
		return zap.NewAtomicLevelAt(zap.ErrorLevel)
	case "dpanic":
		return zap.NewAtomicLevelAt(zap.DPanicLevel)
	case "panic":
		return zap.NewAtomicLevelAt(zap.PanicLevel)
	case "fatal":
		return zap.NewAtomicLevelAt(zap.FatalLevel)
	default:
		return zap.NewAtomicLevelAt(zap.ErrorLevel)
	}
}

package server

import "go.uber.org/zap"

// NewZaplogger creates a new zap logger.
func NewZaplogger() (*zap.Logger, error) {
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.TimeKey = ""
	return cfg.Build()
}

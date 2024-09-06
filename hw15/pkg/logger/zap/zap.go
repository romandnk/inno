package zaplogger

import (
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New(cfg Config) (*zap.Logger, error) {
	err := cfg.Validate()
	if err != nil {
		return nil, err
	}

	//nolint:errcheck
	// skip, because it was validated above
	level, _ := zap.ParseAtomicLevel(cfg.Level)

	config := getConfigByEnv(cfg.Env)
	config.Level = level
	config.Encoding = cfg.Format
	config.OutputPaths = []string{"stdout"}
	config.ErrorOutputPaths = []string{"stdout"}

	logger, err := config.Build()
	if err != nil {
		return nil, err
	}

	return logger, nil
}

func getConfigByEnv(env string) zap.Config {
	var cfg zap.Config
	switch strings.ToLower(env) {
	case "local":
		cfg = getLocalConfig()
	case "dev":
		cfg = getDevConfig()
	case "prod":
		cfg = getProdConfig()
	}
	return cfg
}

func getLocalConfig() zap.Config {
	return zap.Config{
		Development: true,
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "msg",
			LevelKey:   "lvl",
			TimeKey:    "ts",
			EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
				enc.AppendString(t.Format(time.RFC3339))
			},
			EncodeLevel:      zapcore.CapitalColorLevelEncoder,
			ConsoleSeparator: " | ",
		},
	}
}

func getDevConfig() zap.Config {
	return zap.Config{
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "msg",
			LevelKey:   "lvl",
			TimeKey:    "ts",
			EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
				enc.AppendString(t.Format(time.RFC3339))
			},
			EncodeLevel: func(lvl zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
				enc.AppendString(lvl.String())
			},
			ConsoleSeparator: " | ",
		},
	}
}

func getProdConfig() zap.Config {
	return zap.Config{
		DisableCaller:     true,
		DisableStacktrace: true,
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "msg",
			LevelKey:   "lvl",
			TimeKey:    "ts",
			EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
				enc.AppendString(t.Format(time.RFC3339))
			},
			EncodeLevel: func(lvl zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
				enc.AppendString(lvl.String())
			},
		},
	}
}

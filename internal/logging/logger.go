/*
Copyright 2025 ptrvsrg.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package logging

import (
	"os"
	"strings"

	"github.com/samber/lo"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/ptrvsrg/casdoor-operator/config"
)

var (
	loggingFormats = []string{"json", "console"}
)

func SetupLogger(cfg config.LoggingConfig) {
	encoding := parseEncoding(cfg.Format)

	zapConfig := zap.Config{
		Level:            zap.NewAtomicLevelAt(parseLevel(cfg.Level)),
		Development:      encoding == "console",
		Encoding:         encoding,
		EncoderConfig:    getEncoderConfig(encoding),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stdout"},
		InitialFields: map[string]interface{}{
			"pid": os.Getpid(),
		},
	}

	logger := zap.Must(zapConfig.Build())
	zap.ReplaceGlobals(logger)
}

func parseEncoding(format string) string {
	encoding := "json"
	if lo.Contains(loggingFormats, format) {
		encoding = format
	}

	return encoding
}

func getEncoderConfig(encoding string) zapcore.EncoderConfig {
	encoderCfg := zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
	}

	if encoding == "console" {
		encoderCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
		encoderCfg.EncodeCaller = zapcore.FullCallerEncoder
	} else {
		encoderCfg.EncodeLevel = zapcore.LowercaseLevelEncoder
		encoderCfg.EncodeCaller = zapcore.ShortCallerEncoder
	}

	return encoderCfg
}

func parseLevel(level string) zapcore.Level {
	lvl := zap.InfoLevel
	switch strings.ToLower(level) {
	case "debug":
		lvl = zap.DebugLevel
	case "info":
		lvl = zap.InfoLevel
	case "warn":
		lvl = zap.WarnLevel
	case "error":
		lvl = zap.ErrorLevel
	case "fatal":
		lvl = zap.FatalLevel
	}

	return lvl
}

package logger

import (
	"github.com/Niexiawei/golang-utils/pathtool"
	"github.com/Niexiawei/logrotate"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"path/filepath"
	"time"
)

var (
	Logger     *zap.SugaredLogger
	loggerPath string
)

func SetupLogger() {
	loggerPath = pathtool.RuntimePath + "/logs/logger.log"
	{
		dir := filepath.Dir(loggerPath)
		if ok, _ := pathtool.PathExists(dir); !ok {
			_ = os.MkdirAll(dir, 0777)
		}
	}

	core := zapcore.NewTee(
		zapcore.NewCore(stdoutEncoder(), stdoutWriter(), zapcore.DebugLevel),
		zapcore.NewCore(fileEncoder(), fileWriter(), zap.LevelEnablerFunc(func(level zapcore.Level) bool {
			return level >= zap.InfoLevel
		})),
	)
	logger := zap.New(core, zap.AddCaller())
	Logger = logger.Sugar()
}

func stdoutEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func fileEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func fileWriter() zapcore.WriteSyncer {
	filePath, _ := filepath.Abs(loggerPath)
	LoggerWriter, err := logrotate.NewRotateLog(
		filePath+".2006-01-02",
		logrotate.WithCurLogLinkname(filePath),
		logrotate.WithDeleteExpiredFile(time.Hour*24*7, "logger.log.*"),
		logrotate.WithRotateTime(time.Hour),
	)
	if err != nil {
		log.Fatalln(err)
	}
	return zapcore.AddSync(LoggerWriter)
}

func stdoutWriter() zapcore.WriteSyncer {
	return zapcore.AddSync(os.Stdout)
}

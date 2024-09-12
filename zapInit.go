package zapInit

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func Init(path, level string, isWrapperZap bool) {
	var logger *zap.Logger
	l, _ := zapcore.ParseLevel(level)
	fw, _ := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// 修改日志输出格式
	encoderConf := zap.NewProductionEncoderConfig()
	encoderConf.TimeKey = "time"
	encoderConf.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConf.EncodeLevel = zapcore.CapitalLevelEncoder

	encoder := zapcore.NewJSONEncoder(encoderConf)
	if isWrapperZap {
		logger = zap.New(zapcore.NewCore(encoder, zapcore.AddSync(fw), l), zap.AddCaller(), zap.AddCallerSkip(1))
	} else {
		logger = zap.New(zapcore.NewCore(encoder, zapcore.AddSync(fw), l))
	}
	zap.ReplaceGlobals(logger)
}

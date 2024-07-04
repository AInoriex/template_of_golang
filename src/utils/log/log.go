package log

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/TarsCloud/TarsGo/tars"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

var logger *zap.Logger

// @Title   初始化日志
// @Description Tars拦截标准输出并存储成日志，暂无须切割保存日志
//				TLOG2.log 日志输出，需要设置tars/web后台日志等级
// @Author  wzj  (2022/8/5 15:28)
// @Param	savepath		string		日志路径
// 			debug			bool		调试模式
// @Return	null
func InitLogger(logPath string, debug bool) *zap.Logger {
	//var err error
	cfg := tars.GetServerConfig()
	if logPath == "" {
		logPath = filepath.Join(cfg.LogPath, cfg.App, cfg.Server)
		if err := os.MkdirAll(logPath, 0755); err != nil {
			panic(err)
		}
	}

	if debug {

		// 日志切割
		module := "debug"
		// module := cfg.Server
		now := time.Now()
		hook := &lumberjack.Logger{
			Filename:   fmt.Sprintf("%s/%s-%04d%02d%02d%02d.log", logPath, module, now.Year(), now.Month(), now.Day(), now.Hour()), //filePath
			// Filename:   fmt.Sprintf("%s/%s.log", logPath, module), //filePath
			MaxSize:    500,                                       //100
			MaxBackups: 5,
			MaxAge:     30,    //days
			Compress:   true, // disabled by default
		}
		//defer hook.Close()

		writer := zapcore.AddSync(hook)
		core := zapcore.NewCore(getEncoder(), writer, zapcore.DebugLevel)
		logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
		logger.Info("InitLogger init debug mode")

		//core := zapcore.NewCore(getEncoder(), os.Stdout, zapcore.DebugLevel)
		//logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.ErrorLevel))
		//logger.Info("InitLogger init debug mode")

		//fullPath := filepath.Join(logPath, cfg.Server + ".log")
		//st, _ := os.Stat(fullPath)

		//writer := getLogWriter(fullPath)
		//core := zapcore.NewCore(getEncoder(), writer, zapcore.InfoLevel)
		//logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

		//logger, err = zap.NewProduction(zap.AddCaller(), zap.AddCallerSkip(1))
		//if err != nil {
		//	panic(fmt.Sprintf("Got error when init logger,  the error is '%v'", err))
		//}
	} else {

		// 日志切割
		module := cfg.Server
		//now := time.Now()
		hook := &lumberjack.Logger{
			//Filename:   fmt.Sprintf("%s/%s-%04d%02d%02d%02d.log", logPath, module, now.Year(), now.Month(), now.Day(), now.Hour()), //filePath
			Filename:   fmt.Sprintf("%s/%s.log", logPath, module), //filePath
			MaxSize:    500,                                       // 日志文件大小单位: M                                                                                                        //100
			MaxBackups: 50,                                        // 备份数
			MaxAge:     180,                                       // days
			Compress:   true,                                      // disabled by default
		}
		defer hook.Close()

		writer := zapcore.AddSync(hook)
		core := zapcore.NewCore(getEncoder(), writer, zapcore.InfoLevel)
		logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

		//logger, err = zap.NewProduction(zap.AddCaller(), zap.AddCallerSkip(1))
		//if err != nil {
		//	panic(fmt.Sprintf("Got error when init logger,  the error is '%v'", err))
		//}

		logger.Info("InitLogger init production mode")

	}
	zap.ReplaceGlobals(logger)
	return logger
}

func Sync() {
	logger.Sync()
}

func Info(msg string, fields ...zap.Field) {
	zap.L().Info(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	zap.L().Debug(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	zap.L().Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	zap.L().Error(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	zap.L().Fatal(msg, fields...)
}

func getLogWriter(savepath string) zapcore.WriteSyncer {
	file, _ := os.Create(savepath)
	return zapcore.AddSync(file)
}

func Infof(msg string, args ...interface{}) {
	zap.S().Infof(msg, args...)
}

func Infow(msg string, args ...interface{}) {
	zap.S().Infow(msg, args...)
}

func Warnf(msg string, args ...interface{}) {
	zap.S().Warnf(msg, args...)
}

func Warnw(msg string, args ...interface{}) {
	zap.S().Warnw(msg, args...)
}

func Errorf(msg string, args ...interface{}) {
	zap.S().Errorf(msg, args...)
}

func Errorw(msg string, args ...interface{}) {
	zap.S().Errorw(msg, args...)
}

func Debugf(msg string, args ...interface{}) {
	zap.S().Debugf(msg, args...)
}

func Debugw(msg string, args ...interface{}) {
	zap.S().Debugw(msg, args...)
}

func Fatalw(msg string, args ...interface{}) {
	zap.S().Fatalw(msg, args...)
}

func Fatalf(msg string, args ...interface{}) {
	zap.S().Fatalf(msg, args...)
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

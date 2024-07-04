package db

import (
	"fmt"
	"github.com/TarsCloud/TarsGo/tars"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/gorm/logger"

	"os"
	"path/filepath"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Db struct {
	*gorm.DB
	Tx   *gorm.DB
	Data interface{}
}

//定义自己的Writer
type LogWriter struct {
	Log *zap.Logger
}

func (sqlDb *Db) Begins() {
	sqlDb.Tx = sqlDb.DB.Begin()
}

func (sqlDb *Db) Commits() error {
	err := sqlDb.Tx.Commit().Error
	sqlDb.Tx = nil
	return err
}

func (sqlDb *Db) Rollbacks() error {
	err := sqlDb.Tx.Rollback().Error
	sqlDb.Tx = nil
	return err
}

type dialOptions func(db *gorm.DB)

/*
func Dial(host string, options ...dialOptions) *gorm.DB {
	con, err := gorm.Open("mysql", host)
	if err != nil {
		panic(fmt.Sprintf("Got error when connect database, arg = %v the error is '%v'", host, err))
	}
	//设置表名前缀
	//gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	//	return defaultTableName[:]
	//}
	//con.DB().SetConnMaxLifetime(2 * time.Hour)
	for _, option := range options {
		option(con)
	}
	return con
}


func DialAutoMigrate(value ...interface{}) dialOptions {
	return func(db *gorm.DB) {
		db.AutoMigrate(value...)
	}
}
*/
/*
func DialLogMode(enable bool) dialOptions {
	return func(db *gorm.DB) {
		db.LogMode(enable)
	}
}

func DialMaxCon(maxCon int) dialOptions {
	return func(db *gorm.DB) {
		db.DB().SetMaxOpenConns(maxCon)
		idle := maxCon
		if maxCon/3 >= 10 {
			idle = maxCon / 3
		}
		db.DB().SetMaxIdleConns(idle)
	}
}
*/

func NewLogWriter(l *zap.Logger) *LogWriter {
	return &LogWriter{Log: l}
}

//实现gorm/logger.Writer接口
func (m *LogWriter) Printf(format string, v ...interface{}) {
	m.Log.Info(fmt.Sprintf(format, v...))
}

// @Title   初始化Mysql
// @Description mysql基于zap写入日志文件, 但目前日志未切割
// @Author  wzj  (2022/8/10 18:05)
// @Param	args		string			数据库dsn
// 			maxCon		int				最大连接数
// @Return
func NewMysql(args string, maxCon int, arr []interface{}, enable bool) *gorm.DB {
	var con *gorm.DB
	var err error
	var gormCfg = &gorm.Config{
		//Logger:logger.Default.LogMode(logger.Info), //开启sql日志
	}
	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"

	if enable {
		//开启sql日志
		cfg := tars.GetServerConfig()
		logPath := filepath.Join(cfg.LogPath, cfg.App, cfg.Server)
		if err := os.MkdirAll(logPath, 0755); err != nil {
			panic(err)
		}

		// 日志切割
		//now := time.Now()
		hook := &lumberjack.Logger{
			//Filename:   fmt.Sprintf("%s/%s-%04d%02d%02d%02d.log", logPath, module, now.Year(), now.Month(), now.Day(), now.Hour()), //filePath
			Filename:   fmt.Sprintf("%s/db.log", logPath), //filePath
			MaxSize:    500,
			MaxBackups: 50,
			MaxAge:     180,  //days
			Compress:   true, // disabled by default
		}
		defer hook.Close()
		writer := zapcore.AddSync(hook)
		core := zapcore.NewCore(getEncoder(), writer, zapcore.InfoLevel)
		l := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

		gormCfg.Logger = logger.New(
			NewLogWriter(l),
			//log.New(file, "\r\n",log.LstdFlags | log.Lshortfile | log.LUTC), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
			logger.Config{
				SlowThreshold:             time.Second, // 慢 SQL 阈值
				LogLevel:                  logger.Info, // 日志级别
				IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
				Colorful:                  false,       // 禁用彩色打印
			},
		)
	}

	con, err = gorm.Open(mysql.Open(args), gormCfg)
	if err != nil {
		panic(fmt.Sprintf("Got error when connect database, arg = %v the error is '%v'", args, err))
	}
	//设置表名前缀
	//gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	//	return defaultTableName[:]
	//}

	sqlDB, err := con.DB()
	if err != nil {
		panic(fmt.Sprintf("Got error when get con.DB, arg = %v the error is '%v'", args, err))
	}

	idle := maxCon
	if maxCon/3 >= 10 {
		idle = maxCon / 3
	}

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(idle)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(2 * time.Hour)

	//若结构有变，则删除表重新创建
	//dropTable(con, arr...)
	//con.AutoMigrate(arr...) //若没有表，自动生成表
	return con
}

func (sqlDb *Db) Create() (error, interface{}) {
	//var err error
	result := sqlDb.DB.Create(sqlDb.Data)
	return result.Error, result.RowsAffected
}

func Create(db *gorm.DB, value interface{}) (error, interface{}) {
	//var err error
	result := db.Create(value)
	return result.Error, result.RowsAffected
}

func Save(db *gorm.DB, v interface{}) error {
	var err error
	err = db.Save(v).Error
	return err
}

// update语句 不允许使用orm特性
func getLogWriter(savepath string) zapcore.WriteSyncer {
	file, _ := os.OpenFile(savepath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	return zapcore.AddSync(file)
}

// 获取日志格式
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

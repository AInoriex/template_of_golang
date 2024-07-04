package main

import (
	"github.com/TarsCloud/TarsGo/tars"
	"github.com/gin-gonic/gin"
	"singapore/src/server/router/handler"
	"singapore/src/utils/alarm"
	"singapore/src/utils/config"
	"singapore/src/utils/cos"
	"singapore/src/utils/log"
	"singapore/src/utils/obs"
	"singapore/src/utils/rpc"
)

func main() {
	// 选择路由服务启动方式(2选1)
	main_local()
	// main_tars()

	defer func() {
		handler.Alarm2LarkLocal(alarm.LarkServerAlarmTextVariable{
			Level:  "WARN",
			Msg:    "[!] Service is DOWN",
			Detail: "Please check the log.",
		})
		log.Info("[MAIN] Exit.")
	}()

}

// 本地服务
func main_local() {
	config.InitCommonConfigLocal("./conf/local/common.json")
	log.InitLogger(config.CommonCfg.Log.SavePath, true)
	log.Info("初始化Config完毕")

	// cos.InitCos(config.CommonCfg.TencentOSS)
	// log.Info("初始化腾讯云Cos完毕")

	// obs.InitObs(config.CommonCfg.HuaweiOBS)
	// log.Info("初始化华为云Obs完毕")

	// db.InitMysqlAllLocal("root:123456@tcp(127.0.0.1:3306)", "meta", 5, db.Con_Main, true)
	// db.InitMysqlAll(config.DbCfg.Mysql.Host, config.DbCfg.Mysql.Db, config.DbCfg.Mysql.MaxCon, db.Con_Main, config.CommonCfg.OpenDbLog)
	// log.Info("初始化Mysql完毕")

	//uredis.InitRedis("127.0.0.1:6379", "", 0)
	// uredis.InitRedis(config.DbCfg.Redis.Host, config.DbCfg.Redis.Password, config.DbCfg.Redis.Db)
	// log.Info("初始化Redis完毕")

	g := gin.Default()
	handler.InitRouter(g)
	g.Run(config.CommonCfg.HttpServer.Addr)
}

// Tars服务
func main_tars() {
	// config.InitCommonConfigLocal("./conf/beta/common.json")
	config.InitCommonConfig()
	config.InitDBConfig()
	log.InitLogger(config.CommonCfg.Log.SavePath, false)
	log.Info("初始化Config完毕")

	cos.InitCos(config.CommonCfg.TencentOSS)
	log.Info("初始化腾讯云Cos完毕")

	obs.InitObs(config.CommonCfg.HuaweiOBS)
	log.Info("初始化华为云Obs完毕")

	// db.InitMysqlAllLocal("root:123456@tcp(127.0.0.1:3306)", "meta", 5, db.Con_Main, true)
	// db.InitMysqlAll(config.DbCfg.Mysql.Host, config.DbCfg.Mysql.Db, config.DbCfg.Mysql.MaxCon, db.Con_Main, config.CommonCfg.OpenDbLog)
	log.Info("初始化Mysql完毕")

	// uredis.InitRedis("127.0.0.1:6379", "", 0)
	// uredis.InitRedis(config.DbCfg.Redis.Host, config.DbCfg.Redis.Password, config.DbCfg.Redis.Db)
	log.Info("初始化Redis完毕")

	// 初始化路由
	if config.CommonCfg.Env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}
	mux := &tars.TarsHttpMux{}
	g := mux.GetGinEngine()
	handler.InitRouter(g)

	// 初始化tars
	servCfg := tars.GetServerConfig()
	tars.AddHttpServant(mux, servCfg.App+"."+servCfg.Server+".RouterObj") //Register http server

	// rpc调用
	rpc.InitCommTars()

	log.Infof("%s环境初始化完成", config.CommonCfg.RealEnv)
	tars.Run()
}

package config

import (
	"encoding/json"
	"github.com/TarsCloud/TarsGo/tars"
	"io/ioutil"
	"runtime"
)

var (
	CommonCfg        = new(CommonConfig)
	DbCfg            = new(DbConf)
	ConfPath  string = ".\\"
	DirPath   string = "/"
)

type CommonConfig struct {
	*CommonConf
}

type DbConfig struct {
	*DbConf
}

func InitCommonConfig() {
	var path = ""
	var err error
	if runtime.GOOS == "darwin" {
		path = "./config/"
	} else if runtime.GOOS == "windows" {
		path = ConfPath
	} else {
		path = tars.GetServerConfig().BasePath
	}

	data, err := ioutil.ReadFile(path + "/common.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, CommonCfg)
	if err != nil {
		panic(err)
	}
}

func InitDBConfig() {
	// config
	var path = ""
	var err error
	if runtime.GOOS == "darwin" {
		path = "./config/"
	} else if runtime.GOOS == "windows" {
		path = ConfPath
	} else {
		path = tars.GetServerConfig().BasePath
	}

	dataDb, err := ioutil.ReadFile(path + "db.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(dataDb, DbCfg)
	if err != nil {
		panic(err)
	}
}

func InitCommonConfigLocal(localFile string) {
	data, err := ioutil.ReadFile(localFile)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, CommonCfg)
	if err != nil {
		panic(err)
	}
}

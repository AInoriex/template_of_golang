package rpc

import (
	"github.com/TarsCloud/TarsGo/tars"
)

var (
	Suffix = ""

	// Must: Global Init Once
	Comm *tars.Communicator
)

// @Title   初始化Tars服务
// @Description Must: Global Init Once
// @Author  wzj  (2022/7/15 14:45)
func InitCommTars() {
	if Comm == nil {
		Comm = tars.NewCommunicator()
	}
}

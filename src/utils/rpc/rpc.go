package rpc

import (
	"context"
	"singapore/src/utils/errors"
	"singapore/src/utils/log"
	"singapore/src/utils/rpc/define"
	"strconv"

	"github.com/TarsCloud/TarsGo/tars/util/current"
	"go.uber.org/zap"
)

// 定义基本obj_name
const (
	LogicObjName string = "magic.logic.LogicObj"
	UserObjName  string = "magic.user.UserObj"
	RoomObjName  string = "magic.room.RoomObj"
	OrderObjName string = "magic.order.OrderObj"
)

// rpc context 信息
type ContextData struct {
	// 房间ID
	RoomId string
	// 玩家ID
	PlayerId string
	// 应用ID
	AppId int64
}

// @Title   获取context
// @Description 解析context的内容
// @Author  wzj  (2022/7/29 11:06)
func GetContextData(ctx context.Context) (data *ContextData, err error) {
	data = &ContextData{}
	reqContext, ok := current.GetRequestContext(ctx)
	if !ok {
		log.Error("GetContextData fail GetRequestContext not ok")
		return
	}

	//log.Debug("Get context from context: %v", zap.Any("reqContext", reqContext))

	if rid, ok := reqContext[define.SubscribeRoom]; ok {
		data.RoomId = rid
	}
	if uid, ok := reqContext[define.UID]; ok {
		data.PlayerId = uid
	}
	if appid, ok := reqContext[define.AppID]; ok {
		data.AppId, _ = strconv.ParseInt(appid, 10, 64)
	}

	if data.AppId <= 0 {
		log.Error("rpc.GetContextData appid is empty", zap.Any("reqContext", reqContext))
		return data, errors.ErrRelogin
	}

	if data.RoomId == "" {
		log.Error("rpc.GetContextData RoomId is empty", zap.Any("reqContext", reqContext))
		return data, errors.ErrRelogin
	}

	return
}

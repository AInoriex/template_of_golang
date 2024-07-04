package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"singapore/src/utils/alarm"
	"singapore/src/utils/log"
	// "encoding/json"
	// "singapore/src/server/router/model"
	// "singapore/src/utils/config"
	// uerrors "singapore/src/utils/errors"
)

// 测试接口
func HelloWorld(c *gin.Context) {
	// var err error
	req := GetGinBody(c)
	// var attrMap = make(map[string]interface{})
	dataMap := make(map[string]interface{})

	log.Info("DeployReleaseService params", zap.String("req", string(req)))
	// reqbody := &model.DeployReleaseServiceReq{}
	// err = json.Unmarshal(req, &reqbody)
	// if err != nil {
	// 	log.Error("DeployReleaseService unmarshal fail", zap.Error(err), zap.String("req", string(req)))
	// 	Fail(c, uerrors.Parse(uerrors.ErrParam.Error()).Code, uerrors.Parse(uerrors.ErrParam.Error()).Detail)
	// 	return
	// }
	// if reqbody.AnchorId <= 0 {
	// 	log.Error("DeployReleaseService params error", zap.Any("reqbody", reqbody))
	// 	Fail(c, uerrors.Parse(uerrors.ErrParam.Error()).Code, uerrors.Parse(uerrors.ErrParam.Error()).Detail)
	// 	return
	// }

	// 返回数据
	dataMap["res"] = "Hello world."
	Alarm2LarkLocal(alarm.LarkServerAlarmTextVariable{
		Level:  "DEBUG",
		Title:  "TEST DEMO",
		Msg:    "Hello world, this is test message.",
		Detail: fmt.Sprintf("req:%s", string(req)),
	})

	Success(c, dataMap)
}

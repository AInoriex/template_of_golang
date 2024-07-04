package handler

import (
	// "singapore/src/utils/config"
	"singapore/src/utils/alarm"
	"singapore/src/utils/log"
	"singapore/src/utils/net"
	"fmt"
	"go.uber.org/zap"
	"time"
)

// 飞书通知
func Alarm2LarkLocal(v alarm.LarkServerAlarmTextVariable) {
	// if config.CommonCfg.Env == "local" {
	// 	return
	// }
	var err error
	v.Ip, _ = net.GetIpAddress()
	v.Time = time.Now()
	msg := fmt.Sprintf(alarm.TrainDubbingServiceTemplate, v.Title, v.Msg, v.Ip, v.Time, v.Detail)
	err = alarm.PostFeiShu(v.Level, msg, alarm.LaciaBot)
	if err != nil {
		log.Error("Alarm2LarkLocal 告警失败", zap.Error(err))
	}
}

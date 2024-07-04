package alarm

import (
	"time"
)

var (
	AlarmLevelPanic string = "PANIC" // 飞书告警-崩溃级别
	AlarmLevelError string = "ERROR" // 飞书告警-错误级别
	AlarmLevelWarn  string = "WARN"  // 飞书告警-提醒级别
	AlarmLevelInfo  string = "INFO"  // 飞书告警-信息级别
)

// 普通文本信息
type LarkAlarmTextVariable struct {
	MsgType string `json:"msg_type"`
	Content struct {
		Text string `json:"text"`
	} `json:"content"`
}

// 服务告警文本模板
type LarkServerAlarmTextVariable struct {
	Level  string    `json:"level"`
	Title  string    `json:"title"`
	Msg    string    `json:"msg"`
	Ip     string    `json:"ip"`
	Detail string    `json:"detail"`
	Time   time.Time `json:"time"`
}

// 通知文本模板
var (
	TrainUserSubmitTaskTextTemplate string = "用户提交了新的训练任务，请审核"
	TrainUserAuditTaskTextTemplate  string = "用户审核了任务，状态:%v，留言:%s"
	TrainDubbingServiceTemplate     string = "【%s】 \n告警信息\t\t%v \n告警IP\t\t%v \n告警时间\t\t%v \n告警详情\t\t%v"
)

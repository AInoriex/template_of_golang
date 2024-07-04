package alarm

import (
	"testing"
	"time"
)

func TestRobot(t *testing.T) {
	t.Log("TestRobot Start.")

	cardVal := TrainCardVariable{
		Appid:        2,
		TrainService: "127.0.0.1:80",
		Name:         "P·A·I·M·O·N",
		TaskId:       12,
		Content:      "TEST TEST TEST\n测试信息",
		Cost:         "1小时",
		CreateTime:   time.Now(),
	}
	PostCard2FeiShuV2(LaciaBot, LarkCardTemplateIdTrain, cardVal)

	t.Log("TestRobot Done.")
}

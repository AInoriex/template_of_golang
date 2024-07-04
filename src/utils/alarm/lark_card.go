package alarm

import "time"

/*	@Title	飞书卡片使用说明
	@Date	24.03.14 18:00
	@Author	Alex.Xiang
	@Desc	该go struct根据飞书提供的卡片调用结构而制定

	 EXP. POST CARD JSON
	{
		"msg_type": "interactive",
		"card": "{\"type\": \"template\", \"data\": { \"template_id\": \"ctp_AAm2uLFBPpiw\", \"template_variable\": {\"total_count\":\"1小时23分钟\",\"uuid\":\"b320a1aae3657d17c854a515bbf2\",\"create_time\":\"2023-08-23 21:32:22\",\"name\":\"yan_xiaoyin001_female\",\"content\":\"【xxx】成功，开始执行下一步（进度：1 / 4）\\n失败原因：xxx\"} } }"
	}

	{
		"type": "template",
		"data": {
			"template_id": "ctp_2uLFBPpiw",
			"template_variable": {
				"total_count": "1小时23分钟",
				"uuid": "b320a1aae3657d17c854a515bbf2",
				"create_time": "2023-08-23 21:32:22",
				"name": "yan_xiaoyin001_female",
				"content": "【xxx】成功，开始执行下一步（进度：1 / 4）\n失败原因：xxx"
			}
		}
	}
**/

// 卡片信息
type FeiShuCardMsg struct {
	MsgType string `json:"msg_type"`
	Card    string `json:"card"`
}

// 训练平台训练通知卡片
// 预览: https://open.feishu.cn/tool/cardbuilder?templateId=ctp_AAm2uLFBPpiw
const LarkCardTemplateIdTrain string = "ctp_AAm2uLFBPpiw" // 训练卡片

type TrainCardContent struct {
	Type string        `json:"type"`
	Data TrainCardData `json:"data"`
}
type TrainCardData struct {
	Id       string            `json:"template_id"`
	Variable TrainCardVariable `json:"template_variable"`
}
type TrainCardVariable struct {
	Appid        int64     `json:"appid"`
	TrainService string    `json:"train_service"`
	Name         string    `json:"name"`
	TaskId       int64     `json:"task_id"`
	Content      string    `json:"content"`
	Cost         string    `json:"cost"`
	CreateTime   time.Time `json:"create_time"`
}

// 训练平台用户通知卡片
// 预览: https://open.feishu.cn/tool/cardbuilder?templateId=ctp_AA8l4yD6rWjt
const LarkCardTemplateIdTrainUser string = "ctp_AA8l4yD6rWjt" // 用户卡片

type TrainUserCardContent struct {
	Type string            `json:"type"`
	Data TrainUserCardData `json:"data"`
}
type TrainUserCardData struct {
	Id       string                `json:"template_id"`
	Variable TrainUserCardVariable `json:"template_variable"`
}
type TrainUserCardVariable struct {
	AppUserId  string    `json:"app_user_id"`
	AnchorName string    `json:"anchor_name"`
	TaskId     string    `json:"task_id"`
	CreateTime time.Time `json:"create_time"`
	Content    string    `json:"content"`
}

// 4变量通知卡片
// 预览: https://open.feishu.cn/tool/cardbuilder?templateId=ctp_AAUwci9YIeFx
const LarkCardTemplateIdCommon4 string = "ctp_AAUwci9YIeFx" // 4变量卡片

type TranslateCardContent struct {
	Type string          `json:"type"`
	Data Common4CardData `json:"data"`
}
type Common4CardData struct {
	Id       string      `json:"template_id"`
	Variable interface{} `json:"template_variable"`
}
type Common4CardVariable struct {
	Title     string `json:"title"`
	Key1      string `json:"key_1"`
	Val1      string `json:"val_1"`
	Key2      string `json:"key_2"`
	Val2      string `json:"val_2"`
	Key3      string `json:"key_3"`
	Val3      string `json:"val_3"`
	Key4      string `json:"key_4"`
	Val4      string `json:"val_4"`
	ExtraKey  string `json:"extra_key"`
	ExtraVal  string `json:"extra_val"`
	LinkTitle string `json:"link_title"`
	LinkUrl   string `json:"link_url"`
}

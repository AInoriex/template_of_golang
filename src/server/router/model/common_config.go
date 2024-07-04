package model

import "time"

const (
	CcKeyTrainSvcTag                 = "train_service_tag"                  //数字人训练节点标签 beta
	CcKeyTrainSvcRecoverTag          = "train_service_recover_tag"          //数字人训练节点-修复版 dev
	CcKeyTrainDevUid                 = "train_dev_uid"                      //训练开发用户（提测用）
	CcKeyTrainDeployLxProdUser       = "train_deploy_to_lx_prod_user"       //训练平台部署灵犀正式用户用
	CcKeyUpgradeFaceswapTaskIds      = "upgrade_faceswap_task_ids"          //任务升级需要升级换脸逗号隔开
	CcKey11LabsSubscriptionUsed      = "elevenlabs_subscription_used"       //elevenlabs订阅使用量
	CcKey11LabsSubscriptionResetTime = "elevenlabs_subscription_reset_time" //elevenlabs订阅重置时间
)

// 公共配置，获取配置概率
type CommonConfig struct {
	ConfigKey  string    `gorm:"primary_key;column:config_key" json:"config_key"` // 配置key
	Name       string    `gorm:"column:name" json:"name"`                         // 配置名称
	Value      string    `gorm:"column:value" json:"value"`                       // 配置
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`           // 创建时间
}

func (this *CommonConfig) TableName() string {
	return "common_config"
}

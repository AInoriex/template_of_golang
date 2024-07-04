package errors

// 模型训练
//说明：
//      错误码主要用于服务返回，方便客户端展示错误信息以及后续的统计
//      每个服务对应一个错误文件 分两部分，第一部分为错误码，
//      第二部分为错误码对应的错误（ps：错误信息需要动态生成的可不放在这里，但是错误码要存放在这里）
//命名格式：
//      错误码以ErrorCode开头，紧接服务名，最后加上错误描述
//      错误信息以Error开头，紧接服务名，最后加上错误描述
//IM错误码 [25000, 26000)

const (
	ErrorCodeTranslateTask             int32 = 25000
	ErrorCodeTranslateOutOfRetry       int32 = 25001
	ErrorCodeTranslateOutOfTime        int32 = 25002
	ErrorCodeTranslateTaskInvalid      int32 = 25003
	ErrorCodeTranslateTaskExtraInvalid int32 = 25004

	ErrorCodeGetDubbingToken       int32 = 25100
	ErrorCodeCreateDubbingTask     int32 = 25101
	ErrorCodeDownloadDubbingResult int32 = 25102
	ErrorCodeSubmitMergeFailed     int32 = 25103
	ErrorCodeQueryDubbingTask      int32 = 25104
)

var (
	ErrorTranslateTask                  = New("translate", "翻译任务失败", ErrorCodeTranslateTask)
	ErrorTranslateOutOfRetry            = New("translate", "超过重试次数，任务失败", ErrorCodeTranslateOutOfRetry)
	ErrorTranslateOutOfTime             = New("translate", "时间超时，任务失败", ErrorCodeTranslateOutOfTime)
	ErrorTranslateTaskInvalid           = New("translate", "无效的任务 无法提交到流程执行", ErrorCodeTranslateTaskInvalid)
	ErrorTranslateTaskExtraInvalid      = New("translate", "数据库信息错误", ErrorCodeTranslateTaskExtraInvalid)
	ErrorTranslateTaskSubmitMergeFailed = New("translate", "提交视频合成失败", ErrorCodeSubmitMergeFailed)

	ErrorGetDubbingToken       = New("translate", "获取token失败", ErrorCodeGetDubbingToken)
	ErrorCreateDubbingTask     = New("translate", "创建dub任务失败", ErrorCodeCreateDubbingTask)
	ErrorDownloadDubbingResult = New("translate", "获取dub结果失败", ErrorCodeDownloadDubbingResult)
	ErrorQueryDubbingTask      = New("translate", "查询dub任务失败", ErrorCodeQueryDubbingTask)
)

package errors

//说明：
//      错误码主要用于服务返回，方便客户端展示错误信息以及后续的统计
//      每个服务对应一个错误文件 分两部分，第一部分为错误码，
//      第二部分为错误码对应的错误（ps：错误信息需要动态生成的可不放在这里，但是错误码要存放在这里）
//命名格式：
//      错误码以ErrorCode开头，紧接服务名，最后加上错误描述
//      错误信息以Error开头，紧接服务名，最后加上错误描述
// 公共错误定义 [45000, 50000)
const (
	ErrCodeDeployServiceUnknownError = 45000
	ErrCodeDeployCheckFileExistFail  = 45001
	ErrCodeDeployCalFileMd5Fail      = 45002
	ErrCodeDeployActionFail          = 45003
)

var (
	ErrDeployServiceUnknownError = New("", "部署服务发生未知错误", ErrCodeDeployServiceUnknownError)
	ErrDeployCheckFileExistFail  = New("", "部署服务检查文件有效失败", ErrCodeDeployCheckFileExistFail)
	ErrDeployCalFileMd5Fail      = New("", "部署服务计算文件md5失败", ErrCodeDeployCalFileMd5Fail)
	ErrDeployActionFail          = New("", "部署服务同步文件操作失败", ErrCodeDeployActionFail)
)

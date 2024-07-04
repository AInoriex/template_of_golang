package errors

type CodeMsg struct {
	Code int32
	Msg  string
}

//var Success = CodeMsg{0, "success"}
//var DigitalHumanNotGenerate = CodeMsg{30000, "用户模型未生成"}
//var AiGenerateDigitalHumanFail = CodeMsg{30001, "生成捏脸参数失败"}
//var ArgsFail = CodeMsg{30002, "参数有误"}
//var AppConfigError = CodeMsg{30003, "配置有误"}
//var AccessTokenError = CodeMsg{30004, "access_token有误"}
////var AccessTokenExpireTimeError = CodeMsg{30005, "access_token过期"}
//var AccessTokenExpireTimeError = New("api", "access_token过期", 30005)
var (
	DigitalHumanNotGenerate    = New("api", "用户模型未生成", 30000)
	AiGenerateDigitalHumanFail = New("api", "生成捏脸参数失败", 30001)
	ArgsFail                   = New("api", "参数有误", 30002)
	AppConfigError             = New("api", "配置有误", 30003)
	AccessTokenError           = New("api", "access_token有误", 30004)
	AccessTokenExpireTimeError = New("api", "access_token过期", 30005)
	ErrCacheNotExist           = New("api", "缓存不存在", 30006)

	FrameSequenceNotUpload      = New("api", "用户序列帧未上传", 30007)
	FrameSequenceUploadArgsFail = New("api", "上传参数有误", 30008)
	PushFrameSequenceFail       = New("api", "推送u3d序列帧失败", 30009)
	AuditNotPassFail            = New("api", "内容含敏感信息，请修改后重新提问", 30010)
)

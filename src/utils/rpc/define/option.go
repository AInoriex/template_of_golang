package define

// 这里定义所有上行消息 RPCInput 的可选项。内建的名称，只允许小写(a-z)和连字符(-)
const (
	// 长连接鉴权时使用
	UID   = "uid"
	THUID = "th_uid" // 第三方用户id
	Token = "X-Token"
	AppID = "appid"
	//ConnID = "connid" // 同一uid下connid唯一
	// 长连接鉴权时使用，切换房间时使用
	SubscribeRoom      = "roomid"
	TcpSeq             = "tcp_seq"
	ServerTimeStart    = "server_start_time"
	ServerTimeEnd      = "server_end_time"
	ClientTimeStart    = "client_start_time"
	IsAnonymousUser    = "is-anonymous-user"
	HeartbeatThreshold = "heartbeat-threshold"
	// 连接相关信息
	ClientIP          = "client-ip"
	ClientPort        = "client-port"
	AccessPointIP     = "access-point-ip"
	AccessPointPort   = "access-point-port"
	InvokePointIP     = "invoke-point-ip"
	InvokePointIPInt  = "invoke-point-ip-int"
	WiredLogicID      = "wired-logic-id"
	WiredVersion      = "version"
	WiredProtocolCode = "protocol-code"
	WiredAsyncRsp     = "asyncrsp"
	WiredConnID       = "wired-connid"
	LeaveRoomFlag     = "-1"
	Client            = "client"
	PushCmdID         = "cmdId"
	AsyncPush         = "async"
	UserImKey         = "user-im-key"
	Application       = "application"
	ScreenSize        = "screen_size"
	// res_oss 上传返回文件路径
	FileUrl = "file_url"
	//时间
	RECTIME             = "X-rec-tm"
	PROTIME             = "X-pro-tm"
	ENV                 = "env"
	Channel             = "channel"
	NoAuthId            = "no_auth_id"
	PackageId           = "package_id"
	SkyWalkingTracingId = "skywalking_tracingId"
	NoRoom              = "-1"
)

// span Tag
const (
	SpanTagFuncTyp = "func_type"
)

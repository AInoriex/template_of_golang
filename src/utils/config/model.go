package config

//此文件保存公共的结构

const (
	SmsTplDelAccountAsk = "DelAccountAsk"
)

type Mysql struct {
	Host            string `json:"host"`
	Host_rw         string `json:"host_rw"`         // 读写分离连接
	Host_admin_read string `json:"host_admin_read"` // 读写分离连接
	Db              string `json:"db"`
	MaxCon          int    `json:"max_con"`
}

type Redis struct {
	Host     string `json:"host"`
	Password string `json:"password"`
	Db       int    `json:"db"`
}

type Log struct {
	SavePath string `json:"save_path"`
}

type CommonConf struct {
	AppName             string               `json:"app_name"`
	Env                 string               `json:"env"`      //alpha beta pre prod
	RealEnv             string               `json:"real_env"` //alpha beta pre prod
	Debug               bool                 `json:"debug"`    //alpha beta pre prod
	Log                 Log                  `json:"log"`
	AiConf              *AiConfig            `json:"ai_conf"`
	U3dConf             *U3dConfig           `json:"u3d_conf"`
	UploadConf          *UploadConfig        `json:"upload_conf"`
	OpenDbLog           bool                 `json:"open_db_log"`
	TencentOSS          *TencentOSS          `json:"tencent_oss"`
	TencentOSSX         *TencentOSS          `json:"tencent_oss_x"`
	BylinkConf          *BylinkConf          `json:"bylink_conf"`
	BylinkMaigcConf     *BylinkConf          `json:"bylink_magic_conf"`
	ApiHost             string               `json:"api_host"` // api域名
	GrpcServer          *GrpcServerConf      `json:"grpc_server"`
	AIGrpcConf          *AIGrpcConf          `json:"ai_grpc_conf"`
	AVConf              *AVConf              `json:"av_conf"`
	LiveGrpcConf        *LiveGrpcConf        `json:"live_grpc_conf"`
	LiveExtGrpcConf     *LiveGrpcConf        `json:"live_ext_grpc_conf"`
	LiveGrpcServer      *GrpcServerConf      `json:"live_grpc_server"`
	LiveGrpcServerChat  *GrpcServerConf      `json:"live_grpc_server_chat"`
	MagicReportGrpcConf *MagicReportGrpcConf `json:"magic_report_conf"`
	LxConf              *LxConf              `json:"lx_conf"`
	AiChatConf          *AiChatConf          `json:"ai_chat_conf"`
	HttpServer          *HttpServerConf      `json:"http_server"`
	TranslateConf       *TranslateConf       `json:"translate_conf"`
	LabsOf11DubbingConf *LabsOf11DubbingConf `json:"elevenlabs_dubbing_conf"`

	TarsSuffix string `json:"tars_suffix"`
}

type DbConf struct {
	Mysql    Mysql `json:"mysql"`
	MysqlLog Mysql `json:"mysql_log"`
	Redis    Redis `json:"redis"`
	Oss      Oss   `json:"oss"`
}

type Oss struct {
	Endpoint     string `json:"endpoint"`      //数据中心域名
	AccessKey    string `json:"access_key"`    //校验key
	AccessSecret string `json:"access_secret"` //校验密钥
	Bucket       string `json:"bucket"`        //存储空间名称
	Env          string `json:"env"`           //环境 alpha beta prod
}

// AI配置
type AiConfig struct {
	ApiHost   string `json:"api_host"`
	ApiHostV2 string `json:"api_host_v2"`
	TMeta01   string `json:"t_meta_01"`
}

// u3d配置
type U3dConfig struct {
	ApiHost string `json:"api_host"`
}

// 上传配置配置
type UploadConfig struct {
	CdnHost   string `json:"cdn_host"`
	UploadDir string `json:"upload_dir"`
}

type TencentOSS struct {
	Url       string `json:"url"`
	SecretId  string `json:"secret_id"`
	SecretKey string `json:"secret_key"`
	Env       string `json:"env"`
	Cdn       string `json:"cdn"`
}

// 百灵埋点系统
type BylinkConf struct {
	KafkaNode  []string `json:"kafka_node"`
	KafkaTopic string   `json:"kafka_topic"`
	AppName    string   `json:"app_name"`
}

// grpc服务端
type GrpcServerConf struct {
	Addr    string `json:"addr"`
	Timeout int64  `json:"timeout"`
	Network string `json:"network"`
}

// http服务端
type HttpServerConf struct {
	Addr    string `json:"addr"`
	Timeout int64  `json:"timeout"`
	Network string `json:"network"`
}

// ai直播流客户端
type AIGrpcConf struct {
	Addr string `json:"addr"`
}

// AIgrpc服务
type AVConf struct {
	RpcAddr    string `json:"rpc_addr"`
	UdpAddr    string `json:"udp_addr"`
	LanUdpAddr string `json:"lan_udp_addr"` //内网服务
}

// live直播流客户端
type LiveGrpcConf struct {
	Addr string `json:"addr"`
}

// 数据上报客户端
type MagicReportGrpcConf struct {
	Addr string `json:"addr"`
}

// 灵犀配置
type LxConf struct {
	LxUrlBase string `json:"lx_url_base"`
}

type AiChatConf struct {
	Addr string `json:"addr"`
}

// 翻译任务配置
type TranslateConf struct {
	ExpireSecond int64 `json:"expire"`
	TaskMax      int   `json:"task_max"`
}

// ElevenLabs Dubbing配置
type LabsOf11DubbingConf struct {
	ProxyOn       bool     `json:"proxy_on"`
	ProxyHttp     string   `json:"proxy_http"`
	ApiKey        []string `json:"api_key"`
	AudioDefineOn bool     `json:"audio_refine_on"`
}

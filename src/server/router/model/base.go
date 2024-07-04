package model

// 通用参数
type BaseReq struct {
	Appid     int64  `json:"appid"`       //appid
	Appkey    string `json:"appkey"`      //appkey
	AppUserId string `json:"app_user_id"` //业务方用户uid
	OsType    string `json:"os_type"`     //系统类型（ios，android）
	Version   string `json:"version"`     //SDK版本
	Device    string `json:"device_id"`   //设备
	RequestId string `json:"request_id"`  //请求ID
}

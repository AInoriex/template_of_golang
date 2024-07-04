package model

import "singapore/src/utils/config"

type DubbingCallbackReq struct {
	DubbingId    string `json:"dubbing_id"`
	ErrCode      int64  `json:"err_code"`
	Msg          string `json:"msg"`
	OutVideoPath string `json:"out_video_path"`
	ResourceId   int64  `json:"resource_id"`
}

type TranslateDanmuCallbackReq struct {
	DubbingId    string `json:"dubbing_id"`
	ErrCode      int32  `json:"err_code"`
	Msg          string `json:"msg"`
	OutVideoPath string `json:"out_video_path"`
	ResourceId   int64  `json:"resource_id"`
}

// 合成视频回调地址
func GetDubbingCallbackUrl() string {
	return config.CommonCfg.ApiHost + "/translate_api/dubbing_callback"
}

// 字幕合成回调地址
func GetTranslateDanmuCallbackUrl() string {
	return config.CommonCfg.ApiHost + "/translate_api/translate_danmu_callback"
}

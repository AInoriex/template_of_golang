package common

import (
	// "singapore/src/utils/alarm"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	// "singapore/src/utils/bylink"
	uerrors "singapore/src/utils/errors"
	"singapore/src/utils/log"
	"singapore/src/utils/utime"
	"singapore/src/utils/uuid2"
	"net/http"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

const (
	KeyEventTracking          = "eventTracking"
	KeyEventName              = "eventName"
	KeyEventAttrSdkVersion    = "sdk_version"
	KeyEventAttrRequestId     = "request_id"
	KeyEventAttrApiStampBegin = "api_stamp_begin"
	KeyEventAttrApiStampEnd   = "api_stamp_end"
	KeyEventAttrApiStampCost  = "api_stamp_cost"
	KeyEventAttrApiErrorCode  = "api_error_code"
	KeyEventAttrApiErrorMsg   = "api_error_msg"
	KeyEventAttrApiFunc       = "api_func"
	KeyEventAttrApiHost       = "api_host"
	KeyEventAttrStampBegin    = "stamp_begin"
	KeyEventAttrStampEnd      = "stamp_end"
	KeyEventAttrStampCost     = "stamp_cost"
	KeyEventAttrErrorCode     = "error_code"
	KeyEventAttrErrorMsg      = "error_msg"
	KeyEventAttrUid           = "$uid"
	KeyEventAttrAppid         = "$app_id"
)

// 页面版本号
func GetLiveWebVersion() string {

	return "1.11.0"
}

// 响应结构体
type Response struct {
	ErrorCode int32       `json:"code"` // 自定义错误码
	Data      interface{} `json:"data"` // 数据
	Message   string      `json:"msg"`  // 信息
}

// Success 响应成功 ErrorCode 为 0 表示成功
func Success(c *gin.Context, data interface{}) {
	c.Header("Live-Web-Version", GetLiveWebVersion())
	c.JSON(http.StatusOK, Response{
		0,
		data,
		"ok",
	})
}

// Fail 响应失败 ErrorCode 不为 0 表示失败
func Fail(c *gin.Context, errorCode int32, msg string) {
	c.Header("Live-Web-Version", GetLiveWebVersion())
	c.JSON(http.StatusOK, Response{
		errorCode,
		struct{}{},
		msg,
	})
}

// Fail 响应失败 ErrorCode 不为 0 表示失败
func Fail2(c *gin.Context, errorCode int32, msg string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		errorCode,
		data,
		msg,
	})
}

// Success 响应成功 ErrorCode 为 0 表示成功
//func SuccessTrack(c *gin.Context, data interface{}, attrMap map[string]interface{}) {
//	attrMap[KeyEventAttrErrorCode] = 0
//	c.JSON(http.StatusOK, Response{
//		0,
//		data,
//		"ok",
//	})
//}

// Fail 响应失败 ErrorCode 不为 0 表示失败
func FailTrack(c *gin.Context, errorCode int32, msg string, attrMap map[string]interface{}) {
	c.Header("Live-Web-Version", GetLiveWebVersion())
	attrMap[KeyEventAttrErrorCode] = errorCode
	attrMap[KeyEventAttrErrorMsg] = msg
	c.JSON(http.StatusOK, Response{
		errorCode,
		struct{}{},
		msg,
	})
}

// Fail 响应失败 ErrorCode 不为 0 表示失败
func FailTrack2(c *gin.Context, errorCode int32, msg string, attrMap map[string]interface{}, dataMap map[string]interface{}) {
	c.Header("Live-Web-Version", GetLiveWebVersion())
	attrMap[KeyEventAttrErrorCode] = errorCode
	attrMap[KeyEventAttrErrorMsg] = msg
	c.JSON(http.StatusOK, Response{
		errorCode,
		dataMap,
		msg,
	})
}

// @Title   自定义Recovery
// @Description 上报告警,通过飞书alarm_id 查找日志信息
// @Author  wzj  (2022/12/7 17:03)
func Recovery() func(*gin.Context) {
	return gin.CustomRecovery(func(c *gin.Context, err interface{}) {
		//log.Fatal("panic:", zap.String("path", c.Request.URL.Path), zap.Any("err", err))
		uuid := uuid.GetUuid()
		stack := PanicTrace(4)
		log.Error("Recovery Panic: "+err.(string), zap.String("path", c.Request.URL.Path), zap.String("alarm_id", uuid), zap.String("stack", string(stack)))
		// 当前时间
		currentTime := time.Now().Format("2006-01-02 15:04:05")

		// 定义 文件名、行号、方法名
		fileName, line, functionName := "?", 0, "?"

		pc, fileName, line, ok := runtime.Caller(3)
		if ok {
			functionName = runtime.FuncForPC(pc).Name()
			functionName = filepath.Ext(functionName)
			functionName = strings.TrimPrefix(functionName, ".")
		}
		type errorInfo struct {
			Alarm    string `json:"level"`
			Time     string `json:"time"`
			Message  string `json:"message"`
			Filename string `json:"filename"`
			Line     int    `json:"line"`
			Funcname string `json:"func"`
			Path     string `json:"path"`
			Host     string `json:"host"`
			AlarmId  string `json:"alarm_id"`
		}
		var msg = errorInfo{
			Alarm:    "PANIC",
			Time:     currentTime,
			Message:  err.(string),
			Filename: fileName,
			Line:     line,
			Funcname: functionName,
			Path:     c.Request.URL.Path,
			Host:     c.Request.Host,
			AlarmId:  uuid,
		}

		jsons, errs := json.Marshal(msg)
		if errs != nil {
			fmt.Println("json marshal error:", errs)
		}
		log.Errorf("Recovery Alarm:%s", string(jsons))
		// alarm.AlarmPanic("MagicServer Live服务发生PANIC告警:\n" + string(jsons))

		c.String(http.StatusInternalServerError, "Tmeta: Internal Server Error")
	})
}

// @Title   自定义日志中间件
// @Description defer 上报埋点信息（需按固定格式透传）
// @Author  wzj  (2022/12/7 17:02)
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		startM := utime.GetNowUnixM()
		path := c.Request.URL.Path
		method := c.Request.Method
		clientIP := c.ClientIP()
		log.Infof("begin server | %s  %s | %15s |", method, path, clientIP)

		// body
		c.Next()

		end := time.Now()
		//执行时间
		latency := end.Sub(start)
		statusCode := c.Writer.Status()
		log.Infof("end server | %3d | %13v | %15s | %s  %s |",
			statusCode,
			latency,
			clientIP,
			method, path,
		)
		if evt, ok := c.Get(KeyEventTracking); ok && evt != nil {
			if attrMap, ok2 := evt.(map[string]interface{}); ok2 {
				attrMap[KeyEventAttrStampBegin] = startM
				if ename, ok3 := attrMap[KeyEventName]; ok3 && ename != "" {
					DeferEvent(attrMap)
				}
			}
		}

	}
}

// @Title  Gin中间件上报埋点
// @Description desc
// @Author  wzj  (2022/12/7 17:01)
func DeferEvent(attrMap map[string]interface{}) (err error) {
	//defer common.CheckGoPanic()
	// var distinctId string
	var stime, etime, apiSTime, apiETime int64
	var eventBylink string
	if _, ok := attrMap[KeyEventName]; ok {
		eventBylink = attrMap[KeyEventName].(string)
		delete(attrMap, KeyEventName)
	}
	if eventBylink == "" {
		log.Error("DeferHandler report bylink fail", zap.String("event", eventBylink))
		return
	}

	// if id, ok := attrMap["$uid"].(string); ok && id != "" {
	// 	distinctId = id
	// } else if reqId, ok := attrMap["request_id"].(string); ok && distinctId == "" {
	// 	distinctId = reqId
	// }

	//distinctId := attrMap["$uid"].(string)
	attrMap["stamp_end"] = utime.GetNowUnixM()
	if t, ok := attrMap["api_stamp_begin"].(int64); ok {
		apiSTime = t
	}
	if t, ok := attrMap["api_stamp_end"].(int64); ok {
		apiETime = t
	}
	if t, ok := attrMap["stamp_begin"].(int64); ok {
		stime = t
	}
	if t, ok := attrMap["stamp_end"].(int64); ok {
		etime = t
	}
	attrMap["api_stamp_cost"] = apiETime - apiSTime
	attrMap["stamp_cost"] = etime - stime

	// bylink.Report(eventBylink, distinctId, attrMap)
	return
}

// @Title  获取请求参数
// @Description desc
// @Author  wzj  (2022/12/7 17:01)
func GetGinBody(c *gin.Context) (req []byte) {
	req, err := c.GetRawData()
	if err != nil {
		log.Error("GetRequestBody fail", zap.Error(err))
		return
	}
	return
}

// @Title   json请求体转换为埋点信息
// @Description desc
// @Author  wzj  (2022/12/5 18:18)
func ReqToEvent(req []byte, attrMap map[string]interface{}) map[string]interface{} {
	s := struct {
		Appid     int64  `json:"appid"`       //appid
		AppUserId string `json:"app_user_id"` //业务方用户uid
		OsType    string `json:"os_type"`     //系统类型（ios，android）
		Version   string `json:"version"`     //SDK版本
		RequestId string `json:"request_id"`  //请求ID
		Device    string `json:"device"`      //设备ID
		DeviceId  string `json:"device_id"`   //设备ID
	}{}

	err := json.Unmarshal(req, &s)
	if err != nil {
		return attrMap
	}

	attrMap["$uid"] = s.AppUserId
	attrMap["$app_id"] = s.Appid
	attrMap["request_id"] = s.RequestId
	attrMap["sdk_version"] = s.Version
	attrMap["$os"] = s.OsType
	if s.Device != "" {
		attrMap["$device_id"] = s.Device
	} else {
		attrMap["$device_id"] = s.DeviceId
	}

	return attrMap
}

// @Title   error转化埋点信息
// @Description desc
// @Author  wzj  (2022/12/5 18:18)
func ErrorToEvent(err error, attrMap map[string]interface{}) map[string]interface{} {
	if e := uerrors.Parse(err.Error()); e != nil && e.Code != uerrors.Failure {
		attrMap[KeyEventAttrApiErrorCode] = e.Code
		attrMap[KeyEventAttrApiErrorMsg] = e.Detail
	}
	return attrMap
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			//接收客户端发送的origin （重要！）
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			//服务器支持的所有跨域请求的方法
			//c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			////允许跨域设置可以返回其他子段，可以自定义字段
			//c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session")
			//// 允许浏览器（客户端）可以解析的头部 （重要）
			//c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
			////设置缓存时间
			//c.Header("Access-Control-Max-Age", "172800")
			////允许客户端传递校验信息比如 cookie (重要)
			//c.Header("Access-Control-Allow-Credentials", "true")
		}

		//允许类型校验
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "ok!")
		}

		defer func() {
			if err := recover(); err != nil {
				log.Errorf("Panic info is: %v", err)
			}
		}()

		c.Next()
	}
}

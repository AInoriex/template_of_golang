package common

import (
	"bytes"
	"compress/gzip"
	"singapore/src/utils/config"
	"singapore/src/utils/log"
	"encoding/base64"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"os"
	"regexp"
	"runtime"
	"strings"
	"time"
	"io"
)

const (
	AppIdDevDemo  = 2
	AppIdProdDemo = 2
)

// @Title   检查Panic
// @Description 记录日志并上报Panic
// @Author  wzj  (2022/7/15 14:41)
func CheckGoPanic() {
	if e := recover(); e != nil {
		stack := PanicTrace(4)
		fmt.Println("panic:", e, time.Now(), "\n", string(stack))
		log.Fatal("PanicTrace: \n  %s", zap.Any("recover", e), zap.String("stack", string(stack)))
	}
}

// @Title   获取Demo Appid
// @Description desc
// @Author  wzj  (2022/8/24 10:17)
func GetAppIdDemo(debug bool) int64 {
	appId := AppIdDevDemo
	if !debug {
		appId = AppIdProdDemo
	}
	return int64(appId)
}

// @Title   校验房间id
// @Description  "0" "" "-1" 返回true
// @Author  wzj  (2022/8/26 17:24)
func IsEmptyRoomId(rid string) bool {
	return rid == "" || rid == "-1" || rid == "0"
}

// 上述的结果并不包含头部分，所以得自己加一个，比如: data:image/png;base64,
func ImageToBase64(file string) string {
	//f, err := os.Open("ubuntu.png")
	f, err := os.Open(file)
	if err != nil {
		log.Error("打开图片失败", zap.Any("file", file), zap.Any("err", err))
	}
	all, _ := io.ReadAll(f)
	base64 := base64.StdEncoding.EncodeToString(all)
	// log.Debug("打开图片", zap.Any("file", file), zap.Any("base64", base64))
	return base64
}

// http读取gzip
func ReadGzip(resp *http.Response) (res []byte, err error) {
	// 兼容gzip
	// 是否有 gzip
	gzipFlag := false
	for k, v := range resp.Header {
		if strings.ToLower(k) == "content-encoding" && strings.ToLower(v[0]) == "gzip" {
			gzipFlag = true
		}
	}

	if gzipFlag {
		// 创建 gzip.Reader
		gr, err := gzip.NewReader(resp.Body)
		if err != nil {
			fmt.Println(err.Error())
		}
		defer gr.Close()
		res, _ = io.ReadAll(gr)
	} else {
		res, err = io.ReadAll(resp.Body)
	}
	return res, nil
}

func PanicTrace(kb int) []byte {
	s := []byte("/src/runtime/panic.go")
	e := []byte("\ngoroutine ")
	line := []byte("\n")
	stack := make([]byte, kb<<10) //4KB
	length := runtime.Stack(stack, true)
	start := bytes.Index(stack, s)
	stack = stack[start:length]
	start = bytes.Index(stack, line) + 1
	stack = stack[start:]
	end := bytes.LastIndex(stack, line)
	if end != -1 {
		stack = stack[:end]
	}
	end = bytes.Index(stack, e)
	if end != -1 {
		stack = stack[:end]
	}
	stack = bytes.TrimRight(stack, "\n")

	return stack
}

func CheckPhone(phone string) bool {
	//var res bool
	var rule string = "^1[3456789][0-9]{9}$" //"^[0-9]+$"
	//手机号格式校验
	var reg *regexp.Regexp
	reg, _ = regexp.Compile(rule)
	if reg == nil {
		return false
	}
	if !reg.MatchString(phone) {
		return false
	}
	return true
}

func CheckPhoneVerifyCode(code string) bool {
	//var res bool
	var rule string = "[0-9]{6}$" //"^[0-9]+$"
	//手机号格式校验
	var reg *regexp.Regexp
	reg, _ = regexp.Compile(rule)
	if reg == nil {
		return false
	}
	if !reg.MatchString(code) {
		return false
	}
	return true
}

func GetAppName(appid int64) string {
	m := make(map[int64]string)
	m[1] = "kafu"
	m[2] = "tmeta"
	m[6] = "exhibition"
	m[7] = "wechat"

	m[10001] = "kafu"
	m[10003] = "tmeta"
	m[10006] = "exhibition"
	m[10007] = "wechat"
	if _, ok := m[appid]; ok {
		return m[appid]
	}
	return ""
}

// 判断prod环境
func IsProductEnv() bool {
	return config.CommonCfg.RealEnv == "prod"
}

// 判断beta环境
func IsBetaEnv() bool {
	return config.CommonCfg.RealEnv == "beta"
}

// 判断pre环境
func IsPreEnv() bool {
	return config.CommonCfg.RealEnv == "pre"
}

package common

import (
	"encoding/base64"
	"errors"
	"go.uber.org/zap"
	"io/ioutil"
	"os"
	"singapore/src/utils/log"
	"sort"
	"strconv"
	"strings"
	"time"
)

// 字符串转int64
func StringToInt32NotErr(numString string) int32 {
	num := StringToInt64NotErr(numString)
	return int32(num)
}

// 字符串转int64
func StringToInt64NotErr(numString string) int64 {
	num, err := strconv.ParseInt(numString, 10, 64)
	if err != nil {
		log.Error("字符串转int64", zap.Any("numString", numString))
		return 0
	}
	return num
}

// int32转字符串
func Int32ToString(num int32) string {
	return Int64ToString(int64(num))
}

// int64转字符串
func Int64ToString(num int64) string {
	numString := strconv.FormatInt(num, 10)
	return numString
}

// 重新包装code
func RespCodeToInt32(code interface{}) int32 {
	switch code.(type) {
	case int:
		return int32(code.(int))
	case string:
		data, err := strconv.ParseInt(code.(string), 10, 32)
		if err == nil {
			return int32(data)
		}
		return int32(-1)
	case int32:
		return int32(code.(int32))
	case int64:
		return int32(code.(int64))
	case float64:
		return int32(code.(float64))
	case float32:
		return int32(code.(float32))
	default:
		log.Error("is an unknown type.", zap.Any("code", code))
		return int32(-1)
	}
}

func StringInSlice(target string, strArray []string) bool {
	sort.Strings(strArray)
	index := sort.SearchStrings(strArray, target)
	//index的取值：[0,len(str_array)]
	if index < len(strArray) && strArray[index] == target { //需要注意此处的判断，先判断 &&左侧的条件，如果不满足则结束此处判断，不会再进行右侧的判断
		return true
	}
	return false
}

// 保存本地
func SaveFileByBase64(imageBase64 string, basePath string, fileName string) error {
	if imageBase64 == "" || basePath == "" || fileName == "" {
		log.Error("SaveFileByBase64 fail", zap.Any("path", basePath+fileName), zap.Any(
			"imageBase64", imageBase64))
		return errors.New("param error")
	}

	if !strings.HasSuffix(basePath, "/") {
		basePath = basePath + "/"
	}
	path := basePath + fileName

	imageBytes, _ := base64.StdEncoding.DecodeString(imageBase64) //成图片文件并把文件写入到buffer
	err := ioutil.WriteFile(basePath+fileName, imageBytes, 0666)
	if err != nil {
		log.Error("SaveFileByBase64 fail", zap.Any("path", path), zap.Any("err", err))
	}
	log.Info("SaveFileByBase64 success", zap.Any("path", path), zap.Any("err", err))
	return err
}

// 检查本地文件是否存在
func CheckFileExist(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsExist(err)
}

// 从test.txt中读取base64字符串，解码，然后生成文件
func Base64ToFile(f string) (decodeData []byte, err error) {
	decodeData, err = base64.StdEncoding.DecodeString(string(f))
	if err != nil {
		return
	}
	return decodeData, nil
}

// 去掉base64文件头
func ReplaceBase64Header(srcImage string) string {
	headerList := []string{"data:image/jpeg;base64,", "data:image/png;base64,", "data:image/svg;base64,"}
	for _, pre := range headerList {
		srcImage = strings.Replace(srcImage, pre, "", 1)
	}

	return srcImage
}

func Time2Unix(time2 time.Time) int64 {
	if time2.Unix() <= 0 {
		return 0
	}
	return time2.Unix()
}

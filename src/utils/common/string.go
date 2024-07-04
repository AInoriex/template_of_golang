package common

import (
	"singapore/src/utils/log"
	"go.uber.org/zap"
	"strings"
)

// split 多条件分割 举例 splitStrings:[]rune{',', '.', '?', '!', '，', '。', '？', '！'}
func SplitString(s string, splitStrings []rune) []string {
	Split := func(r rune) bool {
		for _, v := range splitStrings {
			if v == r {
				return true
			}
		}
		return false
	}
	a := strings.FieldsFunc(s, Split)
	return a
}

// 获取播放数据
func GetPlayText(text string, PlayNumSum, PlayLenSum, PlayNumMax int32) (bool, string, string, int32, int32) {
	zh := int32(3)
	textLen := int32(len(text))
	if textLen-PlayLenSum > PlayNumMax*zh {
		textPlay := string([]rune(text)[PlayNumSum : PlayNumSum+PlayNumMax])
		PlayNumSum = PlayNumSum + PlayNumMax
		PlayLenSum = PlayLenSum + int32(len(textPlay))
		textLeft := string([]rune(text)[PlayNumMax:])
		return false, textPlay, textLeft, PlayNumSum, PlayLenSum
	} else {
		textPlay := string([]rune(text)[PlayNumSum:])
		PlayNumSum = PlayNumSum + PlayNumMax
		PlayLenSum = PlayLenSum + int32(len(textPlay))
		return true, textPlay, "", PlayNumSum, PlayLenSum
	}
}

// 字符串分割
func SplitSegmentingChat(msg string, isMerge bool) ([]string, string) {
	// 提取断句
	msgList := make([]string, 0)
	msgTemp := ""
	//lenNum := len(msg)
	//fmt.Println("lenNum:", lenNum)
	numAdd := 0
	if IsContainSegmentingEnd(msg) {
		msgList = append(msgList, msg)
	} else {
		for _, v := range msg {
			sv := string(v)
			if IsSegmenting(sv) {
				msgTemp = msgTemp + sv
				msgList = append(msgList, msgTemp)
				msgTemp = ""
				numAdd = 0
			} else if numAdd >= 6 && IsSegmenting2(sv) {
				msgTemp = msgTemp + sv
				msgList = append(msgList, msgTemp)
				msgTemp = ""
				numAdd = 0
			} else {
				numAdd = numAdd + 1
				msgTemp = msgTemp + sv
			}
		}
		if len(msgList) >= 1 {
			lenList := 0
			for _, v := range msgList {
				lenList = lenList + len(v)
			}
			if lenList < 3*6+1 {
				// 少于5个字先不断句
				return make([]string, 0), msg
			}
		}

	}
	if isMerge && len(msgList) > 1 {
		msgMerge := ""
		for _, v := range msgList {
			msgMerge = msgMerge + v
		}
		msgMergeList := make([]string, 0)
		msgMergeList = append(msgMergeList, msgMerge)
		log.Info("合并语句", zap.Any("msgMergeList", msgMergeList))
		return msgMergeList, msgTemp
	}
	return msgList, msgTemp
}

// 是否符合断句
func IsSegmenting(msg string) bool {
	if msg == "；" || msg == "。" || msg == "！" || msg == "？" || msg == "\n" {
		return true
	} else {
		return false
	}
}
func IsSegmenting2(msg string) bool {
	if IsSegmenting(msg) {
		return true
	} else if msg == "," || msg == "，" || msg == "." {
		return true
	} else {
		return false
	}
}

// 是否包含结束标识
func IsContainSegmentingEnd(msg string) bool {
	return strings.Contains(msg, SegmentingEnd()) 
}

// 答案结束标识
func SegmentingEnd() string {
	return "end_end_end"
}

// http chunks 结束标识
func ChunksEnd() string {
	return "eof_eof_eof"
}

// 是否是Unicode空字符串
func IsSpace(msg string) bool {
	if len(msg) == 4 {
		msgList := []rune(msg)
		if msgList[0] == 8203 {
			return true
		}
	}
	return false
}

// 字符串切片去重
func RemoveDuplicates(elements []string) []string {
	// 使用map来存储不重复的元素
	encountered := map[string]bool{}
	result := []string{}

	for i := range elements {
		if encountered[elements[i]] {
			// 做Nothing，元素已经存在.
		} else {
			// 如果元素不存在于map则将其添加到map和结果切片
			encountered[elements[i]] = true
			result = append(result, elements[i])
		}
	}
	return result
}

// 弹幕字符标识
func Danmu() string {
	return "danmu"
}

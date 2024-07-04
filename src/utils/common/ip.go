package common

import (
	"singapore/src/utils/log"
	"go.uber.org/zap"
	"strings"
)

func GetPortByIpPort(ipPort string) int64 {
	list := strings.Split(ipPort, ":")
	if len(list) == 2 {
		return StringToInt64NotErr(list[1])
	} else {
		log.Error("GetPortByIpPort", zap.Any("ipPort", ipPort))
		return 0
	}
}

func GetPortByIpPortString(ipPort string) string {
	list := strings.Split(ipPort, ":")
	if len(list) == 2 {
		return list[1]
	} else {
		log.Error("GetPortByIpPort", zap.Any("ipPort", ipPort))
		return ""
	}
}

func GetHostByIpHost(ipPort string) string {
	list := strings.Split(ipPort, ":")
	if len(list) == 2 {
		return list[0]
	} else {
		log.Error("GetPortByIpPort", zap.Any("ipPort", ipPort))
		return ""
	}

}

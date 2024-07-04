package net

import (
	// "singapore/model"
	// "singapore/util"
	"go.uber.org/zap"
	"net"
	"singapore/src/utils/log"
	"strings"
)

var (
	Ipv4Common01Prefix string = "192." // 常见Ipv4
	Ipv4Common02Prefix string = "172." // 常见Ipv4
	LocalIpPrefix      string = "127." // 本地环回
)

// 获取本机服务器ip地址
func GetIpAddress() (string, error) {
	var retAddr string
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Error("GetIpAddress InterfaceAddrs fail", zap.Error(err))
		return retAddr, err
	}

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if strings.Contains(ipnet.IP.String(), Ipv4Common01Prefix) {
				log.Info("GetIpAddress LingXi prod Bj env ip match.", zap.String("ip", ipnet.IP.String()))
				retAddr = ipnet.IP.String()
				break
			} else if strings.Contains(ipnet.IP.String(), Ipv4Common02Prefix) {
				log.Info("GetIpAddress LingXi prod Gz env ip match.", zap.String("ip", ipnet.IP.String()))
				retAddr = ipnet.IP.String()
				break
			} else if ipnet.IP.To4() != nil && strings.Contains(ipnet.IP.String(), LocalIpPrefix) {
				log.Info("GetIpAddress unknown ipv4 ip.", zap.String("ip", ipnet.IP.String()))
				retAddr = ipnet.IP.String()
				break
			} else {
				continue
			}
		}
	}

	log.Warn("GetIpAddress NO Address match", zap.Any("addrs", addrs))
	return retAddr, nil
}

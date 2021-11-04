package flow

import (
	"fmt"
	"net"
)

func GetLocalIp() []string {
	addrs, err := net.InterfaceAddrs()
	var ipInfo []string
	if err != nil {
		fmt.Println(err)
		return ipInfo
	}
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ipInfo = append(ipInfo, ipnet.IP.String())
				//fmt.Println("本机IP信息: ", ipnet.IP.String())
			}
		}
	}
	return ipInfo
}

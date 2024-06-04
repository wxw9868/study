package main

import (
	"fmt"
	log "github.com/golang/glog"
	"math"
	"net"
	"sort"
	"strings"
)

type ConsulBind struct {
	Addr  string
	IpInt float64
}
type ConsulBindList []ConsulBind

func (s ConsulBindList) Len() int {
	return len(s)
}
func (s ConsulBindList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ConsulBindList) Less(i, j int) bool {
	return s[i].IpInt < s[j].IpInt
}
func (s ConsulBindList) ToStrings() []string {
	ret := make([]string, 0, len(s))
	for _, cbl := range s {
		ret = append(ret, cbl.Addr)
	}
	return ret
}
func BingConsulSort(consulAddrs []string) []string {
	localIpStr, err := GetAgentLocalIP()
	if err != nil {
		return consulAddrs
	}
	localIp := net.ParseIP(localIpStr)
	localIpInt := int64(0)
	if localIp != nil {
		localIpInt = util.InetAton(localIp)
	}
	addrslist := make([]ConsulBind, 0, len(consulAddrs))
	for _, addr := range consulAddrs {
		ads := strings.Split(addr, ":")
		if len(ads) == 2 {
			ip := net.ParseIP(ads[0])
			if ip != nil {
				ipInt := util.InetAton(ip)
				fmt.Println("ip:", ip, ipInt, localIpInt, (ipInt - localIpInt))
				addrslist = append(addrslist, ConsulBind{
					Addr:  addr,
					IpInt: math.Abs(float64(ipInt - localIpInt)),
				})
			}
		}
	}
	consulBindList := ConsulBindList(addrslist)
	sort.Sort(consulBindList)
	log.Infof("sort addrs %v", consulBindList)
	return consulBindList.ToStrings()
}

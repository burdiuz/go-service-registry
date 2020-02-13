package utils

import (
	"fmt"
	"strings"
)

func SplitRemoteAddr(addr string) ([]string, error) {
	index := strings.LastIndex(addr, ":")

	if index < 0 {
		return nil, fmt.Errorf("Cound not find port value in %q remote address", addr)
	}

	return []string{addr[:index], addr[index+1:]}, nil
}

func GetRemoteAddrIp(addr string) string {
	index := strings.LastIndex(addr, ":")

	if index < 0 {
		return addr
	}

	return addr[:index]
}

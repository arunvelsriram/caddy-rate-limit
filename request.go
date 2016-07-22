package ratelimit

import (
	"net"
	"net/http"
	"strings"
)

// IsLocalIpAddress check whether a ip belongs to private network
func IsLocalIpAddress(address string, localIpNets []*net.IPNet) bool {

	ip := net.ParseIP(address)
	if ip != nil {
		for _, ipNet := range localIpNets {
			if ipNet.Contains(ip) {
				return true
			} else {
				return false
			}
		}
	}

	return false
}

// GetRemoteIP returns the ip of requester
// Doesn't care about the ip is real or not
func GetRemoteIP(r *http.Request) string {

	remoteAddress := r.RemoteAddr

	idx := strings.LastIndex(remoteAddress, ":")
	if idx == -1 {
		return remoteAddress
	}
	return remoteAddress[:idx]
}

package util

import (
	"net"
	"unicode"
)

func DomainToIp(domain string) string {
	ips, _ := net.LookupIP(domain)
	for _, ip := range ips {
		if ipv4 := ip.To4(); ipv4 != nil {
			return ipv4.String()
		}
	}
	return ""
}

func ContainsOnlyNumbers(s string) bool {
	for _, r := range s {
		if string(r) == "." {
			continue
		}
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

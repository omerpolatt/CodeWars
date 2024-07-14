package main

import (
	"strconv"
	"strings"
)

/*
Ip adreslerini tam sayı olarak ifade edip  Ip adreslerini karşılaştırmamızı isteyen program
( A program that expresses Ip addresses as integers and asks us to compare Ip addresses )
*/

func IpAccount(ip string) int {

	ipadress := strings.Split(ip, ".")
	result := 0
	for _, ips := range ipadress {

		ip, _ := strconv.Atoi(ips)
		result = result*256 + ip
	}
	return result
}

func IpsBetween(start, end string) int {

	startip := IpAccount(start)
	endip := IpAccount(end)

	return endip - startip

}

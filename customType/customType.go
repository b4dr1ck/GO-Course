package main

import "regexp"

type IP string

func (value IP) validate() bool {
	if value == "" || !regexp.MustCompile(`^(?:[0-9]{1,3}\.){3}[0-9]{1,3}$`).MatchString(string(value)) {
		return false
	}
	return true
}

func main() {
	var ipAddress IP = "192.160.1.1"
	if ipAddress.validate() {
		println("Valid IP address:", ipAddress)
	} else {
		println("Invalid IP address:", ipAddress)
	}
}

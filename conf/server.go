package conf

import (
	"strconv"
)

const (
	Host   = "0.0.0.0"
	Port   = 80
	Scheme = "http://"

	RedisAddr     = "127.0.0.1:6379"
	RedisPassword = ""
	RedisDb       = 0
)

func GetHost() string {
	host := Host + ":" + strconv.Itoa(Port)
	return host
}

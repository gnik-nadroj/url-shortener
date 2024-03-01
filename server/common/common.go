package common

import (
	"os"
)

const (
	RedisUrl = "REDIS_URL"
	RedisPassword = "REDIS_PASSWORD"
	ServerPort = "SERVER_PORT"
	ServerAddr = "SERVER_ADDR"
	SessionSecret = "SESSION_SECRET"
)

func GetEnv(v string) string {
	value := os.Getenv(v)

	if value == "" {
		switch v {
		case RedisUrl:
			value = "localhost:6379"
		case ServerPort:
			value = "3000"
		case ServerAddr:
			value = "http://localhost:3000/"
		case SessionSecret:
			value = "fN#!3bdrL#eKhXz65hXMNR$&&oHb5bYG8zkYRG6J"
		}
	}

	return value
}

func IdToUrl(s string) string {
	host := GetEnv(ServerAddr)

	return host + s
}

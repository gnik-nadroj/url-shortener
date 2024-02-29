package common

import (
	"os"
)

func GetEnv(v string)  string {
	value :=  os.Getenv(v)

	if value == "" {
		switch v {
		case "REDIS_URL":
			value = "localhost:6379"
		case "SERVER_PORT":
			value = "3000"
		}
	}

	return value
}
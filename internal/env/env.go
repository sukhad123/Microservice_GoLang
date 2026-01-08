package env

import (
	"fmt"
	"os"
)

func GetString(key, fallback string) string {
	//if the env doesnt exists retur
	if val := os.Getenv(key); val != "" {
		fmt.Print("key is", key)
		return val
	}
	//if the value doesn't exists return fallback
	return fallback
}

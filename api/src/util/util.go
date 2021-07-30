package util

import (
	"fmt"
	"os"
	"strconv"
)

const (
	portEnv = "PARROT_API_PORT"
)

func GetPort() int {
	port := 8080
	if pString := os.Getenv(portEnv); pString != "" {
		p, err := strconv.Atoi(pString)
		if err != nil {
			fmt.Printf("%s requires an integer value, received: %s\n", portEnv, pString)
			fmt.Printf("Using default port: %d\n", port)
		} else {
			port = p
		}
	}
	return port
}

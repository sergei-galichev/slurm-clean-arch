package main

import (
	"fmt"
	"os"
)

func main() {
	host, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hostname: %s\n", host)
	fmt.Printf("Home dir: %s\n", homeDir)
	fmt.Printf("%s\n", os.Getenv("LANGUAGE"))
}

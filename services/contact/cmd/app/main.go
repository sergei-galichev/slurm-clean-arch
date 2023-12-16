package main

import (
	"fmt"
	"os"
	"slurm-clean-arch/pkg/store/postgres"
)

func main() {
	conn, err := postgres.New(postgres.Settings{})
	if err != nil {
		panic(err)
	}

	defer conn.Pool.Close()
	fmt.Println(conn.Pool.Stat())

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

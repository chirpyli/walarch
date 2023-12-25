package main

import (
	"github.com/chirpyli/walarch/cmd"
	"log"
)

func main() {
	log.Println("PostgreSQL WAL archive to S3 tools.")

	cmd.Execute()
}

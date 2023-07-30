package main

import (
	"minio/cmd"
	"minio/initsystem"
)

func main() {
	initsystem.Help()
	cmd.Option()
}

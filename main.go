package main

import (
	"flag"
	"fmt"
	"os"

	"./config"
	"./db"
	"./server"
)

func main() {
	environment := flag.String("e", "DEV", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
	db.Init()
	server.Init()
}

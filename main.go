package main

import (
	"flag"

	"github.com/wediin/curator/server"
)

var (
	help      bool
	debug     bool
	port      int
	storePath string
)

func init() {
	flag.BoolVar(&help, "h", false, "show this help")
	flag.BoolVar(&debug, "d", false, "enable debug mode")
	flag.IntVar(&port, "p", 9527, "port number")
	flag.StringVar(&storePath, "s", "/usr/lib/curator/", "local storage path")
}

func main() {
	flag.Parse()

	if help {
		flag.Usage()
		return
	}

	config := server.Config{
		Debug:     debug,
		Port:      port,
		StorePath: storePath,
	}
	server.Init(config)
}

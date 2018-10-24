package main

import (
	"flag"
	"log"
	"os"

	"github.com/wediin/curator/server"
)

var (
	help                 bool
	debug                bool
	port                 int
	url                  string
	storePath            string
	mongoServer          string
	mongoDB              string
	photoMongoCollection string
	photoRouter          string
	photoDir             string
)

func init() {
	flag.BoolVar(&help, "help", false, "show this help")
	flag.BoolVar(&debug, "debug", false, "enable debug mode")
	flag.IntVar(&port, "port", 9527, "port number")
	flag.StringVar(&url, "url", "http://localhost:9527", "url for client to access")
	flag.StringVar(&storePath, "store-path", "/var/lib/curator/", "local storage path")
	flag.StringVar(&mongoServer, "mongo-server", "mongodb://localhost:27017", "url to mongodb server")
	flag.StringVar(&mongoDB, "mongo-db", "gallery", "mongo database name")
	flag.StringVar(&photoMongoCollection, "photo-mongo-collection", "photos", "mongo collection name for photos")
	flag.StringVar(&photoRouter, "photo-router", "/photos", "router path of photos")
	flag.StringVar(&photoDir, "photo-dir", "photos", "directory name to store photos under store path")
}

func main() {
	flag.Parse()

	if help {
		flag.Usage()
		return
	}

	config := server.Config{
		Debug:                debug,
		Port:                 port,
		Url:                  url,
		StorePath:            storePath,
		MongoServer:          mongoServer,
		MongoDB:              mongoDB,
		PhotoMongoCollection: photoMongoCollection,
		PhotoRouter:          photoRouter,
		PhotoDir:             photoDir,
	}

	if err := server.Init(config); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	os.Exit(0)
}

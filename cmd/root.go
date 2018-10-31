package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wediin/curator/server"
)

var (
	debug                bool
	port                 int
	url                  string
	webPath              string
	storePath            string
	mongoServer          string
	mongoDB              string
	photoMongoCollection string
	photoRouter          string
	photoDir             string
	thumbWidth           int
)

var RootCmd = &cobra.Command{
	Use:   "curator",
	Short: "gallery backend server",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := server.NewServer(&server.Config{
			Debug:                debug,
			Url:                  url,
			WebPath:              webPath,
			StorePath:            storePath,
			MongoServer:          mongoServer,
			MongoDB:              mongoDB,
			PhotoMongoCollection: photoMongoCollection,
			PhotoRouter:          photoRouter,
			PhotoDir:             photoDir,
			ThumbWidth:           thumbWidth,
		})
		if err != nil {
			return err
		}
		s.Run(port)
		return nil
	},
}

func init() {
	RootCmd.Flags().BoolVarP(&debug, "debug", "d", false, "enable debug mode")
	RootCmd.Flags().IntVarP(&port, "port", "p", 9527, "port number")
	RootCmd.Flags().StringVarP(&url, "url", "", "http://localhost:9527", "url for client to access")
	RootCmd.Flags().StringVarP(&webPath, "web-path", "", "./public", "static web path")
	RootCmd.Flags().StringVarP(&storePath, "store-path", "", "/var/lib/curator/", "local storage path")
	RootCmd.Flags().StringVarP(&mongoServer, "mongo-server", "", "mongodb://localhost:27017", "url to mongodb server")
	RootCmd.Flags().StringVarP(&mongoDB, "mongo-db", "", "gallery", "mongo database name")
	RootCmd.Flags().StringVarP(&photoMongoCollection, "photo-mongo-collection", "", "photos", "mongo collection name for photos")
	RootCmd.Flags().StringVarP(&photoRouter, "photo-router", "", "/photos", "router path of photos")
	RootCmd.Flags().StringVarP(&photoDir, "photo-dir", "", "photos", "directory name to store photos under store path")
	RootCmd.Flags().IntVarP(&thumbWidth, "thumb-width", "", 300, "width in pixel to resize photos to thumbs")
}

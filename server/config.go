package server

type Config struct {
	Debug                bool
	Url                  string
	WebPath              string
	StorePath            string
	MongoServer          string
	MongoDB              string
	PhotoMongoCollection string
	PhotoRouter          string
	PhotoDir             string
	ThumbMaxLen          int
	WebviewMaxLen        int
	SizeMax              int
}

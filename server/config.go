package server

type Config struct {
	Debug                bool
	Url                  string
	StorePath            string
	MongoServer          string
	MongoDB              string
	PhotoMongoCollection string
	PhotoRouter          string
	PhotoDir             string
}

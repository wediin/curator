package db

import (
	"context"

	"github.com/mongodb/mongo-go-driver/mongo"
)

type Client struct {
	Client *mongo.Client
}

func NewClient(url string) (*Client, error) {
	client, err := mongo.NewClient(url)
	if err != nil {
		return nil, err
	}

	return &Client{
		Client: client,
	}, nil
}

func (client *Client) Insert(database string, collection string, document interface{}) error {
	err := client.Client.Connect(context.TODO())
	if err != nil {
		return err
	}
	defer client.Client.Disconnect(context.Background())

	c := client.Client.Database(database).Collection(collection)

	_, err = c.InsertOne(context.Background(), document)
	if err != nil {
		return err
	}

	return nil
}

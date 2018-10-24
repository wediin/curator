package db

import (
	"context"

	"github.com/mongodb/mongo-go-driver/mongo"
)

type Client struct {
	Client *mongo.Client
}

func NewClient(uri string) (*Client, error) {
	client, err := mongo.Connect(context.TODO(), uri)
	if err != nil {
		return nil, err
	}

	return &Client{
		Client: client,
	}, nil
}

func (c *Client) Connect() error {
	return c.Client.Connect(context.TODO())
}

func (c *Client) Ping() error {
	return c.Client.Ping(context.Background(), nil)
}

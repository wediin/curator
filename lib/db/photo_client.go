package db

import (
	"context"

	"github.com/mongodb/mongo-go-driver/mongo"
)

type PhotoClient struct {
	Database   string
	Collection string
	Client     *Client
}

func NewPhotoClient(uri string, database string, collection string) (*PhotoClient, error) {
	client, err := NewClient(uri)
	if err != nil {
		return nil, err
	}

	return &PhotoClient{
		Database:   database,
		Collection: collection,
		Client:     client,
	}, nil
}

func (c *PhotoClient) Insert(model *PhotoModel) error {
	collection, err := c.getCollection()
	if err != nil {
		return err
	}

	_, err = collection.InsertOne(context.Background(), model)
	if err != nil {
		return err
	}

	return nil
}

func (c *PhotoClient) Select() ([]*PhotoModel, error) {
	collection, err := c.getCollection()
	if err != nil {
		return nil, err
	}

	res, err := collection.Find(context.Background(), nil)
	defer res.Close(context.Background())

	photos := make([]*PhotoModel, 0)
	for res.Next(context.Background()) {
		photo := PhotoModel{}
		err := res.Decode(&photo)
		if err != nil {
			return nil, err
		}

		photos = append(photos, &photo)
	}

	return photos, nil
}

func (c *PhotoClient) getCollection() (*mongo.Collection, error) {
	if err := c.Client.Ping(); err != nil {
		if err := c.Client.Connect(); err != nil {
			return nil, err
		}
		if err = c.Client.Ping(); err != nil {
			return nil, err
		}
	}
	return c.Client.Client.Database(c.Database).Collection(c.Collection), nil
}

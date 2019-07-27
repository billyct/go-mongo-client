package go_mongo_client

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Client struct {
	Collection *mongo.Collection
	Client *mongo.Client
	Ctx context.Context
}

func NewClient(uri string, database string, collection string) (*Client, error) {
	c := new(Client)

	c.Ctx = context.TODO()
	client, err := mongo.Connect(c.Ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	c.Client = client
	c.Collection = client.Database(database).Collection(collection)

	return c, nil
}

func (c *Client) Update(filter interface{}, update interface{}) error {
	_, err := c.Collection.UpdateOne(c.Ctx, filter, update)
	return err
}

func (c *Client) Walk(cb func(*mongo.Cursor) error) error {

	findOptions := options.Find()
	findOptions.SetNoCursorTimeout(true)

	filter := bson.D{{}}

	cur, err := c.Collection.Find(c.Ctx, filter, findOptions)
	if err != nil {
		return err
	}

	defer cur.Close(c.Ctx)

	for cur.Next(c.Ctx) {
		err = cb(cur)
		if err != nil {
			return err
		}
	}

	err = cur.Err()
	if err != nil {
		return err
	}
}

package go_mongo_client

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Client struct {
	Uri        string
	Database   string
	Collection string
}

func (c *Client) Walk(cb func(*mongo.Cursor) error) {
	ctx := context.TODO()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(c.Uri))
	checkError(err)

	collection := client.Database(c.Database).Collection(c.Collection)

	findOptions := options.Find()
	findOptions.SetNoCursorTimeout(true)
	findOptions.SetLimit(20)

	filter := bson.D{{}}

	cur, err := collection.Find(ctx, filter, findOptions)
	checkError(err)
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		err = cb(cur)
		checkError(err)
	}

	err = cur.Err()
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

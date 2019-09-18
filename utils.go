package go_mongo_client

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CursorToMap(c *mongo.Cursor) (m map[string]interface{}, err error) {
	var d bson.D

	err = c.Decode(&d)
	if err != nil {
		return
	}

	t, err := bson.MarshalExtJSON(d, true, true)
	if err != nil {
		return
	}

	err = bson.UnmarshalExtJSON(t, true, &m)
	if err != nil {
		return
	}

	return
}

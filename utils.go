package go_mongo_client

import (
	"encoding/json"
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

	err = json.Unmarshal(t, &m)
	if err != nil {
		return
	}

	return
}

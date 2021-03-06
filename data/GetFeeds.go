package data

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func GetFeeds() ([]ChannelFeed, error) {
	client := GetClient()
	defer DisconnectClient(client)

	collection := client.Database("hololive-en").Collection("feeds")
	cur, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}

	var results []ChannelFeed
	if err = cur.All(context.Background(), &results); err != nil {
		return nil, err
	}

	return results, nil
}

package bd

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func BorroTweet(tweetID string, userID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database(DATABASE_NAME)
	col := db.Collection(COLLECTION_TWEET)

	objID, _ := primitive.ObjectIDFromHex(tweetID)

	condicion := bson.M{
		"_id":    objID,
		"userid": userID,
	}
	_, err := col.DeleteOne(ctx, condicion)
	if err != nil {
		log.Fatal(err.Error())
		return false, err
	}
	return true, nil
}

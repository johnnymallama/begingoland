package bd

import (
	"context"
	"log"
	"time"

	"github.com/johnnymallama/begingoland/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func LeoTweet(ID string, pagina int64) ([]*models.DevuelvoTweet, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database(DATABASE_NAME)
	col := db.Collection(COLLECTION_TWEET)

	var resultado []*models.DevuelvoTweet

	condicion := bson.M{
		"userid": ID,
	}

	opcion := options.Find()
	opcion.SetLimit(20)
	opcion.SetSkip((pagina - 1) * 20)
	opcion.SetSort(bson.D{
		{Key: "fecha", Value: -1},
	})

	cursor, err := col.Find(ctx, condicion, opcion)
	if err != nil {
		log.Fatal(err.Error())
		return resultado, false, err
	}
	for cursor.Next(context.TODO()) {
		var registro models.DevuelvoTweet
		err := cursor.Decode(&registro)
		if err != nil {
			return resultado, false, err
		}
		resultado = append(resultado, &registro)
	}
	return resultado, true, nil

}

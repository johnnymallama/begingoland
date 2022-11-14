package bd

import (
	"context"
	"time"

	"github.com/johnnymallama/begingoland/models"
	"go.mongodb.org/mongo-driver/bson"
)

func LeoTweetSeguidor(ID string, pagina int64) ([]models.DevuelvoTweetSeguidor, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database(DATABASE_NAME)
	col := db.Collection(COLLECTION_RELACION)

	skip := (pagina - 1) * 20

	condicion := make([]bson.M, 0)

	condicion = append(condicion, bson.M{"$match": bson.M{"usuarioid": ID}})
	condicion = append(condicion, bson.M{
		"$lookup": bson.M{
			"from":         "tweet",
			"localField":   "usuariorelacionid",
			"foreignField": "userid",
			"as":           "tweet",
		}})
	condicion = append(condicion, bson.M{"$unwind": "$tweet"})
	condicion = append(condicion, bson.M{"$sort": bson.M{"fecha": -1}})
	condicion = append(condicion, bson.M{"$skip": skip})
	condicion = append(condicion, bson.M{"$limit": 20})

	cursor, _ := col.Aggregate(ctx, condicion)
	var resultado []models.DevuelvoTweetSeguidor
	err := cursor.All(ctx, &resultado)
	if err != nil {
		return resultado, false, err
	}
	return resultado, true, nil
}

package bd

import (
	"context"
	"time"

	"github.com/johnnymallama/begingoland/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertoRegistro(u models.Usuario) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database(DATABASE_NAME)
	col := db.Collection(COLLECTION_USUARIOS)

	passwordEncrypt, _ := EncriptarPassword(u.Password)

	u.Password = passwordEncrypt

	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}
	ObjID := result.InsertedID
	return ObjID.(primitive.ObjectID).Hex(), true, nil
}

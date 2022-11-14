package bd

import (
	"context"
	"log"
	"time"

	"github.com/johnnymallama/begingoland/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func LeoUsuarioTodo(usuarioId string, pagina int64, search string, tipo string) ([]*models.Usuario, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database(DATABASE_NAME)
	col := db.Collection(COLLECTION_USUARIOS)

	var resultado []*models.Usuario

	findOption := options.Find()
	findOption.SetSkip((pagina - 1) * 20)
	findOption.SetLimit(20)

	query := bson.M{
		"nombre": bson.M{
			"$regex": `(?i)` + search,
		},
	}
	cursor, err := col.Find(ctx, query, findOption)

	if err != nil {
		log.Fatal(err.Error())
		return resultado, false, err
	}
	for cursor.Next(context.TODO()) {
		var usuario models.Usuario
		err := cursor.Decode(&usuario)
		if err != nil {
			log.Fatal(err.Error())
			return nil, false, err
		}
		var relacion models.Relacion
		relacion.UsuarioID = usuarioId
		relacion.UsuarioRelacionID = usuario.ID.Hex()

		incluir := false

		status, _ := ConsultoRelacion(relacion)

		if tipo == "new" && !status {
			incluir = true
		}
		if tipo == "follow" && status {
			incluir = true
		}

		if relacion.UsuarioRelacionID == usuarioId {
			incluir = false
		}

		if incluir {
			usuario.Password = ""
			usuario.Biografia = ""
			usuario.SitioWeb = ""
			usuario.Ubicacion = ""
			usuario.Avatar = ""
			usuario.Banner = ""
			resultado = append(resultado, &usuario)
		}
	}
	err = cursor.Err()
	if err != nil {
		log.Fatal(err.Error())
		return nil, false, err
	}
	_ = cursor.Close(context.TODO())
	return resultado, true, nil
}

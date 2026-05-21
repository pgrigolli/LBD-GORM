package services

import (
	"LBD/database"
	"LBD/schemas"
	"context"

	"gorm.io/gorm"
)

func CreateArtista(artista schemas.Artista) error {

	db, err := database.ConnectDB()

	if err != nil {
		return err
	}

	ctx := context.Background()
	err = gorm.G[schemas.Artista](db).Create(ctx, &artista)

	return err
}

func GetAllArtista() ([]schemas.Artista, error) {

	db, err := database.ConnectDB()

	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	artistas, err := gorm.G[schemas.Artista](db).Find(ctx)

	return artistas, err

}

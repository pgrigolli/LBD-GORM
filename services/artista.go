package services

import (
	"LBD/database"
	"LBD/schemas"
	"context"
	"fmt"

	"gorm.io/gorm"
)

func CreateArtista(artista schemas.Artista) error {

	db, err := database.ConnectDB()

	if err != nil {
		return err
	}

	ctx := context.Background()
	err = gorm.G[schemas.Artista](db.Debug()).Create(ctx, &artista)

	return err
}

func GetArtista(id uint) (schemas.Artista, error) {

	db, err := database.ConnectDB()

	if err != nil {
		return schemas.Artista{}, err
	}

	ctx := context.Background()
	artista, err := gorm.G[schemas.Artista](db.Debug()).Where("id = ?", id).First(ctx)

	return artista, err
}

func GetAllArtista() ([]schemas.Artista, error) {

	db, err := database.ConnectDB()

	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	artistas, err := gorm.G[schemas.Artista](db.Debug()).Find(ctx)

	return artistas, err

}

func UpdateArtista(artista schemas.Artista) (schemas.Artista, error) {
	db, err := database.ConnectDB()

	if err != nil {
		return schemas.Artista{}, err
	}

	ctx := context.Background()
	rows, err := gorm.G[schemas.Artista](db.Debug()).Where("id = ?", artista.ID).Updates(ctx, artista)

	if err != nil {
		return schemas.Artista{}, err
	}

	if rows == 0 {
		return schemas.Artista{}, fmt.Errorf("artista com id %d nao encontrado", artista.ID)
	}

	return artista, nil

}

func DeleteArtista(artistaID uint) error {

	db, err := database.ConnectDB()

	if err != nil {
		return err
	}

	ctx := context.Background()

	rows, err := gorm.G[schemas.Artista](db.Debug()).Where("id = ?", artistaID).Delete(ctx)
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("artista com id %d nao encontrado", artistaID)
	}

	return nil
}

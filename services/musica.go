package services

import (
	"LBD/database"
	"LBD/schemas"
	"context"
	"fmt"

	"gorm.io/gorm"
)

func CreateMusica(musica schemas.Musica) error {

	db, err := database.ConnectDB()

	if err != nil {
		return err
	}

	ctx := context.Background()
	return gorm.G[schemas.Musica](db).Create(ctx, &musica)
}

func GetMusica(id uint) (schemas.Musica, error) {

	db, err := database.ConnectDB()

	if err != nil {
		return schemas.Musica{}, err
	}

	ctx := context.Background()
	return gorm.G[schemas.Musica](db).Where("id = ?", id).First(ctx)
}

func GetAllMusica() ([]schemas.Musica, error) {

	db, err := database.ConnectDB()

	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	return gorm.G[schemas.Musica](db).Find(ctx)
}

func UpdateMusica(musica schemas.Musica) (schemas.Musica, error) {

	db, err := database.ConnectDB()

	if err != nil {
		return schemas.Musica{}, err
	}

	ctx := context.Background()
	rows, err := gorm.G[schemas.Musica](db).Where("id = ?", musica.ID).Updates(ctx, musica)

	if err != nil {
		return schemas.Musica{}, err
	}

	if rows == 0 {
		return schemas.Musica{}, fmt.Errorf("musica com id %d nao encontrada", musica.ID)
	}

	return musica, nil
}

func DeleteMusica(id uint) error {

	db, err := database.ConnectDB()

	if err != nil {
		return err
	}

	ctx := context.Background()
	rows, err := gorm.G[schemas.Musica](db).Where("id = ?", id).Delete(ctx)

	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("musica com id %d nao encontrada", id)
	}

	return nil
}

package services

import (
	"LBD/schemas"
	"context"
	"fmt"

	"gorm.io/gorm"
)

func CreateMusica(musica schemas.Musica) error {

	db, err := connectDB()

	if err != nil {
		return err
	}

	ctx := context.Background()
	return gorm.G[schemas.Musica](db.Debug()).Create(ctx, &musica)
}

func GetMusica(id uint) (schemas.Musica, error) {

	db, err := connectDB()

	if err != nil {
		return schemas.Musica{}, err
	}

	ctx := context.Background()
	return gorm.G[schemas.Musica](db.Debug()).Where("id = ?", id).First(ctx)
}

func GetAllMusica() ([]schemas.Musica, error) {

	db, err := connectDB()

	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	return gorm.G[schemas.Musica](db.Debug()).Find(ctx)
}

func GetMusicasLedZeppelinMaioresQueMaiorMusicaQueen() ([]schemas.Musica, error) {
	db, err := connectDB()
	if err != nil {
		return nil, err
	}

	maiorDuracaoQueen := db.
		Table(`"MUSICA" mq`).
		Select("MAX(mq.duracao_segundos)").
		Joins(`INNER JOIN "ARTISTA" aq ON aq.id = mq.artista_id`).
		Where("aq.nome = ?", "Queen")

	var musicas []schemas.Musica
	err = db.Debug().
		Table(`"MUSICA" m`).
		Select("m.*").
		Joins(`INNER JOIN "ARTISTA" a ON a.id = m.artista_id`).
		Where("a.nome = ?", "Led Zeppelin").
		Where("m.duracao_segundos > (?)", maiorDuracaoQueen).
		Order("m.duracao_segundos DESC").
		Order("m.titulo ASC").
		Scan(&musicas).Error

	return musicas, err
}

func UpdateMusica(musica schemas.Musica) (schemas.Musica, error) {

	db, err := connectDB()

	if err != nil {
		return schemas.Musica{}, err
	}

	ctx := context.Background()
	rows, err := gorm.G[schemas.Musica](db.Debug()).Where("id = ?", musica.ID).Updates(ctx, musica)

	if err != nil {
		return schemas.Musica{}, err
	}

	if rows == 0 {
		return schemas.Musica{}, fmt.Errorf("musica com id %d nao encontrada", musica.ID)
	}

	return musica, nil
}

func DeleteMusica(id uint) error {

	db, err := connectDB()

	if err != nil {
		return err
	}

	ctx := context.Background()
	rows, err := gorm.G[schemas.Musica](db.Debug()).Where("id = ?", id).Delete(ctx)

	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("musica com id %d nao encontrada", id)
	}

	return nil
}

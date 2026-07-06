package services

import (
	"LBD/schemas"
	"context"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func GetMusicasByUsuarioAndArtista(username, nomeArtista string) ([]schemas.Musica, error) {
	db, err := connectDB()
	if err != nil {
		return nil, err
	}

	ctx := context.Background()

	usuario, err := gorm.G[schemas.Usuario](db.Debug()).Where("username = ?", username).First(ctx)
	if err != nil {
		return nil, fmt.Errorf("usuario '%s' nao encontrado: %w", username, err)
	}

	artista, err := gorm.G[schemas.Artista](db.Debug()).Where("nome = ?", nomeArtista).First(ctx)
	if err != nil {
		return nil, fmt.Errorf("artista '%s' nao encontrado: %w", nomeArtista, err)
	}

	var musicas []schemas.Musica
	err = db.Debug().
		Table(`"MUSICA" m`).
		Select("DISTINCT m.*").
		Joins(`INNER JOIN "MUSICA_PLAYLIST" mp ON mp.musica_id = m.id`).
		Joins(`INNER JOIN "PLAYLIST" p ON p.playlist_id = mp.playlist_id AND p.usuario_id = mp.usuario_id`).
		Where("m.artista_id = ?", artista.ID).
		Where("p.usuario_id = ?", usuario.ID).
		Scan(&musicas).Error

	return musicas, err
}

func GetMusicaComArtista(id uint) (schemas.Musica, error) {
	db, err := connectDB()
	if err != nil {
		return schemas.Musica{}, err
	}

	ctx := context.Background()
	musica, err := gorm.G[schemas.Musica](db.Debug()).
		Joins(clause.Has("Artista"), nil).
		Where(`"MUSICA"."id" = ?`, id).
		First(ctx)
	if err != nil {
		return schemas.Musica{}, fmt.Errorf("musica com id %d nao encontrada: %w", id, err)
	}

	return musica, nil
}

func GetMusicasMaisCurtasQueMediaDoArtista() ([]schemas.Musica, error) {
	db, err := connectDB()
	if err != nil {
		return nil, err
	}

	musicasComMediaDoArtista := db.
		Table(`"MUSICA" m`).
		Select("m.*, AVG(m.duracao_segundos) OVER (PARTITION BY m.artista_id) AS media_duracao_artista")
		// PARTITION BY cria uma "partição" para cada artista,
		// calculando a média da duração das músicas dentro de cada partição.
	var musicas []schemas.Musica
	err = db.Debug().
		Table("(?) AS sub", musicasComMediaDoArtista).
		Select("sub.*").
		Where("sub.duracao_segundos < sub.media_duracao_artista").
		Order("sub.artista_id ASC").
		Order("sub.duracao_segundos ASC").
		Scan(&musicas).Error

	return musicas, err
}

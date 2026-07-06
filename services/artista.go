package services

import (
	"LBD/schemas"
	"context"
	"fmt"

	"gorm.io/gorm"
)

type RankingPopularidadeArtista struct {
	Posicao             int    `gorm:"column:posicao"`
	ArtistaID           uint   `gorm:"column:artista_id"`
	Nome                string `gorm:"column:nome"`
	QuantidadePlaylists int64  `gorm:"column:quantidade_playlists"`
}

func CreateArtista(artista schemas.Artista) error {

	db, err := connectDB()

	if err != nil {
		return err
	}

	ctx := context.Background()
	err = gorm.G[schemas.Artista](db.Debug()).Create(ctx, &artista)

	return err
}

func GetArtista(id uint) (schemas.Artista, error) {

	db, err := connectDB()

	if err != nil {
		return schemas.Artista{}, err
	}

	ctx := context.Background()
	artista, err := gorm.G[schemas.Artista](db.Debug()).Where("id = ?", id).First(ctx)

	return artista, err
}

func GetAllArtista() ([]schemas.Artista, error) {

	db, err := connectDB()

	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	artistas, err := gorm.G[schemas.Artista](db.Debug()).Find(ctx)

	return artistas, err

}

func GetRankingPopularidadeArtista() ([]RankingPopularidadeArtista, error) {
	db, err := connectDB()
	if err != nil {
		return nil, err
	}

	playlistsPorArtista := db.
		Table(`"MUSICA" m`).
		Select("DISTINCT m.artista_id, mp.playlist_id, mp.usuario_id").
		Joins(`INNER JOIN "MUSICA_PLAYLIST" mp ON mp.musica_id = m.id`)

	var ranking []RankingPopularidadeArtista
	err = db.Debug().
		Table(`"ARTISTA" a`).
		Select("a.id AS artista_id, a.nome, COUNT(ppa.playlist_id) AS quantidade_playlists").
		Joins("LEFT JOIN (?) ppa ON ppa.artista_id = a.id", playlistsPorArtista).
		Group("a.id, a.nome").
		Order("quantidade_playlists DESC").
		Order("a.nome ASC").
		Scan(&ranking).Error
	if err != nil {
		return nil, err
	}

	for i := range ranking {
		ranking[i].Posicao = i + 1
	}

	return ranking, nil
}

func UpdateArtista(artista schemas.Artista) (schemas.Artista, error) {
	db, err := connectDB()

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

func GetArtistasSemMusicasEmPlaylist() ([]schemas.Artista, error) {
	db, err := connectDB()
	if err != nil {
		return nil, err
	}

	artistasComMusicaEmPlaylist := db.
		Table(`"MUSICA" m`).
		Select("DISTINCT m.artista_id").
		Joins(`INNER JOIN "MUSICA_PLAYLIST" mp ON mp.musica_id = m.id`)

	var artistas []schemas.Artista
	err = db.Debug().
		Table(`"ARTISTA" a`).
		Select("a.*").
		Where("a.id NOT IN (?)", artistasComMusicaEmPlaylist). //Subquery
		Order("a.nome ASC").
		Scan(&artistas).Error

	return artistas, err
}

func DeleteArtista(artistaID uint) error {

	db, err := connectDB()

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

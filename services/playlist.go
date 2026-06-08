package services

import (
	"LBD/database"
	"LBD/schemas"
	"context"
	"fmt"

	"gorm.io/gorm"
)

func CreatePlaylist(playlist schemas.Playlist) (schemas.Playlist, error) {

	db, err := database.ConnectDB()

	if err != nil {
		return schemas.Playlist{}, err
	}

	ctx := context.Background()
	err = gorm.G[schemas.Playlist](db.Debug()).Create(ctx, &playlist)

	return playlist, err
}

func AddMusicaToPlaylist(musicaId, playlistId, usuarioId uint) error {

	db, err := database.ConnectDB()

	if err != nil {
		return err
	}

	ctx := context.Background()

	existentes, err := gorm.G[schemas.MusicaPlaylist](db.Debug()).
		Where("playlist_id = ? AND usuario_id = ?", playlistId, usuarioId).
		Find(ctx)

	if err != nil {
		return err
	}

	mp := schemas.MusicaPlaylist{
		MusicaId:        musicaId,
		PlaylistId:      playlistId,
		UsuarioId:       usuarioId,
		OrdemNaPlaylist: len(existentes) + 1,
	}

	return gorm.G[schemas.MusicaPlaylist](db.Debug()).Create(ctx, &mp)
}

func RemoveMusicaFromPlaylist(musicaId, playlistId, usuarioId uint) error {

	db, err := database.ConnectDB()

	if err != nil {
		return err
	}

	ctx := context.Background()

	rows, err := gorm.G[schemas.MusicaPlaylist](db.Debug()).
		Where("musica_id = ? AND playlist_id = ? AND usuario_id = ?", musicaId, playlistId, usuarioId).
		Delete(ctx)

	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("musica %d nao encontrada na playlist (%d, %d)", musicaId, playlistId, usuarioId)
	}

	return nil
}

package services

import (
	"LBD/schemas"
	"context"
	"errors"
	"fmt"
	"sort"

	"gorm.io/gorm"
)

type MusicaNaPlaylist struct {
	Titulo          string `gorm:"column:titulo"`
	OrdemNaPlaylist int    `gorm:"column:ordem_na_playlist"`
}

func CreatePlaylist(playlist schemas.Playlist) (schemas.Playlist, error) {
	db, err := connectDB()
	if err != nil {
		return schemas.Playlist{}, err
	}

	ctx := context.Background()
	err = gorm.G[schemas.Playlist](db.Debug()).Create(ctx, &playlist)

	return playlist, err
}

func GetAllPlaylist() ([]schemas.Playlist, error) {
	db, err := connectDB()
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	return gorm.G[schemas.Playlist](db.Debug()).Find(ctx)
}

func AddMusicaToPlaylist(musicaId, playlistId, usuarioId uint) error {
	db, err := connectDB()
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
	db, err := connectDB()
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

func TransferirMusicaEntrePlaylists(musicaId, playlistOrigemId, playlistDestinoId, usuarioId uint) error {
	db, err := connectDB()
	if err != nil {
		return err
	}

	if playlistOrigemId == playlistDestinoId {
		return fmt.Errorf("playlist de origem e destino devem ser diferentes")
	}

	ctx := context.Background()
	return db.Debug().Transaction(func(tx *gorm.DB) error {
		_, err := gorm.G[schemas.Playlist](tx.Debug()).
			Where("playlist_id = ? AND usuario_id = ?", playlistOrigemId, usuarioId).
			First(ctx)
		if err != nil {
			return fmt.Errorf("playlist de origem (%d, %d) nao encontrada: %w", playlistOrigemId, usuarioId, err)
		}

		_, err = gorm.G[schemas.Playlist](tx.Debug()).
			Where("playlist_id = ? AND usuario_id = ?", playlistDestinoId, usuarioId).
			First(ctx)
		if err != nil {
			return fmt.Errorf("playlist de destino (%d, %d) nao encontrada: %w", playlistDestinoId, usuarioId, err)
		}

		_, err = gorm.G[schemas.MusicaPlaylist](tx.Debug()).
			Where("musica_id = ? AND playlist_id = ? AND usuario_id = ?", musicaId, playlistOrigemId, usuarioId).
			First(ctx)
		if err != nil {
			return fmt.Errorf("musica %d nao encontrada na playlist de origem (%d, %d): %w", musicaId, playlistOrigemId, usuarioId, err)
		}

		_, err = gorm.G[schemas.MusicaPlaylist](tx.Debug()).
			Where("musica_id = ? AND playlist_id = ? AND usuario_id = ?", musicaId, playlistDestinoId, usuarioId).
			First(ctx)
		if err == nil {
			return fmt.Errorf("musica %d ja existe na playlist de destino (%d, %d)", musicaId, playlistDestinoId, usuarioId)
		}
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		itensDestino, err := gorm.G[schemas.MusicaPlaylist](tx.Debug()).
			Where("playlist_id = ? AND usuario_id = ?", playlistDestinoId, usuarioId).
			Find(ctx)
		if err != nil {
			return err
		}

		proximaOrdem := 1
		for _, item := range itensDestino {
			if item.OrdemNaPlaylist >= proximaOrdem {
				proximaOrdem = item.OrdemNaPlaylist + 1
			}
		}

		rows, err := gorm.G[schemas.MusicaPlaylist](tx.Debug()).
			Where("musica_id = ? AND playlist_id = ? AND usuario_id = ?", musicaId, playlistOrigemId, usuarioId).
			Delete(ctx)
		if err != nil {
			return err
		}
		if rows == 0 {
			return fmt.Errorf("musica %d nao encontrada na playlist de origem (%d, %d)", musicaId, playlistOrigemId, usuarioId)
		}

		return gorm.G[schemas.MusicaPlaylist](tx.Debug()).Create(ctx, &schemas.MusicaPlaylist{
			MusicaId:        musicaId,
			PlaylistId:      playlistDestinoId,
			UsuarioId:       usuarioId,
			OrdemNaPlaylist: proximaOrdem,
		})
	})
}

func GetMusicasDaPlaylist(nomePlaylist string) ([]MusicaNaPlaylist, error) {
	db, err := connectDB()
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	playlist, err := gorm.G[schemas.Playlist](db.Debug()).Where("nome = ?", nomePlaylist).First(ctx)
	if err != nil {
		return nil, err
	}

	itens, err := gorm.G[schemas.MusicaPlaylist](db.Debug()).
		Preload("Musica", nil).
		Where("playlist_id = ? AND usuario_id = ?", playlist.PlaylistId, playlist.UsuarioId).
		Order("ordem_na_playlist ASC").
		Find(ctx)
	if err != nil {
		return nil, err
	}

	resultado := make([]MusicaNaPlaylist, 0, len(itens))
	for _, item := range itens {
		resultado = append(resultado, MusicaNaPlaylist{
			Titulo:          item.Musica.Titulo,
			OrdemNaPlaylist: item.OrdemNaPlaylist,
		})
	}

	return resultado, nil
}

func GetUsuariosDonoDaMusica(tituloMusica string) ([]string, error) {
	db, err := connectDB()
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	musica, err := gorm.G[schemas.Musica](db.Debug()).Where("titulo = ?", tituloMusica).First(ctx)
	if err != nil {
		return nil, err
	}

	itens, err := gorm.G[schemas.MusicaPlaylist](db.Debug()).
		Preload("Musica", nil).
		Where("musica_id = ?", musica.ID).
		Find(ctx)
	if err != nil {
		return nil, err
	}

	usuariosUnicos := make(map[string]struct{})
	for _, item := range itens {
		playlist, err := gorm.G[schemas.Playlist](db.Debug()).
			Preload("Usuario", nil).
			Where("playlist_id = ? AND usuario_id = ?", item.PlaylistId, item.UsuarioId).
			First(ctx)
		if err != nil {
			return nil, err
		}

		usuariosUnicos[playlist.Usuario.Username] = struct{}{}
	}

	resultado := make([]string, 0, len(usuariosUnicos))
	for username := range usuariosUnicos {
		resultado = append(resultado, username)
	}

	sort.Strings(resultado)
	return resultado, nil
}

func GetPlaylistsByUser(username string) ([]schemas.Playlist, error) {
	db, err := connectDB()
	if err != nil {
		return nil, err
	}

	ctx := context.Background()

	usuario, err := gorm.G[schemas.Usuario](db.Debug()).Where("username = ?", username).First(ctx)
	if err != nil {
		return nil, fmt.Errorf("usuario '%s' nao encontrado: %w", username, err)
	}

	return gorm.G[schemas.Playlist](db.Debug()).Where("usuario_id = ?", usuario.ID).Find(ctx)
}

type PlaylistQuantidadeMusicas struct {
	Nome              string `gorm:"column:nome"`
	QuantidadeMusicas int64  `gorm:"column:quantidade_musicas"`
}

func GetQuantidadeMusicasPorPlaylist() ([]PlaylistQuantidadeMusicas, error) {
	db, err := connectDB()
	if err != nil {
		return nil, err
	}

	var resultado []PlaylistQuantidadeMusicas
	err = db.Debug().
		Table(`"PLAYLIST" p`).
		Select("p.nome, COUNT(mp.musica_id) AS quantidade_musicas").
		Joins(`LEFT JOIN "MUSICA_PLAYLIST" mp ON mp.playlist_id = p.playlist_id AND mp.usuario_id = p.usuario_id`).
		Group("p.playlist_id, p.usuario_id, p.nome").
		Order("quantidade_musicas DESC").
		Order("p.nome ASC").
		Scan(&resultado).Error

	return resultado, err
}

type PlaylistTempoTotal struct {
	NomePlaylist       string `gorm:"column:nome_playlist"`
	UsernameDono       string `gorm:"column:username_dono"`
	TempoTotalSegundos int64  `gorm:"column:tempo_total_segundos"`
}

func GetTempoTotalPorPlaylist() ([]PlaylistTempoTotal, error) {
	db, err := connectDB()
	if err != nil {
		return nil, err
	}

	var resultado []PlaylistTempoTotal
	err = db.Debug().
		Table(`"PLAYLIST" p`).
		Select("p.nome AS nome_playlist, u.username AS username_dono, COALESCE(SUM(m.duracao_segundos), 0) AS tempo_total_segundos").
		Joins(`INNER JOIN "USUARIO" u ON u.id = p.usuario_id`).
		Joins(`LEFT JOIN "MUSICA_PLAYLIST" mp ON mp.playlist_id = p.playlist_id AND mp.usuario_id = p.usuario_id`).
		Joins(`LEFT JOIN "MUSICA" m ON m.id = mp.musica_id`).
		Group("p.playlist_id, p.usuario_id, p.nome, u.username").
		Order("tempo_total_segundos DESC").
		Order("p.nome ASC").
		Scan(&resultado).Error

	return resultado, err
}

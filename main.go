package main

import (
	"LBD/database"
	"LBD/schemas"
	"LBD/services"
	"database/sql"
	"fmt"
	"time"

	"gorm.io/gorm"
)

func main() {
	db, err := database.ConnectDB()
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(
		&schemas.Artista{},
		&schemas.Usuario{},
		&schemas.Musica{},
		&schemas.Playlist{},
		&schemas.MusicaPlaylist{},
	)
	if err != nil {
		panic(err)
	}

	sufixo := time.Now().Format("20060102150405")
	usuario := schemas.Usuario{
		Username: "Usuario Teste " + sufixo,
		Email:    "usuario.teste." + sufixo + "@email.com",
	}

	if err := db.Create(&usuario).Error; err != nil {
		panic(err)
	}

	var artistaID uint
	var musicaID uint
	var playlist schemas.Playlist

	defer func() {
		if playlist.PlaylistId != 0 {
			db.Where("playlist_id = ? AND usuario_id = ?", playlist.PlaylistId, playlist.UsuarioId).Delete(&schemas.MusicaPlaylist{})
			db.Where("playlist_id = ? AND usuario_id = ?", playlist.PlaylistId, playlist.UsuarioId).Delete(&schemas.Playlist{})
		}

		if musicaID != 0 {
			_ = services.DeleteMusica(musicaID)
		}

		if artistaID != 0 {
			_ = services.DeleteArtista(artistaID)
		}

		db.Delete(&schemas.Usuario{}, usuario.ID)
	}()

	fmt.Println("=== TESTES DO SERVICE DE ARTISTA ===")

	nomeArtista := "Artista Teste " + sufixo
	err = services.CreateArtista(schemas.Artista{
		Nome: nomeArtista,
		Nacionalidade: sql.NullString{
			String: "Brasileira",
			Valid:  true,
		},
	})
	printResultado("CreateArtista", err)

	artistas, err := services.GetAllArtista()
	printResultado("GetAllArtista", err)
	if err == nil {
		for _, artista := range artistas {
			if artista.Nome == nomeArtista {
				artistaID = artista.ID
				break
			}
		}
	}

	artista, err := services.GetArtista(artistaID)
	printResultado("GetArtista", err)
	if err == nil {
		fmt.Printf("Artista encontrado: ID=%d Nome=%s\n", artista.ID, artista.Nome)
	}

	artistaAtualizado, err := services.UpdateArtista(schemas.Artista{
		Model: gorm.Model{ID: artistaID},
		Nome:  nomeArtista + " Atualizado",
		Nacionalidade: sql.NullString{
			String: "Portuguesa",
			Valid:  true,
		},
	})
	printResultado("UpdateArtista", err)
	if err == nil {
		fmt.Printf("Artista atualizado: ID=%d Nome=%s\n", artistaAtualizado.ID, artistaAtualizado.Nome)
	}

	fmt.Println("\n=== TESTES DO SERVICE DE MUSICA ===")

	tituloMusica := "Musica Teste " + sufixo
	err = services.CreateMusica(schemas.Musica{
		Titulo:           tituloMusica,
		Duracao_segundos: 210,
		Artista_id:       artistaID,
	})
	printResultado("CreateMusica", err)

	musicas, err := services.GetAllMusica()
	printResultado("GetAllMusica", err)
	if err == nil {
		for _, musica := range musicas {
			if musica.Titulo == tituloMusica {
				musicaID = musica.ID
				break
			}
		}
	}

	musica, err := services.GetMusica(musicaID)
	printResultado("GetMusica", err)
	if err == nil {
		fmt.Printf("Musica encontrada: ID=%d Titulo=%s\n", musica.ID, musica.Titulo)
	}

	musicaAtualizada, err := services.UpdateMusica(schemas.Musica{
		Model:            gorm.Model{ID: musicaID},
		Titulo:           tituloMusica + " Atualizada",
		Duracao_segundos: 240,
		Artista_id:       artistaID,
	})
	printResultado("UpdateMusica", err)
	if err == nil {
		fmt.Printf("Musica atualizada: ID=%d Titulo=%s\n", musicaAtualizada.ID, musicaAtualizada.Titulo)
	}

	fmt.Println("\n=== TESTES DO SERVICE DE PLAYLIST ===")

	playlist, err = services.CreatePlaylist(schemas.Playlist{
		UsuarioId: usuario.ID,
		Nome:      "Playlist Teste " + sufixo,
	})
	printResultado("CreatePlaylist", err)
	if err == nil {
		fmt.Printf("Playlist criada: PlaylistId=%d UsuarioId=%d Nome=%s\n", playlist.PlaylistId, playlist.UsuarioId, playlist.Nome)
	}

	err = services.AddMusicaToPlaylist(musicaID, playlist.PlaylistId, playlist.UsuarioId)
	printResultado("AddMusicaToPlaylist", err)

	err = services.RemoveMusicaFromPlaylist(musicaID, playlist.PlaylistId, playlist.UsuarioId)
	printResultado("RemoveMusicaFromPlaylist", err)

	fmt.Println("\n=== TESTES DE DELETE ===")

	err = services.DeleteMusica(musicaID)
	printResultado("DeleteMusica", err)
	if err == nil {
		musicaID = 0
	}

	err = services.DeleteArtista(artistaID)
	printResultado("DeleteArtista", err)
	if err == nil {
		artistaID = 0
	}
}

func printResultado(nome string, err error) {
	if err != nil {
		fmt.Printf("[ERRO] %s: %v\n", nome, err)
		return
	}

	fmt.Printf("[OK] %s\n", nome)
}

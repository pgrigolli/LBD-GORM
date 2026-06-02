package seed

import (
	"database/sql"

	"LBD/schemas"

	"gorm.io/gorm"
)

func Seed(db *gorm.DB) {

	artistas := []schemas.Artista{
		{
			Nome: "Queen",
			Nacionalidade: sql.NullString{
				String: "Britânica",
				Valid:  true,
			},
		},
		{
			Nome: "Led Zeppelin",
			Nacionalidade: sql.NullString{
				String: "Britânica",
				Valid:  true,
			},
		},
		{
			Nome: "AC/DC",
			Nacionalidade: sql.NullString{
				String: "Australiana",
				Valid:  true,
			},
		},
		{
			Nome: "Banda X (Pop)",
			Nacionalidade: sql.NullString{
				String: "Brasileira",
				Valid:  true,
			},
		},
	}

	db.Create(&artistas)

	usuarios := []schemas.Usuario{
		{
			Username: "Pablo",
			Email:    "pablo@aluno.com",
		},
		{
			Username: "Josue",
			Email:    "josue@aluno.com",
		},
		{
			Username: "Alexandre",
			Email:    "alexandre@aluno.com",
		},
	}

	db.Create(&usuarios)

	musicas := []schemas.Musica{
		{
			Titulo:           "Bohemian Rhapsody",
			Duracao_segundos: 354,
			Artista_id:       artistas[0].ID,
		},
		{
			Titulo:           "Stairway to Heaven",
			Duracao_segundos: 482,
			Artista_id:       artistas[1].ID,
		},
		{
			Titulo:           "Back In Black",
			Duracao_segundos: 255,
			Artista_id:       artistas[2].ID,
		},
		{
			Titulo:           "We Will Rock You",
			Duracao_segundos: 160,
			Artista_id:       artistas[0].ID,
		},
		{
			Titulo:           "Musica Pop Brasileira",
			Duracao_segundos: 180,
			Artista_id:       artistas[3].ID,
		},
		{
			Titulo:           "Thunderstruck",
			Duracao_segundos: 292,
			Artista_id:       artistas[2].ID,
		},
	}

	db.Create(&musicas)

	playlists := []schemas.Playlist{
		{
			UsuarioId: usuarios[0].ID,
			Nome:      "Rock do Pablo",
		},
		{
			UsuarioId: usuarios[1].ID,
			Nome:      "Baladas do Josue",
		},
		{
			UsuarioId: usuarios[0].ID,
			Nome:      "Heavy Riffs",
		},
	}

	db.Create(&playlists)

	// Playlist "Rock do Pablo": Bohemian Rhapsody, Back In Black, We Will Rock You
	db.Create(&[]schemas.MusicaPlaylist{
		{MusicaId: musicas[0].ID, PlaylistId: playlists[0].PlaylistId, UsuarioId: playlists[0].UsuarioId, OrdemNaPlaylist: 1},
		{MusicaId: musicas[2].ID, PlaylistId: playlists[0].PlaylistId, UsuarioId: playlists[0].UsuarioId, OrdemNaPlaylist: 2},
		{MusicaId: musicas[3].ID, PlaylistId: playlists[0].PlaylistId, UsuarioId: playlists[0].UsuarioId, OrdemNaPlaylist: 3},
	})

	// Playlist "Baladas do Josue": Stairway to Heaven
	db.Create(&[]schemas.MusicaPlaylist{
		{MusicaId: musicas[1].ID, PlaylistId: playlists[1].PlaylistId, UsuarioId: playlists[1].UsuarioId, OrdemNaPlaylist: 1},
	})

	// Playlist "Heavy Riffs": Back In Black, Thunderstruck
	db.Create(&[]schemas.MusicaPlaylist{
		{MusicaId: musicas[2].ID, PlaylistId: playlists[2].PlaylistId, UsuarioId: playlists[2].UsuarioId, OrdemNaPlaylist: 1},
		{MusicaId: musicas[5].ID, PlaylistId: playlists[2].PlaylistId, UsuarioId: playlists[2].UsuarioId, OrdemNaPlaylist: 2},
	})
}

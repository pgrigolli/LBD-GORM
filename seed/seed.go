package seed

import (
	"database/sql"
	"time"

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

	now := time.Now()

	playlists := []schemas.Playlist{
		{
			Usuario_id: usuarios[0].ID,
			Nome:       "Rock do Pablo",
			Data_criacao: sql.NullTime{
				Time:  now,
				Valid: true,
			},
		},
		{
			Usuario_id: usuarios[1].ID,
			Nome:       "Baladas do Josue",
			Data_criacao: sql.NullTime{
				Time:  now,
				Valid: true,
			},
		},
		{
			Usuario_id: usuarios[0].ID,
			Nome:       "Heavy Riffs",
			Data_criacao: sql.NullTime{
				Time:  now,
				Valid: true,
			},
		},
	}

	db.Create(&playlists)

	db.Model(&playlists[0]).Association("Musicas").Append(
		&musicas[0],
		&musicas[2],
		&musicas[3],
	)

	db.Model(&playlists[1]).Association("Musicas").Append(
		&musicas[1],
	)

	db.Model(&playlists[2]).Association("Musicas").Append(
		&musicas[2],
		&musicas[5],
	)
}

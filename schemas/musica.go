package schemas

import (
	"gorm.io/gorm"
)

type Musica struct {
	gorm.Model
	Titulo           string
	Duracao_segundos uint
	Artista_id       uint
	Playlists        []Playlist `gorm:"many2many:musica_playlists"`
}

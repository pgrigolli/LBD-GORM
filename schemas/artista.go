package schemas

import (
	"database/sql"

	"gorm.io/gorm"
)

type Artista struct {
	gorm.Model
	Nome          string `gorm:"unique"`
	Nacionalidade sql.NullString
	Musicas       []Musica `gorm:"foreignKey:Artista_id;constraint:OnDelete:CASCADE"`
}

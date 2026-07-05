package schemas

import (
	"database/sql"
)

type Artista struct {
	ID            uint           `gorm:"column:id;primaryKey;autoIncrement"`
	Nome          string         `gorm:"column:nome;unique;not null"`
	Nacionalidade sql.NullString `gorm:"column:nacionalidade"`
	Musicas       []Musica       `gorm:"foreignKey:Artista_id;references:ID;constraint:OnDelete:CASCADE"`
}

func (Artista) TableName() string {
	return "ARTISTA"
}

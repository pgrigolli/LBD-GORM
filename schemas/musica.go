package schemas

type Musica struct {
	ID               uint    `gorm:"column:id;primaryKey;autoIncrement"`
	Titulo           string  `gorm:"column:titulo;not null"`
	Duracao_segundos uint    `gorm:"column:duracao_segundos;not null"`
	Artista_id       uint    `gorm:"column:artista_id;not null;index"`
	Artista          Artista `gorm:"foreignKey:Artista_id;references:ID;constraint:OnDelete:RESTRICT"`
	MusicaPlaylists  []MusicaPlaylist `gorm:"foreignKey:MusicaId;references:ID;constraint:OnDelete:CASCADE"`
}

func (Musica) TableName() string {
	return "MUSICA"
}

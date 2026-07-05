package schemas

import "time"

type Playlist struct {
	PlaylistId  uint      `gorm:"column:playlist_id;primaryKey;autoIncrement"`
	UsuarioId   uint      `gorm:"column:usuario_id;primaryKey;not null;index"`
	Nome        string    `gorm:"column:nome;not null"`
	DataCriacao time.Time `gorm:"column:data_criacao;autoCreateTime"`
	Usuario     Usuario   `gorm:"foreignKey:UsuarioId;references:ID;constraint:OnDelete:CASCADE"`
}

func (Playlist) TableName() string {
	return "PLAYLIST"
}

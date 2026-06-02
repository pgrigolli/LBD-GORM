package schemas

import "time"

type Playlist struct {
	PlaylistId  uint      `gorm:"column:playlist_id;primaryKey;autoIncrement"`
	UsuarioId   uint      `gorm:"column:usuario_id;primaryKey;not null"`
	Nome        string    `gorm:"not null"`
	DataCriacao time.Time `gorm:"column:data_criacao;autoCreateTime"`
}

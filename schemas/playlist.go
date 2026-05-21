package schemas

import (
	"database/sql"

	"gorm.io/gorm"
)

type Playlist struct {
	gorm.Model
	Usuario_id   uint `gorm:"primaryKey;autoIncrement:false"`
	Nome         string
	Data_criacao sql.NullTime
}

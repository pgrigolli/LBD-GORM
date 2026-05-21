package schemas

import (
	"gorm.io/gorm"
)

type Usuario struct {
	gorm.Model
	Username  string
	Email     string
	Playlists []Playlist
}

package schemas

type Usuario struct {
	ID        uint       `gorm:"column:id;primaryKey;autoIncrement"`
	Username  string     `gorm:"column:username;unique;not null"`
	Email     string     `gorm:"column:email;unique;not null"`
	Playlists []Playlist `gorm:"foreignKey:UsuarioId;references:ID;constraint:OnDelete:CASCADE"`
}

func (Usuario) TableName() string {
	return "USUARIO"
}

package schemas

type MusicaPlaylist struct {
	MusicaId        uint     `gorm:"column:musica_id;primaryKey"`
	PlaylistId      uint     `gorm:"column:playlist_id;primaryKey;uniqueIndex:idx_ordem_na_playlist"`
	UsuarioId       uint     `gorm:"column:usuario_id;primaryKey;uniqueIndex:idx_ordem_na_playlist"`
	OrdemNaPlaylist int      `gorm:"column:ordem_na_playlist;not null;uniqueIndex:idx_ordem_na_playlist"`
	Musica          Musica   `gorm:"foreignKey:MusicaId;constraint:OnDelete:CASCADE"`
}

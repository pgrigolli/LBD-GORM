package main

import (
	"context"
	"database/sql"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Artista struct {
	gorm.Model
	nome          string `gorm:"unique"`
	nacionalidade sql.NullString
	musicas       []Musica
}

type Usuario struct {
	gorm.Model
	username  string
	email     string
	playlists []Playlist
}

type Playlist struct {
	gorm.Model
	usuario_id   uint `gorm:"primaryKey;autoIncrement:false"`
	nome         string
	data_criacao sql.NullTime
}

type Musica struct {
	gorm.Model
	titulo           string
	duracao_segundos uint
	artista_id       uint
	Playlists        []Playlist `gorm:"many2many:musica_playlists"`
}

func main() {

	dsn := "host=localhost user=admin password=admin123 dbname=LBDGORM port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema
	db.AutoMigrate(&Artista{}, &Usuario{})

	ctx := context.Background()

	user, err := gorm.G[Artista](db).Take(ctx)

	fmt.Println(db)

	fmt.Println(user)
	fmt.Println(err)

	fmt.Println("Hello, World!")
}

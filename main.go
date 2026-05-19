package main

import (
	"context"
	"database/sql"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Artista struct {
	gorm.Model
	nome          string
	nacionalidade sql.NullString
}

type Usuario struct {
	gorm.Model
	username string
	email    string
}

type Playlist struct {
	playlist_id uint `gorm:"primaryKey"`
	usuario_id  uint `gorm:"primaryKey"`
}

func main() {

	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
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

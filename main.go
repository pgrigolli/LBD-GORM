package main

import (
	"fmt"

	"LBD/database"
	"LBD/schemas"
)

func main() {

	db, err := database.ConnectDB()

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(
		&schemas.Artista{},
		&schemas.Usuario{},
		&schemas.Musica{},
		&schemas.Playlist{},
	)

	//seed.Seed(db)

	fmt.Println("Banco populado com sucesso!")
}

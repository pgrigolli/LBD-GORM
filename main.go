package main

import (
	"LBD/database"
	"LBD/schemas"
	"LBD/services"
	"fmt"
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

	//fmt.Println("Banco populado com sucesso!")

	// artista := schemas.Artista{
	// 	Nome:          "Laufey",
	// 	Nacionalidade: sql.NullString{String: "", Valid: false},
	// 	Musicas:       []schemas.Musica{}}

	// err = services.CreateArtista(artista)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("Artista criado com sucesso")
	// }

	artistas, err := services.GetAllArtista()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(artistas[1])
	}
}

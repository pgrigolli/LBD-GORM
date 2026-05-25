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
		for index, element := range artistas {

			fmt.Println("Artista %d", index)
			fmt.Println(element)
			fmt.Println("")
		}
	}

	// //Exemplo de UpdateArtista
	// artistaUpdate := schemas.Artista{
	// 	Model: gorm.Model{ID: 1}, // ajuste o ID conforme necessário
	// 	Nome:  "Laufey Updated",
	// }

	// _, err = services.UpdateArtista(artistaUpdate)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("Artista atualizado com sucesso")
	// }

	// //Exemplo de DeleteArtista
	// err = services.DeleteArtista(1) // ajuste o ID conforme necessário
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("Artista deletado com sucesso")
	// }

}

package services

import (
	"LBD/database"
	"LBD/schemas"
	"context"

	"gorm.io/gorm"
)

func GetAllUsuario() ([]schemas.Usuario, error) {
	db, err := database.ConnectDB()

	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	return gorm.G[schemas.Usuario](db.Debug()).Find(ctx)
}

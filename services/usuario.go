package services

import (
	"LBD/schemas"
	"context"

	"gorm.io/gorm"
)

func GetAllUsuario() ([]schemas.Usuario, error) {
	db, err := connectDB()

	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	return gorm.G[schemas.Usuario](db.Debug()).Find(ctx)
}

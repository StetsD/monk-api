package repositories

import (
	"context"
	"fmt"
	"github.com/stetsd/monk-api/internal/app/schemas"
	"github.com/stetsd/monk-api/internal/domain/infrasInterfaces"
	"github.com/stetsd/monk-api/internal/domain/models"
	"github.com/stetsd/monk-api/internal/infrastructure/logger"
)

type PgRepoUserStore struct {
	pgd infrasInterfaces.DbDriver
}

func NewPgRepoUserStore(pgDriver infrasInterfaces.DbDriver) *PgRepoUserStore {
	return &PgRepoUserStore{pgDriver}
}

func (pgRUS *PgRepoUserStore) Registration(data *schemas.RegistrationBody) (int, error) {
	// TODO: password not salted
	rows, err := pgRUS.pgd.Query(`
		INSERT INTO "User" (name, email, password) 
		VALUES ($1, $2, $3) RETURNING id`,
		data.Name,
		data.Email,
		data.Password,
	)

	defer func() {
		if err := rows.Close(); err != nil {
			logger.Log.Error(err.Error())
		}
	}()

	if err != nil {
		return 0, err
	}

	var id int

	for rows.Next() {
		if err := rows.Scan(&id); err != nil {
			logger.Log.Error(err.Error())
		}
	}

	return id, nil
}

func (pgRUS *PgRepoUserStore) Login(_ context.Context, _ *models.User) error {
	fmt.Println("plug")
	return nil
}

func (pgRUS *PgRepoUserStore) Logout(ctx context.Context, id int) error {
	fmt.Println("plug")
	return nil
}

func (pgRUS *PgRepoUserStore) Put(ctx context.Context, user *models.User) (*models.User, error) {
	fmt.Println("plug")
	return &models.User{}, nil
}

func (pgRUS *PgRepoUserStore) Delete(ctx context.Context, id int) error {
	fmt.Println("plug")
	return nil
}

func (pgRUS *PgRepoUserStore) GetById(ctx context.Context, id int) (*models.User, error) {
	fmt.Println("plug")
	return &models.User{}, nil
}

package repositories

import (
	"context"
	"fmt"
	"github.com/stetsd/monk-api/internal/app/schemas"
	"github.com/stetsd/monk-api/internal/domain/infrasInterfaces"
	"github.com/stetsd/monk-api/internal/domain/models"
)

type PgRepoUserStore struct {
	pgd infrasInterfaces.DbDriver
}

func NewPgRepoUserStore(pgDriver infrasInterfaces.DbDriver) *PgRepoUserStore {
	return &PgRepoUserStore{pgDriver}
}

func (pgRUS *PgRepoUserStore) Registration(data *schemas.RegistrationBody) error {
	// TODO: password not salted
	err := pgRUS.pgd.Query(`
		INSERT INTO "User" (name, email, password) 
		VALUES ($1, $2, $3)`,
		data.Name,
		data.Email,
		data.Password,
	)

	if err != nil {
		return err
	}

	return nil
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

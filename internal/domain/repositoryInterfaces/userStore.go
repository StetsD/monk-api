package repositoryInterfaces

import (
	"context"
	"github.com/stetsd/monk-api/internal/app/schemas"
	"github.com/stetsd/monk-api/internal/domain/models"
)

type UserStore interface {
	Login(ctx context.Context, user *models.User) error
	Registration(data *schemas.RegistrationBody) error
	Logout(ctx context.Context, id int) error
	Put(ctx context.Context, user *models.User) (*models.User, error)
	Delete(ctx context.Context, id int) error
	GetById(ctx context.Context, id int) (*models.User, error)
}

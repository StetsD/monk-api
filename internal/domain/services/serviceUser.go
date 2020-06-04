package services

import (
	"github.com/stetsd/monk-api/internal/app/schemas"
	"github.com/stetsd/monk-api/internal/domain/repositoryInterfaces"
)

type ServiceUser struct {
	UserStore repositoryInterfaces.UserStore
}

const ServiceUserName = "ServiceUser"

func (su *ServiceUser) Login() error {

	return nil
}

func (su *ServiceUser) Registration(data *schemas.RegistrationBody) (int, error) {
	id, err := su.UserStore.Registration(data)

	if err != nil {
		return 0, err
	}

	return id, nil
}

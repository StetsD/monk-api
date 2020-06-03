package tools

import (
	"github.com/stetsd/monk-api/internal/domain/repositories"
	"github.com/stetsd/monk-api/internal/domain/services"
	"github.com/stetsd/monk-db-driver"
)

type ServiceCollection struct {
	ServiceUser  services.ServiceUser
	ServiceEvent services.ServiceEvent
}

func Bind(driver *monk_db_driver.DbDriver, serviceNames ...string) ServiceCollection {
	serviceCollection := ServiceCollection{}
	for _, service := range serviceNames {
		switch service {
		case services.ServiceUserName:
			pgRepoUserStore := repositories.NewPgRepoUserStore(driver)
			serviceCollection.ServiceUser = services.ServiceUser{
				UserStore: pgRepoUserStore,
			}
		case services.ServiceEventName:
			serviceCollection.ServiceEvent = services.ServiceEvent{}
		}
	}

	return serviceCollection
}

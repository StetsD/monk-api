package services

import (
	"github.com/stetsd/monk-api/internal/api"
	"github.com/stetsd/monk-api/internal/app/schemas"
	"github.com/stetsd/monk-api/internal/domain/repositoryInterfaces"
)

type ServiceEvent struct {
	GrpcConn repositoryInterfaces.GrpcClient
}

const ServiceEventName = "ServiceEvent"

func (se *ServiceEvent) SendEvent(msg *schemas.EventBody) (*api.EventResult, error) {
	eventResult, err := se.GrpcConn.SendEvent(msg)
	if err != nil {
		return nil, err
	}

	return eventResult, nil
}

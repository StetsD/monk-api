package services

import (
	"github.com/stetsd/monk-api/internal/domain/repositoryInterfaces"
)

type ServiceEvent struct {
	GrpcConn repositoryInterfaces.GrpcClient
}

const ServiceEventName = "ServiceEvent"

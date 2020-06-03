package repositoryInterfaces

import (
	"github.com/stetsd/monk-api/internal/api"
	"github.com/stetsd/monk-api/internal/app/schemas"
)

type GrpcClient interface {
	SendEvent(msg *schemas.EventBody) (*api.EventResult, error)
	Close() error
}

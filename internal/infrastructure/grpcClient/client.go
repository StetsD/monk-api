package grpcClient

import (
	"context"
	"github.com/golang/protobuf/ptypes"
	"github.com/stetsd/monk-api/internal/api"
	"github.com/stetsd/monk-api/internal/app/schemas"
	"google.golang.org/grpc"
)

type GrpcConn struct {
	connection *grpc.ClientConn
	apiClient  api.ApiClient
}

func NewGrpcConn() (*GrpcConn, error) {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		return nil, err
	}

	apiClient := api.NewApiClient(connection)

	return &GrpcConn{connection: connection, apiClient: apiClient}, nil
}

func (c *GrpcConn) SendEvent(msg *schemas.EventBody) (*api.EventResult, error) {
	dateStart, err := ptypes.TimestampProto(msg.DateStart)
	if err != nil {
		return nil, err
	}
	dateEnd, err := ptypes.TimestampProto(msg.DateEnd)
	if err != nil {
		return nil, err
	}

	eventResult, err := c.apiClient.SendEvent(context.Background(), &api.Event{
		Title:       msg.Title,
		Description: msg.Description,
		UserId:      msg.UserId,
		DateStart:   dateStart,
		DateEnd:     dateEnd,
	})

	if err != nil {
		return nil, err
	}

	return eventResult, nil
}

func (c *GrpcConn) Close() error {
	err := c.connection.Close()
	if err != nil {
		return err
	}
	return nil
}

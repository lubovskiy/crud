package service

import (
	"context"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/lubovskiy/crud/pkg/crud"
)

type ParcelChangesService struct {
}

func NewParcelChangesService() *ParcelChangesService {
	return &ParcelChangesService{
	}
}

func (p ParcelChangesService) ListContacts(ctx context.Context, request *crud.ListContactsRequest) (*crud.ListContactsResponse, error) {
	log, _ := zap.NewProduction()
	defer log.Sync()

	log.Debug("gRPC: Run ListContacts")

	return nil, errors.New("oooo")
}

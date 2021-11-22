package service

import (
	"context"

	"github.com/lubovskiy/crud/pkg/crud"
)

type ParcelChangesService struct {
}

func NewParcelChangesService() *ParcelChangesService {
	return &ParcelChangesService{
	}
}

func (p ParcelChangesService) ListContacts(ctx context.Context, request *crud.ListContactsRequest) (*crud.ListContactsResponse, error) {
	panic("implement me")
}

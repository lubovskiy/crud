package service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/lubovskiy/crud/infrastructure/database"
	phonebookRepo "github.com/lubovskiy/crud/internal/repository/phonebook"
	"github.com/lubovskiy/crud/pkg/crud"
)

var (
	idToField = map[int32]string{
		0: "id",
		1: "name",
		2: "phone",
	}
)

type PhonebookService struct {
	repo phonebookRepo.Repositer
}

func NewPhonebookService(repo phonebookRepo.Repositer) *PhonebookService {
	return &PhonebookService{
		repo: repo,
	}
}

func (p *PhonebookService) ListContacts(ctx context.Context, request *crud.ListContactsRequest) (*crud.ListContactsResponse, error) {
	var filter *phonebookRepo.Filter
	if request.Filter != nil {
		filter = &phonebookRepo.Filter{
			ID:    request.Filter.Ids,
			Name:  request.Filter.Names,
			Phone: request.Filter.Phones,
		}
	}

	params := &database.Params{
		ReturnedFields: getReturnFields(request.GetReturnedFields()),
	}
	if request.Limit != nil {
		params.Limits = &request.Limit.Value
	}

	res, err := p.repo.GetList(ctx, filter, params)
	if err != nil {
		return nil, status.Error(codes.NotFound, "id was not found")
	}

	contacts := toContacts(res)

	return &crud.ListContactsResponse{
		Contacts: contacts,
	}, nil
}

func getReturnFields(returnedFields []crud.ContactFieldNum) []string {
	res := make([]string, 0)
	for _, v := range returnedFields {
		if f, ok := idToField[int32(v)]; ok {
			res = append(res, f)
		}
	}
	return res
}

func toContacts(r []*phonebookRepo.Model) []*crud.Contact {
	res := make([]*crud.Contact, len(r))

	for i, v := range res {
		res[i] = &crud.Contact{
			Id:    v.Id,
			Name:  v.Name,
			Phone: v.Phone,
		}
	}

	return res
}

func (p *PhonebookService) AddContact(ctx context.Context, request *crud.AddContactRequest) (*crud.IsErr, error) {
	panic("implement me")
}

func (p *PhonebookService) UpdateContact(ctx context.Context, request *crud.UpdateContactRequest) (*crud.IsErr, error) {
	panic("implement me")
}

func (p *PhonebookService) DeleteContact(ctx context.Context, request *crud.DeleteContactRequest) (*crud.IsErr, error) {
	panic("implement me")
}

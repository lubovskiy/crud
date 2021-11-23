package service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/golang/protobuf/ptypes/wrappers"
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
	err := request.Validate()
	if err != nil {
		return nil, err
	}

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

	for i, v := range r {
		res[i] = &crud.Contact{
			Id:    &wrappers.Int64Value{Value: v.ID},
			Name:  &wrappers.StringValue{Value: v.Name},
			Phone: &wrappers.StringValue{Value: v.Phone},
		}
	}

	return res
}

func (p *PhonebookService) AddContact(ctx context.Context, request *crud.AddContactRequest) (*crud.IsErr, error) {
	err := request.Validate()
	if err != nil {
		return nil, err
	}

	model := &phonebookRepo.Model{
		Name:  request.Name.Value,
		Phone: request.Phone.Value,
	}

	_, err = p.repo.Upsert(ctx, model)
	if err != nil {
		return &crud.IsErr{
			Err: &wrappers.StringValue{Value: err.Error()},
		}, nil
	}

	return &crud.IsErr{}, nil
}

func (p *PhonebookService) UpdateContact(ctx context.Context, request *crud.UpdateContactRequest) (*crud.IsErr, error) {
	err := request.Validate()
	if err != nil {
		return nil, err
	}

	model := &phonebookRepo.Model{
		ID: request.Id.Value,
	}

	if request.Name != nil {
		model.Name = request.Name.Value
	}
	if request.Phone != nil {
		model.Phone = request.Phone.Value
	}

	err = p.repo.Update(ctx, model)
	if err != nil {
		return &crud.IsErr{
			Err: &wrappers.StringValue{Value: err.Error()},
		}, nil
	}

	return &crud.IsErr{}, nil
}

func (p *PhonebookService) DeleteContact(ctx context.Context, request *crud.DeleteContactRequest) (*crud.IsErr, error) {
	err := request.Validate()
	if err != nil {
		return nil, err
	}

	filter := &phonebookRepo.Filter{
		ID: []int64{request.Id.Value},
	}

	err = p.repo.Delete(ctx, filter)
	if err != nil {
		return &crud.IsErr{
			Err: &wrappers.StringValue{Value: err.Error()},
		}, nil
	}

	return &crud.IsErr{}, nil
}

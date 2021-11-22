package phonebook

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/lubovskiy/crud/infrastructure/database"
)

var (
	ErrNil = errors.New("model can't be nil")
)

type Repositer interface {
	Upsert(ctx context.Context, model *Model) (*Model, error)
	GetList(ctx context.Context, filter *Filter, params *database.Params) ([]*Model, error)
	Delete(ctx context.Context, filter *Filter) error
}

type Repository struct {
	Pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{pool}
}

func (r *Repository) Upsert(ctx context.Context, model *Model) (*Model, error) {
	if model == nil {
		return nil, ErrNil
	}

	query := `
		INSERT INTO phonebook
			(name, phone)
		VALUES
			($1, $2)
		ON CONFLICT ON CONSTRAINT phonebook_pkey
			DO UPDATE SET LOAD = EXCLUDED.load
		RETURNING ID`


	arguments := []interface{}{(*model).Name, (*model).Phone}

	var id int64
	err := r.Pool.QueryRow(
		ctx, query, arguments...,
	).Scan(
		&id,
	)
	if err != nil {
		return nil, err
	}

	model.ID = id

	return model, nil
}

func (r *Repository) GetList(ctx context.Context, filter *Filter, params *database.Params) ([]*Model, error) {
	fields := []string{
		"id",
		"name",
		"phone",
	}

	limits, retFields := params.Gen(fields)

	qRetFields := strings.Join(retFields, ", ")
	query := fmt.Sprintf("SELECT %s FROM phonebook", qRetFields)

	conditions, arguments := filter.Filter()

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND\n ")
	}

	if len(limits) > 0 {
		query += strings.Join(limits, " ")
	}

	rows, err := r.Pool.Query(ctx, query, arguments...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	list := make([]*Model, 0)
	for rows.Next() {
		var m Model
		err = rows.Scan(
			&m.ID,
			&m.Name,
			&m.Phone,
		)
		if err != nil {
			return nil, err
		}

		list = append(list, &m)
	}

	return list, nil
}


func (r *Repository) Delete(ctx context.Context, filter *Filter) error {
	query := "DELETE FROM phonebook \n"

	conditions, arguments := filter.Filter()

	if len(conditions) > 0 {
		query += " WHERE "
		query += strings.Join(conditions, " AND\n ")
	}

	_, err := r.Pool.Exec(ctx, query, arguments...)
	if err != nil {
		return err
	}

	return nil
}

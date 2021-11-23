package phonebook

import (
	"fmt"

	"github.com/lubovskiy/crud/helpers/bind"
)

type Model struct {
	ID    int64
	Name  string
	Phone string
}

type Filter struct {
	ID    []int64
	Name  []string
	Phone []string
}

func (f *Filter) Filter() (conditions []string, arguments []interface{}) {
	gen := bind.SequentialGenerator{}
	gen.Next()

	if f.ID != nil {
		conditions = append(conditions, fmt.Sprintf("id = ANY(%s::bigint[])", gen.NextBind()))
		arguments = append(arguments, f.ID)
	}

	if f.Name != nil {
		conditions = append(conditions, fmt.Sprintf("name = ANY(%s::text[])", gen.NextBind()))
		arguments = append(arguments, f.Name)
	}

	if f.Phone != nil {
		conditions = append(conditions, fmt.Sprintf("phone = ANY(%s::text[])", gen.NextBind()))
		arguments = append(arguments, f.Phone)
	}

	return
}

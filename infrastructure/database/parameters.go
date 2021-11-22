package database

import (
	"fmt"
	"github.com/lubovskiy/crud/helpers/intersect"
)

type Params struct {
	ReturnedFields []string
	Limits         *int
}

func (p *Params) Gen(fields []string) (limits []string, retFields []string) {
	if len(p.ReturnedFields) == 0 {
		retFields = fields
	} else {
		retFields = intersect.IntersectStr(p.ReturnedFields, fields)
	}

	if p.Limits != nil && *p.Limits > 0 {
		limits = []string{fmt.Sprintf(" LIMIT %d ", *p.Limits)}
	}

	return
}

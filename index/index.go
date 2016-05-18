package index

import (
	"fmt"
	"github.com/csimplestring/go-mem-store/document"
)

type Index interface {
	Search(op Op, args ...interface{}) ([]document.ObjectID, error)
	Field() string
}

type Op string

const (
	OpEq    = "eq"
	OpGt    = "gt"
	OpGte   = "gte"
	OpLt    = "lt"
	OpLte   = "lte"
	OpRange = "range"
)

type IndexManager struct {
	indices map[string]Index
}

func (i *IndexManager) AddIndex(idx Index) error {
	if _, exist := i.indices[idx.Field()]; exist {
		return fmt.Errorf("Duplicate index on field %s", idx.Field())
	}

	i.indices[idx.Field()] = idx
	return nil
}

func (i *IndexManager) FindIndexByField(field string) Index {
	return i.indices[field]
}

func NewManager() *IndexManager {
	return &IndexManager{
		indices: make(map[string]Index),
	}
}

package index

import (
	"github.com/csimplestring/go-mem-store/document"
	"hash"
)

type Index interface {
	Search(op Op, args ...interface{}) ([]document.ObjectID, error)
	Fields() []string
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
	hash    hash.Hash
}

func (i *IndexManager) AddIndex(idx Index) {
	key := i.getKeyFor(idx.Fields())
	i.indices[key] = idx
}

func (i *IndexManager) FindIndexByFields(fields ...string) Index {
	for i = len(fields); i>=0; i-- {
		key := i.getKeyFor(fields[:i])
		if idx, ok := i.indices[key]; ok {
			return idx
		}
	}

	return nil
}

func (i *IndexManager) getKeyFor(fields ...string) string {
	i.hash.Reset()
	for _, f := range fields {
		i.hash.Write([]byte(f))
	}

	return string(i.hash.Sum(nil))
}

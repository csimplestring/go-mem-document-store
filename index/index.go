package index

import (
	"github.com/csimplestring/go-mem-store/document"
)

type Index interface {
	Find(key interface{}) ([]document.ObjectID, error)

	FindGreaterThan(key interface{}) ([]document.ObjectID, error)

	FindGreaterThanEqual(key interface{}) ([]document.ObjectID, error)

	FindLessThan(key interface{}) ([]document.ObjectID, error)

	FindLessThanEqual(key interface{}) ([]document.ObjectID, error)

	FindRange(start, end interface{}) ([]document.ObjectID, error)

	Fields() []string

	Name() string

	Match(fields... string) bool

	Type() string

	PrefixField() string
}
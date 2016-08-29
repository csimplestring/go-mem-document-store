package document

import (
	"fmt"
	"strconv"
	"errors"
)

// Document represents a scheme-less data-structure, whose field value can be retrieved
// by key path.
type Document interface  {
	ID() ObjectID
	Get(key string) interface{}
}

type ObjectID interface {
	Bytes() []byte
}

type objectID struct  {
	raw []byte
}

func (o *objectID) Bytes() []byte {
	return o.raw
}


type doc struct {
	id   *objectID
	data map[string]interface{}
}

func (d *doc) ID() ObjectID {
	return d.id
}

func (d *doc) Get(key string) interface{} {
	return d.data[key]
}

func NewFromMap(m map[string]interface{}) (*doc, error) {

	if _, exist := m["id"]; !exist {
		return nil, errors.New("The document id is missing.")
	}

	id, ok := m["id"].([]byte)
	if !ok {
		return nil, errors.New("The document id must be []byte.")
	}

	return &doc{
		id : &objectID{raw: id},
		data: scanMap(m, ""),
	}, nil
}

func scanMap(m map[string]interface{}, parent string) (map[string]interface{}) {
	flatten := make(map[string]interface{})

	for key, value := range m {
		if len(parent) > 0 {
			key = parent + "." + key
		}

		switch v := value.(type) {

		case []interface{}:
			flatten = mergeMap(flatten, scanArray(v, key))
		case map[string]interface{}:
			flatten = mergeMap(flatten, scanMap(v, key))
		default:
			flatten[key] = v
		}
	}
	return flatten
}

func scanArray(l []interface{}, parent string) (map[string]interface{}) {

	flatten := make(map[string]interface{})

	for i, value := range l {
		var key string

		if len(parent) > 0 {
			key = fmt.Sprintf("%s[%s]", parent, strconv.Itoa(i))
		} else {
			key = strconv.Itoa(i)
		}

		switch v := value.(type) {

		case []interface{}:
			flatten = mergeMap(flatten, scanArray(v, key))
		case map[string]interface{}:
			flatten = mergeMap(flatten, scanMap(v, key))
		default:
			flatten[key] = v
		}
	}
	return flatten
}

func mergeMap(dst, merged map[string]interface{}) map[string]interface{} {
	for k, v := range merged {
		dst[k] = v
	}
	return dst
}

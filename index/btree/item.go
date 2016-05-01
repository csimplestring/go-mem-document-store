package btree

import (
	"fmt"
	"github.com/csimplestring/go-mem-store/document"
	"github.com/petar/GoLLRB/llrb"
	"github.com/csimplestring/go-mem-store/types"
)


// btreeItem represents a single object in the b-tree, including the key and id list.
type btreeItem struct {
	key interface{}
	ids []document.ObjectID
}

// Less defines the comparison of item. Note that the type of key also participates the comparison.
// Therefore the mixed type items are supported, namely, the btree can store int-key, float-key, string-key etc.
func (b *btreeItem) Less(than llrb.Item) bool {
	other := than.(*btreeItem)

	lType := types.Of(b.key)
	rType := types.Of(other.key)
	if lType != rType {
		return lType < rType
	}

	return lessThan(b.key, other.key)
}

// createItem creates a new btreeItem and returns a pointer to it.
func createItem(key interface{}, ids []document.ObjectID) (*btreeItem, error) {
	if types.Of(key) == types.Unsupported {
		return nil, fmt.Errorf("Unsupported type for btree item")
	}

	return &btreeItem{
		key: key,
		ids: ids,
	}, nil
}

// lessThan returns True if l < r, otherwise returns False. Note that only compare the same type value!
func lessThan(l, r interface{}) bool {
	switch l.(type) {
	case int:
		return l.(int) < r.(int)
	case float32:
		return l.(float32) < r.(float32)
	case float64:
		return l.(float64) < r.(float64)
	case string:
		return l.(string) < r.(string)
	default:
		return false
	}
}

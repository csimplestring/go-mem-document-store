package btree

import (
	"github.com/csimplestring/go-mem-store/document"
	"github.com/petar/GoLLRB/llrb"
	"github.com/csimplestring/go-mem-store/types"

)

// iterator implements btree.ItemIterator interface,
// store passed item's object ids in results.
type iterator struct {
	results  []document.ObjectID
	itemType types.T
	filter   func(item llrb.Item) bool
}

// visit is usually used as a closure passed to btree's AscendXXX method,
// return True if the iteration continues, otherwise returns False.
func (b *iterator) visit(item llrb.Item) bool {
	node := item.(*btreeItem)

	if types.Of(node.key) != b.itemType {
		return false
	}

	if b.filter != nil && b.filter(item) {
		return true
	}

	b.results = append(b.results, node.ids...)
	return true
}

// setType tells the iterator only accepts items of a specific type.
func (b *iterator) setType(t types.T) *iterator {
	b.itemType = t
	return b
}

// setFilter add a filter closure to decide if the item will be rejected or not during traversal of btree.
func (b *iterator) setFilter(filter func(item llrb.Item) bool) *iterator {
	b.filter = filter
	return b
}

// buildIterator builds a new iterator.
func buildIterator() *iterator {
	return &iterator{}
}



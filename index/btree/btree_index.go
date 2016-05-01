package btree

import (
	"fmt"
	"github.com/csimplestring/go-mem-store/document"
	"github.com/csimplestring/go-mem-store/types"
	"github.com/petar/GoLLRB/llrb"
)

// btreeIndex represents an index using b-tree.
type btreeIndex struct {
	field string
	name  string
	btree *llrb.LLRB
}

// Insert takes a new document and add a new btree entry for it.
func (b *btreeIndex) Insert(doc document.Document) error {
	key, err := createItem(doc.Get(b.field), nil)
	if err != nil {
		return err
	}

	newItem, err := createItem(doc.Get(b.field), []document.ObjectID{doc.ID()})
	if err != nil {
		return err
	}

	found := b.btree.Get(key)
	if found != nil {
		list := found.(*btreeItem).ids
		newItem.ids = append(list, newItem.ids...)
	}

	b.btree.ReplaceOrInsert(newItem)

	return nil
}

// FindLessThan finds all the documents whose field is between the start and the end key.
// Only the ObjectIDs of documents are returned.
func (b *btreeIndex) FindRange(start, end interface{}) ([]document.ObjectID, error) {
	itemType := types.Of(start)
	if itemType != types.Of(end) {
		return nil, fmt.Errorf("The type of Start and End in Range is not the same.")
	}

	startItem, err := createItem(start, nil)
	if err != nil {
		return nil, err
	}

	endItem, err := createItem(end, nil)
	if err != nil {
		return nil, err
	}

	iter := buildIterator().setType(itemType)
	b.btree.AscendRange(startItem, endItem, iter.visit)

	return iter.results, nil
}

// FindLessThan finds all the documents whose field is less than the key.
// Only the ObjectIDs of documents are returned.
func (b *btreeIndex) FindLessThan(key interface{}) ([]document.ObjectID, error) {
	pivot, err := createItem(key, nil)
	if err != nil {
		return nil, err
	}

	filter := func(item llrb.Item) bool {
		node := item.(*btreeItem)
		shouldSkip := (node.key == key)

		return shouldSkip
	}

	iter := buildIterator().setType(types.Of(key)).setFilter(filter)
	b.btree.DescendLessOrEqual(pivot, iter.visit)

	return iter.results, nil
}

// FindLessThanEqual finds all the documents whose field is less than or equal to the key.
// Only the ObjectIDs of documents are returned.
func (b *btreeIndex) FindLessThanEqual(key interface{}) ([]document.ObjectID, error) {
	pivot, err := createItem(key, nil)
	if err != nil {
		return nil, err
	}

	iter := buildIterator().setType(types.Of(key))
	b.btree.DescendLessOrEqual(pivot, iter.visit)

	return iter.results, nil
}

// FindGreaterThan finds all the documents whose field is greater than the key.
// Only the ObjectIDs of documents are returned.
func (b *btreeIndex) FindGreaterThan(key interface{}) ([]document.ObjectID, error) {
	pivot, err := createItem(key, nil)
	if err != nil {
		return nil, err
	}

	filter := func(item llrb.Item) bool {
		node := item.(*btreeItem)
		shouldSkip := (node.key == key)

		return shouldSkip
	}

	iter := buildIterator().setType(types.Of(key)).setFilter(filter)
	b.btree.AscendGreaterOrEqual(pivot, iter.visit)

	return iter.results, nil
}

// FindGreaterThanEqual finds all the documents whose field is greater than or equal the key.
// Only the ObjectIDs of documents are returned.
func (b *btreeIndex) FindGreaterThanEqual(key interface{}) ([]document.ObjectID, error) {
	pivot, err := createItem(key, nil)
	if err != nil {
		return nil, err
	}

	iter := buildIterator().setType(types.Of(key))
	b.btree.AscendGreaterOrEqual(pivot, iter.visit)

	return iter.results, nil
}

// Find finds all the documents whose field is equal to the key.
// Only the ObjectIDs of documents are returned.
func (b *btreeIndex) Find(value interface{}) ([]document.ObjectID, error) {
	var results []document.ObjectID

	key, err := createItem(value, nil)
	if err != nil {
		return results, err
	}

	item := b.btree.Get(key)
	if item != nil {
		results = item.(*btreeItem).ids
	}

	return results, nil
}

// New a b-tree index.
func New(name, field string) *btreeIndex {
	return &btreeIndex{
		field: field,
		name:  name,
		btree: llrb.New(),
	}
}



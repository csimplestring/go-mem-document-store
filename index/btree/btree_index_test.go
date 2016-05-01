package btree

import (
	"bytes"
	"github.com/csimplestring/go-mem-store/document"
	"github.com/stretchr/testify/assert"
	"testing"
	"strconv"
)

type doc struct {
	id   *bytes.Buffer
	data map[string]interface{}
}

func (d *doc) ID() document.ObjectID {
	return d.id
}

func (d *doc) Get(key string) interface{} {
	return d.data[key]
}

func newTestingDoc(id, field string, v interface{}) *doc {
	return &doc{
		id: bytes.NewBufferString(id),
		data: map[string]interface{}{
			field: v,
		},
	}
}

func getTestingMixedTypeDocs() []*doc {
	return []*doc{
		newTestingDoc("1", "f1", 1),
		newTestingDoc("2", "f1", "a"),
		newTestingDoc("3", "f1", 1.1),
		newTestingDoc("4", "f1", 2.3),
		newTestingDoc("5", "f1", 2),
		newTestingDoc("6", "f1", "b"),
		newTestingDoc("7", "f1", 1),
		newTestingDoc("8", "f1", "c"),
	}
}

func getTestingSequentialDocs(start, end int) []*doc {
	var docs []*doc
	for i:=start; i<end; i++ {
		docs = append(docs, newTestingDoc(strconv.Itoa(i), "f1", i))
	}

	return docs
}

func TestInsert_OK(t *testing.T) {
	index := New("f1-index", "f1")

	for _, d := range getTestingMixedTypeDocs() {
		err := index.Insert(d)
		assert.NoError(t, err)
	}
}

func TestFindRange_OK(t *testing.T) {
	index := New("f1-index", "f1")

	for _, d := range getTestingSequentialDocs(0,100) {
		err := index.Insert(d)
		assert.NoError(t, err)
	}

	actual, err := index.FindRange(3, 6)
	assert.NoError(t, err)

	expected := [][]byte {
		[]byte("3"), []byte("4"), []byte("5"),
	}

	assertEqualResults(expected, actual, t)
}

func TestFindRange_NotFound(t *testing.T) {
	index := New("f1-index", "f1")

	for _, d := range getTestingSequentialDocs(0,100) {
		err := index.Insert(d)
		assert.NoError(t, err)
	}

	actual, err := index.FindRange(101, 200)
	assert.NoError(t, err)
	assert.Empty(t, actual)
}

func TestFindByLessThanEqual_NotFound(t *testing.T) {
	index := New("f1-index", "f1")

	for _, d := range getTestingMixedTypeDocs() {
		err := index.Insert(d)
		assert.NoError(t, err)
	}

	tests := []struct {
		key      interface{}
		expected [][]byte
	}{
		{
			key: 0,
		},
		{
			key: 1.0,
		},
		{
			key: "",
		},
	}

	for _, test :=range tests {
		actual, err := index.FindLessThanEqual(test.key)
		assert.NoError(t, err)
		assert.Empty(t, actual)
	}
}

func TestFindByLessThanEqual_OK(t *testing.T) {
	index := New("f1-index", "f1")

	for _, d := range getTestingMixedTypeDocs() {
		err := index.Insert(d)
		assert.NoError(t, err)
	}

	tests := []struct {
		key      interface{}
		expected [][]byte
	}{
		{
			key: 2,
			expected: [][]byte{
				[]byte("5"),[]byte("1"), []byte("7"),
			},
		},
		{
			key: "c",
			expected: [][]byte{
				[]byte("8"), []byte("6"), []byte("2"),
			},
		},
		{
			key: 2.3,
			expected: [][]byte{
				[]byte("4"), []byte("3"),
			},
		},
	}

	for _, test :=range tests {
		actual, err := index.FindLessThanEqual(test.key)
		assert.NoError(t, err)

		assertEqualResults(test.expected, actual, t)
	}
}

func TestFindByLessThan_NotFound(t *testing.T) {
	index := New("f1-index", "f1")

	for _, d := range getTestingMixedTypeDocs() {
		err := index.Insert(d)
		assert.NoError(t, err)
	}

	tests := []struct {
		key      interface{}
		expected [][]byte
	}{
		{
			key: 1,
		},
		{
			key: 1.1,
		},
		{
			key: "a",
		},
	}

	for _, test :=range tests {
		actual, err := index.FindLessThan(test.key)
		assert.NoError(t, err)
		assert.Empty(t, actual)
	}
}

func TestFindByLessThan_OK(t *testing.T) {
	index := New("f1-index", "f1")

	for _, d := range getTestingMixedTypeDocs() {
		err := index.Insert(d)
		assert.NoError(t, err)
	}

	tests := []struct {
		key      interface{}
		expected [][]byte
	}{
		{
			key: 2,
			expected: [][]byte{
				[]byte("1"), []byte("7"),
			},
		},
		{
			key: "c",
			expected: [][]byte{
				[]byte("6"), []byte("2"),
			},
		},
		{
			key: 2.3,
			expected: [][]byte{
				[]byte("3"),
			},
		},
	}

	for _, test :=range tests {
		actual, err := index.FindLessThan(test.key)
		assert.NoError(t, err)
		
		assertEqualResults(test.expected, actual, t)
	}
}

func TestFindByGreaterThan_NotFound(t *testing.T) {
	index := New("f1-index", "f1")

	for _, d := range getTestingMixedTypeDocs() {
		err := index.Insert(d)
		assert.NoError(t, err)
	}

	tests := []struct {
		key interface{}
	}{
		{
			key: 100,
		},
		{
			key: 2.33,
		},
		{
			key: "d",
		},
	}

	for _, test := range tests {
		actual, err := index.FindGreaterThan(test.key)
		assert.NoError(t, err)
		assert.Nil(t, actual)
	}
}

func TestFindByGreaterThan_OK(t *testing.T) {
	index := New("f1-index", "f1")

	for _, d := range getTestingMixedTypeDocs() {
		err := index.Insert(d)
		assert.NoError(t, err)
	}

	tests := []struct {
		key      interface{}
		expected [][]byte
	}{
		{
			key: 1,
			expected: [][]byte{
				[]byte("5"),
			},
		},
		{
			key: "b",
			expected: [][]byte{
				[]byte("8"),
			},
		},
		{
			key: 1.1,
			expected: [][]byte{
				[]byte("4"),
			},
		},
	}

	for _, test :=range tests {
		actual, err := index.FindGreaterThan(test.key)
		assert.NoError(t, err)

		assertEqualResults(test.expected, actual, t)
	}
}

func TestFindByGreaterThanEqual_NotFound(t *testing.T) {
	index := New("f1-index", "f1")

	for _, d := range getTestingMixedTypeDocs() {
		err := index.Insert(d)
		assert.NoError(t, err)
	}

	tests := []struct {
		key interface{}
	}{
		{
			key: 100,
		},
		{
			key: 2.33,
		},
		{
			key: "d",
		},
	}

	for _, test := range tests {
		actual, err := index.FindGreaterThanEqual(test.key)
		assert.NoError(t, err)
		assert.Nil(t, actual)
	}
}

func TestFindByGreaterThanEqual_OK(t *testing.T) {
	index := New("f1-index", "f1")

	for _, d := range getTestingMixedTypeDocs() {
		err := index.Insert(d)
		assert.NoError(t, err)
	}

	tests := []struct {
		key      interface{}
		expected [][]byte
	}{
		{
			key: 1,
			expected: [][]byte{
				[]byte("1"), []byte("7"), []byte("5"),
			},
		},
		{
			key: "b",
			expected: [][]byte{
				[]byte("6"), []byte("8"),
			},
		},
		{
			key: 1.01,
			expected: [][]byte{
				[]byte("3"), []byte("4"),
			},
		},
	}

	for _, test :=range tests {
		actual, err := index.FindGreaterThanEqual(test.key)
		assert.NoError(t, err)

		assertEqualResults(test.expected, actual, t)
	}
}

func TestFindByEqual_NotFound(t *testing.T) {
	index := New("f1-index", "f1")

	for _, d := range getTestingMixedTypeDocs() {
		err := index.Insert(d)
		assert.NoError(t, err)
	}

	tests := []struct {
		key interface{}
	}{
		{
			key: 100,
		},
		{
			key: 1.111,
		},
		{
			key: "abc",
		},
	}

	for _, test := range tests {
		actual, err := index.Find(test.key)
		assert.NoError(t, err)
		assert.Nil(t, actual)
	}
}

func TestFindByEqual_OK(t *testing.T) {
	index := New("f1-index", "f1")

	for _, d := range getTestingMixedTypeDocs() {
		err := index.Insert(d)
		assert.NoError(t, err)
	}

	tests := []struct {
		key      interface{}
		expected [][]byte
	}{
		{
			key: 1,
			expected: [][]byte{
				[]byte("1"),
				[]byte("7"),
			},
		},
		{
			key: "a",
			expected: [][]byte{
				[]byte("2"),
			},
		},
		{
			key: 1.1,
			expected: [][]byte{
				[]byte("3"),
			},
		},
		{
			key: "b",
			expected: [][]byte{
				[]byte("6"),
			},
		},
		{
			key: 2,
			expected: [][]byte{
				[]byte("5"),
			},
		},
		{
			key: 2.3,
			expected: [][]byte{
				[]byte("4"),
			},
		},
	}

	for _, test := range tests {
		actual, err := index.Find(test.key)
		assert.NoError(t, err)

		assertEqualResults(test.expected, actual, t)
	}
}

func assertEqualResults(expected [][]byte, actual []document.ObjectID, t *testing.T)  {
	if len(expected) != len(actual) {
		assert.Fail(t, "length of results are not the same")
	}

	for i := range actual {
		assert.Equal(t, expected[i], actual[i].Bytes())
	}
}
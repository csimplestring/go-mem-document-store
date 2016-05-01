package document

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewFromMap(t *testing.T)  {
	m := map[string]interface{}{
		"id": []byte("test-id-1"),
		"a": 1,
		"b": "foo",
		"c": []interface{} {
			1,
			"zoo",
			3.23,
			map[string]interface{} {
				"f": 1,
			},
		},
		"d": map[string]interface{} {
			"e": "bar",
		},
	}

	doc, err := NewFromMap(m)
	assert.NoError(t, err)

	assert.Equal(t, []byte("test-id-1"), doc.ID().Bytes())
//
//	keys := doc.Keys()
//	assert.Contains(t, keys, "id")
//	assert.Contains(t, keys, "a")
//	assert.Contains(t, keys, "b")
//	assert.Contains(t, keys, "c[0]")
//	assert.Contains(t, keys, "c[1]")
//	assert.Contains(t, keys, "c[2]")
//	assert.Contains(t, keys, "c[3].f")
//	assert.Contains(t, keys, "d.e")

	assert.Equal(t, doc.Get("id"), []byte("test-id-1"))
	assert.Equal(t, doc.Get("a"), 1)
	assert.Equal(t, doc.Get("b"), "foo")
	assert.Equal(t, doc.Get("c[0]"), 1)
	assert.Equal(t, doc.Get("c[1]"), "zoo")
	assert.Equal(t, doc.Get("c[2]"), 3.23)
	assert.Equal(t, doc.Get("c[3].f"), 1)
	assert.Equal(t, doc.Get("d.e"), "bar")
}
package exp
import (
	"testing"
	"bytes"
	"github.com/csimplestring/go-mem-store/document"

	"github.com/stretchr/testify/assert"
	"fmt"
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

func TestMatch(t *testing.T)  {
	id := "1"
	positives := []struct{
		doc *doc
		exp Exp

	}{
		{
			doc: newTestingDoc(id, "a", 1),
			exp: NewEq("a", 1),
		},
		{
			doc: newTestingDoc(id, "a", "foo"),
			exp: NewEq("a", "foo"),
		},
		{
			doc: newTestingDoc(id, "a", 1),
			exp: NewGt("a", 0),
		},
		{
			doc: newTestingDoc(id, "a", "foo"),
			exp: NewGt("a", "foa"),
		},
		{
			doc: newTestingDoc(id, "a", 1.1),
			exp: NewGt("a", 1.0),
		},
		{
			doc: newTestingDoc(id, "a", 1),
			exp: NewGte("a", 0),
		},
		{
			doc: newTestingDoc(id, "a", 0),
			exp: NewGte("a", 0),
		},
		{
			doc: newTestingDoc(id, "a", 1),
			exp: NewLt("a", 2),
		},
		{
			doc: newTestingDoc(id, "a", "foo"),
			exp: NewLt("a", "fox"),
		},
		{
			doc: newTestingDoc(id, "a", 1),
			exp: NewLte("a", 1),
		},
		{
			doc: newTestingDoc(id, "a", 1),
			exp: NewLte("a", 2),
		},
		{
			doc: newTestingDoc(id, "a", 1.1),
			exp: NewLte("a", 1.2),
		},
		{
			doc: newTestingDoc(id, "a", "foo"),
			exp: NewLte("a", "foo"),
		},
		{
			doc: &doc{
				id: bytes.NewBufferString(id),
				data: map[string]interface{}{
					"a": 1,
					"b": "foo",
					"c": 1.3,
				},
			},
			exp: NewAndExp([]Exp{
				NewEq("a", 1), NewGt("c", 1.1),
			}),
		},
		{
			doc: &doc{
				id: bytes.NewBufferString(id),
				data: map[string]interface{}{
					"a": 2,
					"b": "foo",
					"c": 1.3,
				},
			},
			exp: NewOrExp([]Exp{
				NewEq("a", 1), NewGt("c", 1.1),
			}),
		},
	}

	for i, test :=range positives {
		expr := test.exp
		doc := test.doc

		assert.True(t, expr.Match(doc), fmt.Sprintf("test case %d wrong", i))
	}
}

func TestNotMatch(t *testing.T)  {
	id:="1"
	negatives := []struct{
		doc *doc
		exp Exp

	}{
		{
			doc: newTestingDoc(id, "a", 1),
			exp: NewEq("a", 2),
		},
		{
			doc: newTestingDoc(id, "a", "2"),
			exp: NewEq("a", 2),
		},
		{
			doc: newTestingDoc(id, "a", 2.0),
			exp: NewEq("a", 2),
		},
		{
			doc: newTestingDoc(id, "a", 1),
			exp: NewGt("a", 1),
		},
		{
			doc: newTestingDoc(id, "a", 1),
			exp: NewGt("a", "1"),
		},
		{
			doc: newTestingDoc(id, "a", 1),
			exp: NewGte("a", 2),
		},
		{
			doc: newTestingDoc(id, "a", 1),
			exp: NewGte("a", "1"),
		},
		{
			doc: newTestingDoc(id, "a", 1),
			exp: NewLt("a", 1),
		},
		{
			doc: newTestingDoc(id, "a", 1),
			exp: NewLte("a", 0),
		},
		{
			doc: &doc{
				id: bytes.NewBufferString(id),
				data: map[string]interface{}{
					"a": 2,
					"b": "foo",
					"c": 1.3,
				},
			},
			exp: NewOrExp([]Exp{
				NewEq("a", 1), NewGt("c", 1.4), NewLt("b", "a"),
			}),
		},
		{
			doc: &doc{
				id: bytes.NewBufferString(id),
				data: map[string]interface{}{
					"a": 2,
					"b": "foo",
					"c": 1.3,
				},
			},
			exp: NewAndExp([]Exp{
				NewEq("a", 2), NewGt("c", 1.2), NewLt("b", "a"),
			}),
		},
	}

	for _, test :=range negatives {
		expr := test.exp
		doc := test.doc

		assert.False(t, expr.Match(doc))
	}
}

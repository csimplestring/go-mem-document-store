package exp
import (
	"github.com/csimplestring/go-mem-store/document"
	"github.com/csimplestring/go-mem-store/types"
)

type Gt struct {
	*BinaryCmpExp
}

func (g *Gt) Match(d document.Document) bool {
	matched := g.checkTypeOrNil(d)
	if !matched {
		return false
	}

	arg := d.Get(g.field)

	switch types.Of(g.value) {
	case types.Int:
		return arg.(int) > g.value.(int)
	case types.Float32:
		return arg.(float32) > g.value.(float32)
	case types.Float64:
		return arg.(float64) > g.value.(float64)
	case types.String:
		return arg.(string) > g.value.(string)
	default:
		return false
	}
}

func NewGt(field string, value interface{}) *Gt {
	return &Gt{
		BinaryCmpExp: &BinaryCmpExp{
			field: field,
			value: value,
		},
	}
}

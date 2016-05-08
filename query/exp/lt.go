package exp
import (
	"github.com/csimplestring/go-mem-store/document"
	"github.com/csimplestring/go-mem-store/types"
)

type Lt struct  {
	*BinaryExp
}

func (l *Lt) Match(d document.Document) bool {
	matched := l.checkTypeOrNil(d)
	if !matched {
		return false
	}

	arg := d.Get(l.field)

	switch types.Of(l.value) {
	case types.Int:
		return arg.(int) < l.value.(int)
	case types.Float32:
		return arg.(float32) < l.value.(float32)
	case types.Float64:
		return arg.(float64) < l.value.(float64)
	case types.String:
		return arg.(string) < l.value.(string)
	default:
		return false
	}
}

func NewLt(field string, value interface{}) *Lt {
	return &Lt{
		BinaryExp: &BinaryExp{
			field: field,
			value: value,
		},
	}
}

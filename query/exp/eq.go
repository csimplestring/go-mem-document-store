package exp
import "github.com/csimplestring/go-mem-store/document"

type Eq struct {
	*BinaryExp
}

func (eq *Eq) Match(d document.Document) bool {
	matched := eq.checkTypeOrNil(d)
	if !matched {
		return false
	}

	return eq.value == d.Get(eq.field)
}

func NewEq(field string, value interface{}) *Eq {
	return &Eq{
		BinaryExp: &BinaryExp{
			field: field,
			value: value,
		},
	}
}

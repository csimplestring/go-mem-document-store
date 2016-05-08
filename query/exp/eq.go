package exp
import "github.com/csimplestring/go-mem-store/document"

type Eq struct {
	*BinaryCmpExp
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
		BinaryCmpExp: &BinaryCmpExp{
			field: field,
			value: value,
		},
	}
}

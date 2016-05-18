package exp

import "github.com/csimplestring/go-mem-store/document"

type Lte struct {
	*Lt
	*Eq
}

func (l *Lte) Match(d document.Document) bool {
	return l.Lt.Match(d) || l.Eq.Match(d)
}

func (l *Lte) Field() string {
	return l.Eq.Field()
}

func NewLte(field string, value interface{}) *Lte {
	return &Lte{
		Lt: NewLt(field, value),
		Eq: NewEq(field, value),
	}
}

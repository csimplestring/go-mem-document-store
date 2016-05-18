package exp

import (
	"github.com/csimplestring/go-mem-store/document"
)

type Gte struct {
	*Gt
	*Eq
}

func (g *Gte) Match(d document.Document) bool {
	return g.Gt.Match(d) || g.Eq.Match(d)
}

func (g *Gte) Field() string {
	return g.Eq.Field()
}

func NewGte(field string, value interface{}) *Gte {
	return &Gte{
		Gt: NewGt(field, value),
		Eq: NewEq(field, value),
	}
}

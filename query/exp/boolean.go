package exp
import "github.com/csimplestring/go-mem-store/document"

type And struct {
	expressions []Exp
}

func (a *And) Match(d document.Document) bool {
	for _, expr := range a.expressions {
		if expr.Match(d) == false {
			return false
		}
	}

	return true
}

type Or struct {
	expressions []Exp
}

func (o *Or) Match(d document.Document) bool {
	for _, expr := range o.expressions {
		if expr.Match(d) == true {
			return true
		}
	}

	return false
}

func NewAndExp(expressions []Exp) *And {
	return &And{
		expressions: expressions,
	}
}

func NewOrExp(expressions []Exp) *Or {
	return &Or{
		expressions: expressions,
	}
}


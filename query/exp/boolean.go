package exp
import "github.com/csimplestring/go-mem-store/document"

type BooleanExp struct  {
	expressions []Exp
}

func (b *BooleanExp) Field() string {
	return ""
}

func (b *BooleanExp) Children() []Exp {
	return b.expressions
}

type And struct {
	*BooleanExp
}

func (a *And) Match(d document.Document) bool {
	for _, expr := range a.expressions {
		if expr.Match(d) == false {
			return false
		}
	}

	return true
}

//func (a *And) Field() string {
//	return ""
//}

type Or struct {
	*BooleanExp
}

func (o *Or) Match(d document.Document) bool {
	for _, expr := range o.expressions {
		if expr.Match(d) == true {
			return true
		}
	}

	return false
}

//func (o *Or) Field() string {
//	return ""
//}

func NewAndExp(expressions []Exp) *And {
	return &And{
		BooleanExp: &BooleanExp{
			expressions: expressions,
		},
	}
}

func NewOrExp(expressions []Exp) *Or {
	return &Or{
		BooleanExp: &BooleanExp{
			expressions: expressions,
		},
	}
}


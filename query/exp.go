package query

type Exp interface {
	Eval(args... interface{}) bool
	Field() string
}

type Eq struct  {
	field string
	value interface{}
}

func (e *Eq) Eval(v interface{}) bool {
	return true
}






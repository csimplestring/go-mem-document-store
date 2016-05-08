package exp

import (
	"github.com/csimplestring/go-mem-store/types"
	"github.com/csimplestring/go-mem-store/document"
)

type Exp interface {
	Match(d document.Document) bool
}

type BinaryExp struct {
	op    string
	field string
	value interface{}
}

func (b *BinaryExp) checkTypeOrNil(d document.Document) bool {
	arg := d.Get(b.field)
	if arg == nil {
		return false
	}

	t := types.Of(b.value)
	if t != types.Of(arg) || t == types.Unsupported {
		return false
	}

	return true
}





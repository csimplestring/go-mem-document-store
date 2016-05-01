package executor
import (
	"github.com/csimplestring/go-mem-store/query"
	"github.com/csimplestring/go-mem-store/document"
)

// for each exp in the query, if one exp is 'ne', use linear scan
// for 'and' logic, take the intersect
// for 'or' logic, take the union
type Executor interface {
	Execute(query query.Query) []document.ObjectID
}

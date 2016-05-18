package plan
import (
	"testing"
	"github.com/csimplestring/go-mem-store/index"
	"github.com/csimplestring/go-mem-store/index/mocks"
	"github.com/csimplestring/go-mem-store/query/exp"
	"github.com/csimplestring/go-mem-store/query"
)

func getIndexManager() *index.IndexManager {
	manager := index.NewManager()

	idx1 := new(mocks.MockedIndex)
	idx1.On("Field").Return("a")

	manager.AddIndex(idx1)

	return manager
}

func TestScanPlanBuilder_Build(t *testing.T)  {
	queryNode := &query.QueryNode{
		Bool: query.BooleanAND,
		Children: []*query.QueryNode{
			&query.QueryNode{
				Exp: exp.NewEq("a", 1),
			},
			&query.QueryNode{
				Exp: exp.NewEq("b", 1),
			},
		},
	}

	builder := &ScanPlanBuilder{
		indexManager: getIndexManager(),
	}

	plan := builder.Build(queryNode)
	t.Log(plan.Explain())
}

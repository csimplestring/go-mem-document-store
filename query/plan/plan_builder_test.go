package plan

import (
	"github.com/csimplestring/go-mem-store/index"
	"github.com/csimplestring/go-mem-store/index/mocks"
	"github.com/csimplestring/go-mem-store/query"
	"github.com/csimplestring/go-mem-store/query/exp"
	"github.com/stretchr/testify/assert"
	"testing"
)

func getIndexManager() *index.IndexManager {
	manager := index.NewManager()

	idx1 := new(mocks.MockedIndex)
	idx1.On("Field").Return("a")

	manager.AddIndex(idx1)

	return manager
}

func TestScanPlanBuilder_Build(t *testing.T) {
	tests := []struct {
		query    *query.QueryNode
		expected string
	}{
		{
			query: &query.QueryNode{
				Exp: exp.NewEq("a", 1),
			},
			expected: "",
		},
		{
			query: &query.QueryNode{
				Bool: query.BooleanAND,
				Children: []*query.QueryNode{
					&query.QueryNode{
						Exp: exp.NewEq("a", 1),
					},
					&query.QueryNode{
						Exp: exp.NewEq("b", 1),
					},
				},
			},
			expected: "",
		},
	}

	builder := &ScanPlanBuilder{
		indexManager: getIndexManager(),
	}

	for _, test := range tests {
		plan := builder.Build(test.query)
		t.Log(plan.Explain())

		assert.Equal(t, test.expected, plan.Explain())
	}
}

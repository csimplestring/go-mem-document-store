package plan

import (
	"github.com/csimplestring/go-mem-store/index"
	"github.com/csimplestring/go-mem-store/query"
)

type ScanPlanBuilder struct {
	indexManager *index.IndexManager
}

func (p *ScanPlanBuilder) Build(node *query.QueryNode) ScanPlan {
	if node.IsLeaf() {
		fieldName := node.Exp.Field()
		idx := p.indexManager.FindIndexByField(fieldName)

		if idx != nil {
			return &IndexScanPlan{
				index: idx,
				exp: node.Exp,
			}
		}

		return &SimpleScanPlan{
			exp: node.Exp,
		}
	}

	if node.IsAnd() {
		var filters []ScanPlan
		for _, subNode := range node.Children {
			subPlan := p.Build(subNode)

			if filter, ok := subPlan.(*IndexScanPlan); ok {
				filters = append(filters, filter)
			}
		}

		return &AndScanPlan{
			exp: node.Exp,
			filters: filters,
		}
	}

	if node.IsOr() {
		useFilter := true
		var filters []ScanPlan

		for _, subNode := range node.Children {
			subPlan := p.Build(subNode)
			if _, ok := subPlan.(*IndexScanPlan); !ok {
				useFilter = false
				break
			}
			filters = append(filters, subPlan)
		}

		orScanPlan := &OrScanPlan{
			exp: node.Exp,
		}

		if useFilter {
			orScanPlan.filters = filters
		}

		return orScanPlan
	}

	return nil
}

package query

//import "github.com/csimplestring/go-mem-store/query/exp"

//type PlanBuilder struct {
//}
//
//func (p *PlanBuilder) Build(node *QueryNode) PlanNode {
//	return nil
//}
//
//func (p *PlanBuilder) BuildExpPlan(exp exp.Exp) *SimpleScanPlan {
//	return &SimpleScanPlan{}
//}
//
//func (p *PlanBuilder) BuildAndQueryPlan(and *QueryNode) *SimpleScanPlan {
//	var candidatePlans []PlanNode
//	for _, subNode := range and.Children {
//		subPlan := p.Build(subNode)
//		candidatePlans = append(candidatePlans, subPlan)
//	}
//
//	bestPlan := p.chooseBestPlan(candidatePlans)
//
//	// conjunction of results
//	finalPlan := &SimpleScanPlan{
//		idx:  bestPlan.idx,
//		exp:  and.Exp,
//		cost: bestPlan.cost,
//	}
//
//
//
//	return finalPlan
//}
//
//func (p *PlanBuilder) chooseBestPlan(plans []*SimpleScanPlan) *SimpleScanPlan {
//	return nil
//}

package query
import "github.com/csimplestring/go-mem-store/index"

type QueryOptimizer struct {
	indices []index.Index
}

func (q *QueryOptimizer) findIndexByFields(fields... string) index.Index {
	for _, idx := range q.indices {
		if idx.Match(fields...) {
			return idx
		}
	}

	return nil
}

func (q *QueryOptimizer) canUseIndex()  {

}

func CanUseIndex(optimizer *QueryOptimizer, node *QueryNode) bool {
	if node.Type == LeafNode {
		return optimizer.findIndexByFields(node.Exp.Field()) != nil
	}

	if node.Type == InnerNode && node.Logic == BooleanAND {
		canUse := false
		for _, subNode := range node.Children {
			canUse = canUse || CanUseIndex(optimizer, subNode)
		}
		return canUse
	}

	if node.Type == InnerNode && node.Logic == BooleanOR {
		canUse := true
		for _, subNode := range node.Children {
			canUse = canUse && CanUseIndex(optimizer, subNode)
		}
		return canUse
	}

	return false;
}

func canUseIndexForAND(indices []index.Index, andQuery QueryNode)  {

}

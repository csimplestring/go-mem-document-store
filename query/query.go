package query

type Query interface {
	GetExpTree() QueryNodeTree
}

type QueryNodeTree interface {
	GetLogicalType() bool
	Children() []QueryNodeTree
	GetAnd() []QueryNodeTree
	GetOr() []QueryNodeTree
	GetExp() QueryNode
}

type NodeType uint8
type BooleanLogic bool

const (
	BooleanAND = BooleanLogic(true)
	BooleanOR  = BooleanLogic(false)

	RootNode = NodeType(0)
	InnerNode = NodeType(1)
	LeafNode = NodeType(2)
)

// Type is BooleanAND by default
type QueryNode struct {
	Type NodeType

	Exp Exp

	Logic    BooleanLogic
	Children []*QueryNode
}

func newQueryNode(logic BooleanLogic, children []*QueryNode) *QueryNode {
	return &QueryNode{
		Logic: logic,
		Children: children,
	}
}

//
//type jsonExpTree struct {
//	And []*jsonExpTree `json:"and,omitempty"`
//	Or  []*jsonExpTree `json:"or,omitempty"`
//	Exp *jsonExp       `json:"exp,omitempty"`
//}
//
//func (j *jsonExpTree) GetAnd() []QueryNodeTree {
//	and := make([]QueryNodeTree, len(j.And))
//	for i, node := range j.And {
//		and[i] = node
//	}
//
//	return and
//}
//
//func (j *jsonExpTree) GetOr() []QueryNodeTree {
//	or := make([]QueryNodeTree, len(j.Or))
//	for i, node := range j.And {
//		or[i] = node
//	}
//
//	return or
//}
//
//func (j *jsonExpTree) GetExp() QueryNode {
//	return j.Exp
//}

////
//type jsonExp struct {
//	Field string      `json:"field"`
//	Op    string      `json:"operator"`
//	Value interface{} `json:"value"`
//}
//
//func (j *jsonExp) GetField() string {
//	return j.Field
//}
//
//func (j *jsonExp) GetOperator() string {
//	return j.Op
//}
//
//func (j *jsonExp) GetValue() interface{} {
//	return j.Value
//}
//
//func (j *jsonExp) Match(doc document.Document) bool {
//	doc.Get(j.Field)
//
//}
//
////
//type JsonQuery struct {
//	expTree *jsonExpTree
//}
//
//func (q *JsonQuery) GetExpTree() QueryNodeTree {
//	return q.expTree
//}
//
//func New(queryString string) (Query, error) {
//	tree, err := parseExpTree(queryString)
//	if err != nil {
//		return nil, err
//	}
//
//	return &JsonQuery{
//		expTree: tree,
//	}, nil
//}
//
//func parseExpTree(q string) (tree *jsonExpTree, err error) {
//	tree = new(jsonExpTree)
//	err = json.Unmarshal([]byte(q), &tree)
//
//	return
//}

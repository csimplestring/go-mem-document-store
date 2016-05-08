package query

import (
	"fmt"
	"github.com/csimplestring/go-mem-store/document"
	"github.com/csimplestring/go-mem-store/index"
	"github.com/csimplestring/go-mem-store/query/exp"
)

type ScanPlan interface {
	Run(docs []*document.Document) []*document.ObjectID
	Cost() float64
	Explain() string
}

type SimpleScanPlan struct {
	exp exp.Exp
}

func (s *SimpleScanPlan) Run(docs []*document.Document) []*document.ObjectID {
	return nil
}

func (s *SimpleScanPlan) Cost() float64 {
	return 0.0
}

func (s *SimpleScanPlan) Explain() string {
	return "simple scan;"
}

type IndexScanPlan struct {
	index index.Index
	exp   exp.Exp
}

func (i *IndexScanPlan) Run(docs []*document.Document) []*document.ObjectID {
	return nil
}

func (i *IndexScanPlan) Cost() float64 {
	return 0.0
}

func (i *IndexScanPlan) Explain() string {
	return fmt.Sprintf("using index on %s;", i.index.Fields())
}

type AndScanPlan struct {
	filters []ScanPlan
	exp     exp.Exp
}

func (a *AndScanPlan) Run(docs []*document.Document) []*document.ObjectID {
	return nil
}

func (a *AndScanPlan) Cost() float64 {
	return 0.0
}

func (a *AndScanPlan) Explain() string {
	if len(a.filters) == 0 {
		return "simple scan;"
	}

	explained := ""
	for _, filter := range a.filters {
		explained = explained + " " + filter.Explain()
	}

	return explained
}

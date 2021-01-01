package articlemanager

import "github.com/aws/aws-sdk-go/service/dynamodb"

type Paginator struct {
	itemsRetreived   int
	pages            int
	max              int
	lastEvaluatedKey map[string]*dynamodb.AttributeValue
}

func NewPaginator(max int, lastEvaluatedKey map[string]*dynamodb.AttributeValue) *Paginator {
	return &Paginator{
		itemsRetreived:   0,
		pages:            0,
		lastEvaluatedKey: nil,
		max:              max,
	}
}

func (p *Paginator) updatePaginator(output *dynamodb.ScanOutput) {
	p.itemsRetreived += len(output.Items)
	p.pages++
	p.lastEvaluatedKey = output.LastEvaluatedKey
}

func (p *Paginator) PagesRemain() bool {
	return p.lastEvaluatedKey == nil || len(p.lastEvaluatedKey) > 0
}

func (p *Paginator) getMax64() int64 {
	return int64(p.max)
}

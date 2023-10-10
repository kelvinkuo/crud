package convert

import (
    "github.com/kelvinkuo/crud/internal/db"
    "github.com/kelvinkuo/crud/internal/protocol"
)

type CommonConverter struct {
    Creators []ItemCreator
    Filters  []ColumnFilter
}

func NewCommonConverter() CommonConverter {
    return CommonConverter{
        Creators: []ItemCreator{},
    }
}

func (p *CommonConverter) AddItemCreator(creator ItemCreator) {
    p.Creators = append(p.Creators, creator)
}

func (p *CommonConverter) CreateItems(table db.Table, service string) ([]protocol.Item, error) {
    var items []protocol.Item
    for _, creator := range p.Creators {
        item, err := creator.ItemCreate(table, service, p.Filters)
        if err != nil {
            return nil, err
        }
        items = append(items, item)
    }
    
    return items, nil
}

func (p *CommonConverter) AddColumnFilter(filter ColumnFilter) {
    p.Filters = append(p.Filters, filter)
}

func FilterOut(col db.Column, filters []ColumnFilter) bool {
    for _, filter := range filters {
        if filter.FilterOut(col) {
            return true
        }
    }
    
    return false
}

package main

import (
	"fmt"
	"toll-calculator/types"
)

const basePrice = 3.15

type Aggregator interface {
	AggregateDistance(types.Distance) error
	CalculateInvoice(int) (*types.Invoice, error)
}

type Storer interface {
	Insert(types.Distance) error
	Read(int) (float64, error)
}

type InvoiceAggregator struct {
	store Storer
}

func NewInvoiceAggregator(store Storer) Aggregator {
	return &InvoiceAggregator{
		store: store,
	}
}

func (i *InvoiceAggregator) AggregateDistance(dist types.Distance) error {
	return i.store.Insert(dist)
}

func (i *InvoiceAggregator) CalculateInvoice(obuId int) (*types.Invoice, error) {
	dist, err := i.store.Read(obuId)
	if err != nil {
		return nil, fmt.Errorf("obu id (%d) not found", obuId)
	}

	inv := &types.Invoice{
		OBUID: obuId,
		TotalDistance: dist,
		TotalCost: basePrice * dist,

	}
	return inv, nil
}
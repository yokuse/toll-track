package client

import (
	"context"
	"toll-calculator/types"
)

type Client interface {
	Aggregate(context.Context, *types.AggregateDistanceRequest) error
	GetInvoice(context.Context, int) (*types.Invoice, error)
}
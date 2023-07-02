package client

import (
	"context"
	"toll-calculator/types"

	"google.golang.org/grpc"
)

type GPRCClient struct {
	Endpoint string
	// when you dont give a var name, you are embedding
	// this will give you direct access to every function that that struct has
	client types.DistanceAggregatorClient
}

func NewGRPCClient(endpoint string) (*GPRCClient, error) {
	conn, err := grpc.Dial(endpoint, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	c := types.NewDistanceAggregatorClient(conn)

	return &GPRCClient{
		Endpoint:                 endpoint,
		client: c,
	}, nil
}

func (c *GPRCClient) Aggregate(ctx context.Context, r *types.AggregateDistanceRequest) error {
	_, err := c.client.AggregateDistance(ctx, r)
	return err
}

func (c *GPRCClient) GetInvoice(ctx context.Context, id int) (*types.Invoice, error) {
	return nil, nil
}

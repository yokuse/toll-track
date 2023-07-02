package main

import (
	"context"
	"toll-calculator/types"
)

type GRPCAggregatorServer struct {
	types.UnimplementedDistanceAggregatorServer
	svc Aggregator
	
}

func NewGRPCAggregatorServer(svc Aggregator) *GRPCAggregatorServer {
	return &GRPCAggregatorServer{
		svc: svc,
	}
}

func (s *GRPCAggregatorServer) AggregateDistance(ctx context.Context, req *types.AggregateDistanceRequest) (*types.None, error) {
	distance := types.Distance{
		OBUID: int(req.ObuID),
		Value: req.Value,
		Unix: req.Unix,
	}

	// sine the svc here is of type aggregator interface, there is no need to add logging
	// all logging is already done in the logging middle ware service which is also implements the aggregator ingerface
	// which is the function that we are calling here
	return &types.None{}, s.svc.AggregateDistance(distance)
}
package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"toll-calculator/types"
)

type HTTPClient struct {
	Endpoint string
}

func NewClient(endpoint string) *HTTPClient {
	return &HTTPClient{
		Endpoint: endpoint,
	}
}

// for gateway to call then aggregator will call all other microservices
func (c *HTTPClient) GetInvoice(ctx context.Context, id int) (*types.Invoice, error) {
	invReq := types.GetInvoiceRequest{
		ObuID: int32(id),
	}

	b, err := json.Marshal(&invReq)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.Endpoint + "/invoice", bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("service responded with a non 200 status code %d", resp.StatusCode)
	}

	var inv types.Invoice
	if err := json.NewDecoder(resp.Body).Decode(&inv); err != nil {
		return nil, err
	}
	// close the body for resources
	defer resp.Body.Close()
	return &inv, nil
}

// give this clients the name of your endpoint
func (c *HTTPClient) Aggregate(ctx context.Context, r *types.AggregateDistanceRequest) error {
	b, err := json.Marshal(r)
	if err != nil {
		return err
	}

	// body that we are sending needs to be an io reader so we put in newreader
	req, err := http.NewRequest("POST", c.Endpoint + "/aggregate", bytes.NewReader(b))
	if err != nil {
		return err
	}
	
	// make the request
	httpc := http.DefaultClient
	res, err := httpc.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("service responded with a non 200 status code %d", res.StatusCode)
	}

	return nil
}
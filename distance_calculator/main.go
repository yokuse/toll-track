package main

import (
	"log"
	"toll-calculator/aggregator/client"
)

const aggregatorEndpoint = "http://127.0.0.1:30001"

func main() {

	// we can do this because NewLogMiddleware returns the same interface, calculatorservicer, 
	// So essentially we are extending the functionality of the implemented functions
	// NewCalculatorService gives the main business math logic
	// NewLogMiddleware adds logging and calls the main business logic function
	svc := NewCalculatorService()	// initialise service
	svc = NewLogMiddleware(svc)		// add middle ware functions to it

	// httpClient := client.NewClient(aggregatorEndpoint)
	grpcClient, err := client.NewGRPCClient(aggregatorEndpoint)
	if err != nil {
		log.Fatal(err)
	}

	kafkaConsumer, err := NewKafkaConsumer("obu-data", svc, grpcClient)
	if err != nil {
		log.Fatal(err)
	}

	kafkaConsumer.Start()
}
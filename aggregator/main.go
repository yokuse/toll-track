package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"
	"toll-calculator/types"

	"google.golang.org/grpc"
)

func main() {
	// create a flag that could be passed to the commant line. this returns a pointer to the string
	httpListenAddr := flag.String("httpAddr", ":30001", "the listen address of the HTTP server")
	grpcListenAddr := flag.String("grpcAddr", ":30002", "the listen address of the GRPC server")

	// after all flags are defined, call to pass the defined flags to the command line so that they can be used
	flag.Parse()

	store := NewMemoryStore()
	svc := NewInvoiceAggregator(store)
	svc = NewLogMiddleware(svc)

	// you can also make a switch case something like a factory method
	// make grpc transport in a new goroutine
	go func() {
		// any error in starting grpc server will be logged
		log.Fatal(makeGRPCTransport(*grpcListenAddr, svc))
	}()

	// then here you can make all the different channels for endpoints
	// i.e. GRPC, HTTP, Kafka etc
	log.Fatal(makeHTTPTransport(*httpListenAddr, svc))
}

func makeHTTPTransport(listenAddr string, svc Aggregator) error {
	fmt.Println("HTTP transport running on port", listenAddr)
	// expose API endpoints
	// httpe end point with handlerfunc function defined to handle any calls made to this endpoint
	http.HandleFunc("/aggregate", aggregateData(svc))
	http.HandleFunc("/invoice", handleGetInvoice(svc)) // endpoint for api handler for frontend to display invoice for obu id

	// Serve on which port, use the previously defined flag, the handler here is usually nil as mentioned in the docs
	return http.ListenAndServe(listenAddr, nil)
}

func makeGRPCTransport(listenAddr string, svc Aggregator) error {
	fmt.Println("GRPC transport running on port", listenAddr)
	// make tcp listener
	ln, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return err
	}
	defer ln.Close()

	// make a new grpc native server with options
	server := grpc.NewServer([]grpc.ServerOption{}...)
	// register our grpc server implementation to the grpc package
	types.RegisterDistanceAggregatorServer(server, NewGRPCAggregatorServer(svc))
	return server.Serve(ln)
}

func handleGetInvoice(svc Aggregator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get query string params
		value := r.URL.Query().Get("obu")
		if value == "" {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "missing obu id"})
			return
		}
		
		obuId, err := strconv.Atoi(value) 
		if err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid obu id"})
			return
		}

		inv, err := svc.CalculateInvoice(obuId)
		if err != nil {
			writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}

		writeJSON(w, http.StatusOK, map[string]*types.Invoice{"distance": inv})
	}
}

// this is a decorator pattern, basically we dont want http in the aggregate data class
// so we decorate the aggregate function with the http function
func aggregateData(svc Aggregator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var distance types.Distance
		if err := json.NewDecoder(r.Body).Decode(&distance); err != nil {
			// got error decoding the data coming through
			// so write bad request, can also return body. normal http request stuff
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}

		if err := svc.AggregateDistance(distance); err != nil {
			writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}

	}
}

func writeJSON(rw http.ResponseWriter, status int, v any) error {
	rw.WriteHeader(status)
	rw.Header().Add("Content-Type", "application/json")

	// v is any type, so we need to encode it to json
	// can be any interface that you are putting in
	return json.NewEncoder(rw).Encode(v)
}
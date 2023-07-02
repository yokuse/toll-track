package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"toll-calculator/aggregator/client"

	"github.com/sirupsen/logrus"
)

// decorator patterns, adding functionality to existing functions
// we want to return errors from handler function, http.HandlerFunc does not allow us to do that
// so we extend its functionality by creating a new type apiFunc which returns an error
type apiFunc func(w http.ResponseWriter, r *http.Request) error

func main() {
	listenAddr := flag.String("listen-addr", ":30003", "server listen address")
	flag.Parse()

	client := client.NewClient("http://localhost:30001")
	invoiceHandler := NewInvoiceHandler(client)

	// part 3 of decorator, wrap the handler function, which returns the type we want.
	http.HandleFunc("/invoice", makAPIFunc(invoiceHandler.handleGetInvoice))
	logrus.Infof("Starting server on port, %s", *listenAddr)
	log.Fatal(http.ListenAndServe(*listenAddr, nil))
}

type InvoiceHandler struct {
	client client.Client
}

func NewInvoiceHandler(client client.Client) *InvoiceHandler {
	return &InvoiceHandler{
		client: client,
	}
}

func (h *InvoiceHandler) handleGetInvoice(w http.ResponseWriter, r *http.Request) error {
	// need agg client
	inv , err := h.client.GetInvoice(context.Background(), 234987)
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, inv)
}

func writeJSON(w http.ResponseWriter, code int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(v)
}

// this is the second part of the decorator pattern
// this helper function to wrap around handlerfunc
func makAPIFunc(fn apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// so here we call our own function and return error if there is an error
		if err := fn(w, r); err != nil {
			writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
	}
}
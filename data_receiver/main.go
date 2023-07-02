package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"toll-calculator/types"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

func main() {

	recv, err := NewDataReceiver()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/ws", recv.handleWs)
	http.ListenAndServe(":30000", nil)
}

type DataReceiver struct {
	msgch chan types.OBUData
	conn  *websocket.Conn
	prod *LogMiddleware
}

func NewDataReceiver() (*DataReceiver, error) {
	kafkaTopic := "obu-data"		// hardcoded for now
	// Initialise new producer that connects to localhost 9092
	// 9092 is the port exposed for the broker in docker compose
	p, err := NewKafkaProducer(kafkaTopic)
	if err != nil {
		log.Fatal("Failed to create Kafka producer")
		return nil, err
	}

	return &DataReceiver{
		msgch: make(chan types.OBUData, 128),
		prod:  NewLogMiddleware(p),
	}, nil
}

func (dr *DataReceiver) pushData(data types.OBUData) error {
	return dr.prod.PushData(data)
}

func (dr *DataReceiver) handleWs(w http.ResponseWriter, r *http.Request) {
	// websocket server - read from websocket client (the vehicles/ obu data generator)
	c, err := websocket.Accept(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close(websocket.StatusInternalError, "the sky is falling")

	// Create a separate context for WebSocket operations
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	dr.conn = c

	go dr.wsReceiveLoop(ctx)

	// Handle the HTTP request context separately
	select {
	case <-r.Context().Done():
		// HTTP request canceled or timed out, stop the WebSocket communication
		cancel()
	case <-ctx.Done():
		// WebSocket communication finished or canceled, do nothing
	}
}

func (dr *DataReceiver) wsReceiveLoop(ctx context.Context) {
	fmt.Println("OBU client connected!")
	for {
		var data types.OBUData
		if err := wsjson.Read(ctx, dr.conn, &data); err != nil {
			// handle error
			log.Println(err)
			continue
		}
		if err := dr.pushData(data); err != nil {
			fmt.Println("Failed to push message to kafka, err: ", err)
		}
	}
}

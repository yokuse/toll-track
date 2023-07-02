package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
	"toll-calculator/types"
)

var sendInterval = time.Second * 5
const wsEndPoint = "ws://localhost:30000/ws"

func genLatLong() (float64,float64) {
	return genCoord(), genCoord()
}

func genCoord() float64 {
	n := float64(rand.Intn(100) + 1) // so that we never have 0
	f := rand.Float64()               // so that we have a decimal, random decimal
	return n + f
}

// mock vehicles sneding data
func main() {
	// openwebsocket 
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	c, _, err := websocket.Dial(ctx, wsEndPoint, nil)
	if err != nil {
		//  Failed to dial: websocket: bad handshake
		log.Fatal(err)
	}

	obuIds := generateOBUIDs(20)	// mock obu data
	for {
		// generate 20 OBU data
		for i := 0; i < len(obuIds); i++ {
			lat, long := genLatLong()
			data := types.OBUData{
				OBUID: obuIds[i],
				Lat:   lat,
				Long:  long,
			}
			
			// write to websocket connection
			if err := wsjson.Write(ctx, c, data); err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%+v\n", data)
		}
		time.Sleep(sendInterval)
	}
}

func generateOBUIDs(n int) []int {
	ids := make([]int, n)
	for i := 0; i < n; i++ {
		ids[i] = rand.Intn(math.MaxInt)
	}
	return ids
}

// init runs before main, ensure true random
func init() {
	rand.Seed(time.Now().UnixNano())
}

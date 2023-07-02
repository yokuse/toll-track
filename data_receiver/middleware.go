package main

import (
	"time"
	"toll-calculator/types"

	"github.com/sirupsen/logrus"
)

type LogMiddleware struct {
	next DataProducer
}

func NewLogMiddleware(next DataProducer) *LogMiddleware {
	return &LogMiddleware{
		next: next, 
	}
}

func (l *LogMiddleware) PushData(data types.OBUData) error {

	// log, defer means return after this function ends so we can measure along with this function
	// logging done here instead of new goroutine which the kafka docs has
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"obuID": data.OBUID,
			"lat": data.Lat,
			"long": data.Long,
			"timestamp": time.Since(start),
		}).Info("Producing to kafka")
	}(time.Now())
	
	return l.next.PushData(data)
}

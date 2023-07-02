package main

import (
	"time"
	"toll-calculator/types"

	"github.com/sirupsen/logrus"
)

type LogMiddleware struct {
	next Aggregator
}

func NewLogMiddleware(next Aggregator) Aggregator {
	return &LogMiddleware{next: next}
}

func (m *LogMiddleware) AggregateDistance(d types.Distance) (err error) {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"took": time.Since(start),
			"err": err,
		}).Info("AggregateDistance")
	}(time.Now())

	err = m.next.AggregateDistance(d)

	return err
}

func (m *LogMiddleware) CalculateInvoice(obuId int) (inv *types.Invoice, err error) {
	defer func(start time.Time) {
		var (
			distance float64
			cost float64
		)

		if inv != nil {
			distance = inv.TotalDistance
			cost = inv.TotalCost
		}

		logrus.WithFields(logrus.Fields{
			"took": time.Since(start),
			"err": err,
			"distance": distance,
			"cost": cost,
		}).Info("CalculateInvoice")
	}(time.Now())

	inv, err = m.next.CalculateInvoice(obuId)

	return inv, err
}
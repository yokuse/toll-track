package main

import (
	"math"
	"toll-calculator/types"
)

type CalculatorServicer interface {
	CalculateDistance(types.OBUData) (float64, error)
}

type CalculatorService struct {
	prevPoints []float64
}

func NewCalculatorService() CalculatorServicer {
	return &CalculatorService{
		prevPoints: []float64{0.0, 0.0},
	}
}

func (s *CalculatorService) CalculateDistance(data types.OBUData) (float64, error) {
	// just mock calculate distance, by getting the diff of the 2 distance
	distance := 0.0
	distance = calculateDistance(s.prevPoints[0], s.prevPoints[1], data.Lat, data.Long)
	s.prevPoints = []float64{data.Lat, data.Long}
	
	return distance, nil
}

func calculateDistance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2 - x1, 2) + math.Pow(y2 - y1, 2))
}
package main

import (
	"github.com/stretchr/testify/mock"
)

// TestPoint is a mocked object that should implement all the interfaces of TestPoint
type TestPoint struct {
	mock.Mock
}

func (tp *TestPoint) GetX() int {
	args := tp.Called()
	return args.Int(0)
}

func (tp *TestPoint) GetY() int {
	args := tp.Called()
	return args.Int(0)
}

func (tp *TestPoint) Distance(tp2 Plottable) float64 {
	args := tp.Called(tp2)
	return args.Get(0).(float64)
}

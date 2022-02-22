package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAreEquiDistantPoints(t *testing.T) {
	p1 := new(TestPoint)
	p2 := new(TestPoint)
	p3 := new(TestPoint)
	p4 := new(TestPoint)

	require.NotNil(t, p1)
	require.NotNil(t, p2)
	require.NotNil(t, p3)
	require.NotNil(t, p4)

	p1.On("Distance", p2).Return(1.0)
	p3.On("Distance", p4).Return(1.0)

	ret := AreEquiDistantPoints(p1, p2, p3, p4)

	assert.True(t, ret)
}

func TestAreNotEquiDistantPoints(t *testing.T) {
	p1 := new(TestPoint)
	p2 := new(TestPoint)
	p3 := new(TestPoint)
	p4 := new(TestPoint)

	require.NotNil(t, p1)
	require.NotNil(t, p2)
	require.NotNil(t, p3)
	require.NotNil(t, p4)

	p1.On("Distance", p2).Return(1.0)
	p3.On("Distance", p4).Return(2.0)

	ret := AreEquiDistantPoints(p1, p2, p3, p4)

	assert.False(t, ret)
}

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcess(t *testing.T) {
	agg := &Aggregator{
		Market:        1,
		TotalVolume:   17.0,
		TotalPrice:    33.0,
		WeightedPrice: 190.0,
		TotalCount:    3,
		BuyCount:      1,
	}
	tx := Tx{
		Id:     100,
		Market: 1,
		Price:  9.0,
		Volume: 20.0,
		IsBuy:  true,
	}
	agg.Process(tx)
	assert.Equal(t, &Aggregator{
		Market:        1,
		TotalVolume:   37.0,
		TotalPrice:    42.0,
		WeightedPrice: 370.0,
		TotalCount:    4,
		BuyCount:      2,
	}, agg)
}

func TestAggregate(t *testing.T) {
	agg := &Aggregator{
		Market:        1,
		TotalVolume:   37.0,
		TotalPrice:    42.0,
		WeightedPrice: 370.0,
		TotalCount:    4,
		BuyCount:      2,
	}
	res := agg.Aggregate()
	assert.Equal(t, int64(1), res.Market)
	assert.Equal(t, 37.0, res.TotalVolume)
	assert.Equal(t, 10.5, res.MeanPrice)
	assert.Equal(t, 9.25, res.MeanVolume)
	assert.Equal(t, 10.0, res.VolumeWeightedAveragePrice)
	assert.Equal(t, 50.0, res.PercentageBuy)
}

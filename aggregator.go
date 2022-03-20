package main

type Aggregator struct {
	Market        int64
	TotalVolume   float64
	TotalPrice    float64
	WeightedPrice float64
	TotalCount    int64
	BuyCount      int64
}

type AggregatedResult struct {
	Market                     int64   `json:"market"`
	TotalVolume                float64 `json:"total_volume"`
	MeanPrice                  float64 `json:"mean_price"`
	MeanVolume                 float64 `json:"mean_volume"`
	VolumeWeightedAveragePrice float64 `json:"volume_weighted_average_price"`
	PercentageBuy              float64 `json:"percentage_buy"`
}

func (a *Aggregator) Process(tx Tx) {
	a.TotalVolume += tx.Volume
	a.TotalPrice += tx.Price
	a.WeightedPrice += tx.Volume * tx.Price
	a.TotalCount++
	if tx.IsBuy {
		a.BuyCount++
	}
}

func (a *Aggregator) Aggregate() AggregatedResult {
	return AggregatedResult{
		Market:                     a.Market,
		TotalVolume:                a.TotalVolume,
		MeanPrice:                  a.TotalPrice / float64(a.TotalCount),
		MeanVolume:                 a.TotalVolume / float64(a.TotalCount),
		VolumeWeightedAveragePrice: a.WeightedPrice / a.TotalVolume,
		PercentageBuy:              float64(a.BuyCount) / float64(a.TotalCount) * 100,
	}
}

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Tx struct {
	Id     int64   `json:"id"`
	Market int64   `json:"market"`
	Price  float64 `json:"price"`
	Volume float64 `json:"volume"`
	IsBuy  bool    `json:"is_buy"`
}

func main() {
	aggs := map[int64]*Aggregator{}

	scanner := bufio.NewScanner(os.Stdin)
	// Skip the first line 'BEGIN\n'
	_ = scanner.Scan()
	for scanner.Scan() {
		line := scanner.Text()
		var tx Tx
		err := json.Unmarshal([]byte(line), &tx)
		if err != nil {
			if line == "END" {
				break
			} else {
				log.Fatal(err)
			}
		}
		var marketAgg *Aggregator
		if agg, ok := aggs[tx.Market]; !ok {
			aggs[tx.Market] = &Aggregator{Market: tx.Market}
			marketAgg = aggs[tx.Market]
		} else {
			marketAgg = agg
		}
		marketAgg.Process(tx)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for _, agg := range aggs {
		result := agg.Aggregate()
		json, err := json.Marshal(result)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(json))
	}
}

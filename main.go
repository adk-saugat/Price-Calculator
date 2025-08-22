package main

import "fmt"

func main(){
	prices := []float64{10, 20, 30}
	taxRates := []float64{0, 0.1, 0.15}

	result := make(map[float64][]float64)

	for _, taxValue := range taxRates{
		taxIncludedPrices := make([]float64, len(prices))
		for priceIndex, priceValue := range prices{
			taxIncludedPrices[priceIndex] = priceValue * (1 + taxValue)
		}
		result[taxValue] = taxIncludedPrices
	}
	fmt.Println(result)
}
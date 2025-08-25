package main

import (
	"fmt"

	"github.com/price-calculator/filemanager"
	"github.com/price-calculator/prices"
)

func main(){
	taxRates := []float64{0, 0.1, 0.15}

	for _, taxValue := range taxRates{
		fm := filemanager.New("prices.txt",fmt.Sprintf("result_%.0f.json", taxValue * 100))
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxValue)
		priceJob.Process()
	}
}
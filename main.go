package main

import (
	"fmt"

	"github.com/price-calculator/filemanager"
	"github.com/price-calculator/prices"
)

func main(){
	taxRates := []float64{0, 0.1, 0.15}

	for _, taxValue := range taxRates{
		fm := filemanager.New("pricess.txt",fmt.Sprintf("result_%.0f.json", taxValue * 100))
		// cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxValue)
		err := priceJob.Process()

		if err != nil{
			fmt.Println(err)
		}
	}
}
package prices

import (
	"fmt"

	"github.com/price-calculator/conversion"
	"github.com/price-calculator/filemanager"
)

type TaxIncludedPriceJob struct{
	TaxRate 					float64
	InputPrices 				[]float64
	TaxIncludedPrices 			map[string]string
}

func (job *TaxIncludedPriceJob) LoadData(){
	lines, err := filemanager.ReadLines("prices.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	prices, err := conversion.StringsToFloats(lines)

	if err != nil {
		fmt.Println(err)
		return
	}

	job.InputPrices = prices
}

func (job *TaxIncludedPriceJob) Process(){
	job.LoadData()

	result := make(map[string]string)
	for _, priceValue := range job.InputPrices{
		taxIncludedPrice := priceValue * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f",priceValue)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrices = result

	filemanager.WriteJSON(fmt.Sprintf("result_%.0f.json", job.TaxRate * 100), job)
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob{
	return &TaxIncludedPriceJob{
		TaxRate: taxRate,
		InputPrices: []float64{},
	}
}
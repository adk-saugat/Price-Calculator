package prices

import "fmt"

type TaxIncludedPriceJob struct{
	TaxRate 					float64
	InputPrices 				[]float64
	TaxIncludedPrices map[string]float64
}

func (job TaxIncludedPriceJob) Process(){
	result := make(map[string]float64)
	for _, priceValue := range job.InputPrices{
		result[fmt.Sprintf("%.2f",priceValue)] = priceValue * (1 + job.TaxRate)
	}

	fmt.Println(result)
			
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob{
	return &TaxIncludedPriceJob{
		TaxRate: taxRate,
		InputPrices: []float64{10, 20, 30},
	}
}
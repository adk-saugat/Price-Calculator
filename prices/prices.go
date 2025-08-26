package prices

import (
	"fmt"

	"github.com/price-calculator/conversion"
	"github.com/price-calculator/iomanager"
)

type TaxIncludedPriceJob struct{
	IOManager			iomanager.IOManager 	`json:"-"`
	TaxRate 			float64					`json:"taxRate"`
	InputPrices 		[]float64				`json:"inputPrices"`
	TaxIncludedPrices 	map[string]string		`json:"taxIncludedPrices"`
}

func (job *TaxIncludedPriceJob) LoadData() error{
	lines, err := job.IOManager.ReadLines()

	if err != nil {
		fmt.Println(err)
		return err
	}

	prices, err := conversion.StringsToFloats(lines)

	if err != nil {
		fmt.Println(err)
		return err
	}

	job.InputPrices = prices
	return nil
}

func (job *TaxIncludedPriceJob) Process() error{
	err := job.LoadData()

	if err != nil {
		return err
	}

	result := make(map[string]string)
	for _, priceValue := range job.InputPrices{
		taxIncludedPrice := priceValue * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f",priceValue)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrices = result

	return job.IOManager.WriteResult(job)
}

func NewTaxIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob{
	return &TaxIncludedPriceJob{
		IOManager: iom,
		TaxRate: taxRate,
		InputPrices: []float64{},
	}
}
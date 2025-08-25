package prices

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type TaxIncludedPriceJob struct{
	TaxRate 					float64
	InputPrices 				[]float64
	TaxIncludedPrices map[string]float64
}

func (job *TaxIncludedPriceJob) LoadData(){
	file, err := os.Open("prices.txt")

	if err != nil{
		fmt.Println("An error occured:\n", err)
		return
	}
	
	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil{
		fmt.Println("An error occured:\n", err)
		file.Close()
		return
	}

	prices := make([]float64, len(lines))
	for index, lineValue := range lines{
		floatPrice, err := strconv.ParseFloat(lineValue, 64)

		if err != nil{
			fmt.Println("An error occured:\n", err)
			file.Close()
			return
		}

		prices[index] = floatPrice
	}

	file.Close()
	job.InputPrices = prices
}

func (job *TaxIncludedPriceJob) Process(){
	job.LoadData()

	result := make(map[string]string)
	for _, priceValue := range job.InputPrices{
		taxIncludedPrice := priceValue * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f",priceValue)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	fmt.Println(result)
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob{
	return &TaxIncludedPriceJob{
		TaxRate: taxRate,
		InputPrices: []float64{},
	}
}
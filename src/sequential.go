package main

import (
	"encoding/csv"
	"os"
	"strconv"
)

func SequentialRunWithAccuracy() float64 {
	file, _ := os.Open("../data/car_sales_clean_updated.csv")
	defer file.Close()

	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()

	var correctPredictions int

	for i, record := range records {
		if i == 0 {
			continue
		}
		year, _ := strconv.Atoi(record[2])
		price, _ := strconv.ParseFloat(record[3], 64)
		mileage, _ := strconv.Atoi(record[4])
		model, _ := strconv.Atoi(record[1])
		country, _ := strconv.Atoi(record[5])

		car := Car{
			Brand:     0,
			Model:     model,
			Year:      year,
			Price:     price,
			Mileage:   mileage,
			Country:   country,
			Condition: record[6],
		}

		predicted := PredictRisk(car)
		if predicted == record[7] {
			correctPredictions++
		}
	}

	total := len(records) - 1
	return (float64(correctPredictions) / float64(total)) * 100
}

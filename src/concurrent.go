package main

import (
	"encoding/csv"
	"os"
	"strconv"
	"sync"
)

func ConcurrentRunWithAccuracy() float64 {
	file, _ := os.Open("../data/car_sales_clean_updated.csv")
	defer file.Close()

	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()

	var correctPredictions int
	var wg sync.WaitGroup
	var mutex sync.Mutex

	batchSize := 10000

	for i := 0; i < len(records); i += batchSize {
		end := i + batchSize
		if end > len(records) {
			end = len(records)
		}

		wg.Add(1)
		go func(batch [][]string) {
			defer wg.Done()
			localCorrect := 0

			for j, record := range batch {
				if i+j == 0 {
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
					localCorrect++
				}
			}

			mutex.Lock()
			correctPredictions += localCorrect
			mutex.Unlock()
		}(records[i:end])
	}

	wg.Wait()

	total := len(records) - 1
	return (float64(correctPredictions) / float64(total)) * 100
}

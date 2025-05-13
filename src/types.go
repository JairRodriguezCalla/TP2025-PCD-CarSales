package main

type Car struct {
	Brand     int
	Model     int
	Year      int
	Price     float64
	Mileage   int
	Country   int
	Condition string
}

func PredictRisk(car Car) string {
	if car.Price < 15000 && car.Mileage < 50000 && car.Year >= 2018 {
		return "Bajo"
	} else if car.Price >= 15000 && car.Price <= 30000 && car.Mileage <= 100000 && car.Year >= 2015 {
		return "Moderado"
	} else {
		return "Alto"
	}
}

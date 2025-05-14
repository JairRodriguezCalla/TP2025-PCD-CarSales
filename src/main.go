package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n=== Sistema de Clasificación de Riesgo de Inversión ===")
		fmt.Println("[1] Ingresar datos manualmente y calcular riesgo")
		fmt.Println("[2] Cargar archivo CSV y generar resultados")
		fmt.Println("[3] Ejecutar Benchmark de Rendimiento")
		fmt.Println("[4] Salir")
		fmt.Print("Seleccione una opción: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			manualInput(reader)
		case "2":
			processCSV(reader)
		case "3":
			BenchmarkTest()
		case "4":
			fmt.Println("Saliendo...")
			return
		default:
			fmt.Println("Opción no válida. Intente nuevamente.")
		}
	}
}

func manualInput(reader *bufio.Reader) {
	fmt.Print("\nIngrese el año del vehículo: ")
	yearStr, _ := reader.ReadString('\n')
	year, _ := strconv.Atoi(strings.TrimSpace(yearStr))

	fmt.Print("Ingrese el precio del vehículo: ")
	priceStr, _ := reader.ReadString('\n')
	price, _ := strconv.ParseFloat(strings.TrimSpace(priceStr), 64)

	fmt.Print("Ingrese el kilometraje del vehículo: ")
	mileageStr, _ := reader.ReadString('\n')
	mileage, _ := strconv.Atoi(strings.TrimSpace(mileageStr))

	fmt.Print("Ingrese el código del modelo (número): ")
	modelStr, _ := reader.ReadString('\n')
	model, _ := strconv.Atoi(strings.TrimSpace(modelStr))

	fmt.Print("Ingrese el código del país (número): ")
	countryStr, _ := reader.ReadString('\n')
	country, _ := strconv.Atoi(strings.TrimSpace(countryStr))

	car := Car{
		Brand:     0,
		Model:     model,
		Year:      year,
		Price:     price,
		Mileage:   mileage,
		Country:   country,
		Condition: "N/A",
	}

	risk := PredictRisk(car)
	fmt.Printf(">> Riesgo de Inversión: %s\n", risk)
}

func processCSV(reader *bufio.Reader) {
	fmt.Print("\nIngrese la ruta del archivo CSV: ")
	path, _ := reader.ReadString('\n')
	path = strings.TrimSpace(path)

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println("Error al leer el archivo:", err)
		return
	}

	// Agregar la nueva columna de Riesgo
	header := append(records[0], "Riesgo_Inversion")
	var newRecords [][]string
	newRecords = append(newRecords, header)

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

		risk := PredictRisk(car)
		newRecord := append(record, risk)
		newRecords = append(newRecords, newRecord)
	}

	outputPath := "../results/resultado_riesgo.csv"
	outputFile, _ := os.Create(outputPath)
	defer outputFile.Close()

	writer := csv.NewWriter(outputFile)
	writer.WriteAll(newRecords)
	writer.Flush()

	fmt.Printf(">> Archivo procesado exitosamente. Resultado: %s\n", outputPath)
}

package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"time"
)

func BenchmarkTest() {
	const iterations = 1000

	var seqTimes, conTimes, seqAccs, conAccs []float64

	for i := 0; i < iterations; i++ {
		start := time.Now()
		acc := SequentialRunWithAccuracy()
		seqTimes = append(seqTimes, time.Since(start).Seconds()*1000)
		seqAccs = append(seqAccs, acc)
	}

	for i := 0; i < iterations; i++ {
		start := time.Now()
		acc := ConcurrentRunWithAccuracy()
		conTimes = append(conTimes, time.Since(start).Seconds()*1000)
		conAccs = append(conAccs, acc)
	}

	fmt.Println("\n=== Media Recortada de Tiempos (ms) ===")
	fmt.Printf("Secuencial: %.2f ms\n", TrimmedMean(seqTimes))
	fmt.Printf("Concurrente: %.2f ms\n", TrimmedMean(conTimes))

	fmt.Println("\n=== Precisión Promedio del Modelo ===")
	fmt.Printf("Secuencial: %.2f%%\n", Mean(seqAccs))
	fmt.Printf("Concurrente: %.2f%%\n", Mean(conAccs))

	SaveBenchmarkResults(seqTimes, conTimes, seqAccs, conAccs)
}

func TrimmedMean(times []float64) float64 {
	sort.Float64s(times)
	trimmed := times[50 : len(times)-50]
	sum := 0.0
	for _, v := range trimmed {
		sum += v
	}
	return sum / float64(len(trimmed))
}

func Mean(values []float64) float64 {
	sum := 0.0
	for _, v := range values {
		sum += v
	}
	return sum / float64(len(values))
}

func SaveBenchmarkResults(seqTimes, conTimes, seqAccs, conAccs []float64) {
	file, _ := os.Create("../results/benchmark_results.csv")
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"Iteration", "Sequential_Time(ms)", "Concurrent_Time(ms)", "Sequential_Accuracy(%)", "Concurrent_Accuracy(%)"})

	for i := 0; i < len(seqTimes); i++ {
		row := []string{
			fmt.Sprintf("%d", i+1),
			fmt.Sprintf("%.2f", seqTimes[i]),
			fmt.Sprintf("%.2f", conTimes[i]),
			fmt.Sprintf("%.2f", seqAccs[i]),
			fmt.Sprintf("%.2f", conAccs[i]),
		}
		writer.Write(row)
	}
}

package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

type Record struct {
	Date        string
	BuyingRate  float64
	MidRate     float64
	SellingRate float64
}

func main() {
	// Read data from CSV file
	records, err := readCSV("data.csv")
	if err != nil {
		log.Fatal(err)
	}

	// Compute statistics
	buyingRates := extractColumn(records, "BuyingRate")
	midRates := extractColumn(records, "MidRate")
	sellingRates := extractColumn(records, "SellingRate")

	buyingStats := computeStatistics(buyingRates)
	midStats := computeStatistics(midRates)
	sellingStats := computeStatistics(sellingRates)

	// Print statistics
	fmt.Println("Statistics for Buying Rate:")
	printStatistics(buyingStats)

	fmt.Println("\nStatistics for Mid Rate:")
	printStatistics(midStats)

	fmt.Println("\nStatistics for Selling Rate:")
	printStatistics(sellingStats)
}

func readCSV(filename string) ([]Record, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var data []Record

	for _, row := range records[1:] {
		buyingRate, err := strconv.ParseFloat(row[2], 64)
		if err != nil {
			return nil, err
		}

		midRate, err := strconv.ParseFloat(row[3], 64)
		if err != nil {
			return nil, err
		}

		sellingRate, err := strconv.ParseFloat(row[4], 64)
		if err != nil {
			return nil, err
		}

		record := Record{
			Date:        row[0],
			BuyingRate:  buyingRate,
			MidRate:     midRate,
			SellingRate: sellingRate,
		}

		data = append(data, record)
	}

	return data, nil
}

func extractColumn(records []Record, columnName string) []float64 {
	var column []float64

	switch columnName {
	case "BuyingRate":
		for _, record := range records {
			column = append(column, record.BuyingRate)
		}
	case "MidRate":
		for _, record := range records {
			column = append(column, record.MidRate)
		}
	case "SellingRate":
		for _, record := range records {
			column = append(column, record.SellingRate)
		}
	}

	return column
}

func computeStatistics(data []float64) map[string]float64 {
	stats := make(map[string]float64)
	stats["Mean"] = mean(data)
	stats["Median"] = median(data)
	stats["StdDev"] = stdDev(data)
	return stats
}

func mean(data []float64) float64 {
	sum := 0.0
	for _, value := range data {
		sum += value
	}
	return sum / float64(len(data))
}

func median(data []float64) float64 {
	sort.Float64s(data)
	n := len(data)
	if n%2 == 0 {
		return (data[n/2-1] + data[n/2]) / 2.0
	}
	return data[n/2]
}

func stdDev(data []float64) float64 {
	m := mean(data)
	sum := 0.0
	for _, value := range data {
		sum += (value - m) * (value - m)
	}
	return (sum / float64(len(data))) / 2.0
}

func printStatistics(stats map[string]float64) {
	fmt.Printf("Mean:   %.2f\n", stats["Mean"])
	fmt.Printf("Median: %.2f\n", stats["Median"])
	fmt.Printf("StdDev: %.2f\n", stats["StdDev"])
}

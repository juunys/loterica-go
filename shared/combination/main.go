package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

const DEFAULT_NAME_FILE = "../lotofacil-combinations.csv"
const MAX_NUMBER = 25
const NUMBER_COMBINATIONS = 15

func combine(n int, k int) [][]string {
	return combineRecursive(1, n, k)
}

func combineRecursive(min, max, k int) [][]string {
	var out [][]string

	if k == 1 {
		for i := min; i <= max; i++ {
			out = append(out, []string{strconv.Itoa(i)})
		}

		return out
	}

	for i := min; i < max; i++ {
		for _, line := range combineRecursive(i+1, max, k-1) {
			out = append(out, append([]string{strconv.Itoa(i)}, line...))
		}
	}

	return out
}

func main() {
	csvFile, err := os.Create(DEFAULT_NAME_FILE)
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	defer csvFile.Close()

	csvwriter := csv.NewWriter(csvFile)
	allCombinations := combine(MAX_NUMBER, NUMBER_COMBINATIONS)

	for _, empRow := range allCombinations {
		_ = csvwriter.Write(empRow)
	}

	csvwriter.Flush()
	defer csvFile.Close()
}

package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func ReadCombinations(filePath string) [][]int {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	reader := csv.NewReader(file)

	combinations := [][]int{}

	for {
		records, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		ints := make([]int, len(records))
		for i, s := range records {
			ints[i], _ = strconv.Atoi(s)
		}
		combinations = append(combinations, ints)
	}
	return combinations
}

func ReadDatabase(filePath string) [][]int {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	reader := csv.NewReader(file)

	gamesRunned := [][]int{}

	for {
		records, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		l := strings.Split(records[0], ";")
		ints := make([]int, len(l))
		for i, s := range l {
			ints[i], _ = strconv.Atoi(s)
		}
		sort.Ints(ints)
		gamesRunned = append(gamesRunned, ints)
	}

	return gamesRunned
}

func testEq(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func main() {
	gamesRunned := ReadDatabase("lotofacil-database.csv")
	myGame := ReadDatabase("mygame-db.csv")

	for index, mg := range myGame {
		available := true
		sort.Ints(mg)
		for _, value := range gamesRunned {
			if testEq(value, mg) {
				available = false
				fmt.Printf("JOGO %d INDISPONÍVEL: %v\n\n", index+1, value)
			}
		}

		if available {
			fmt.Printf("JOGO DISPONÍVEL %d: %v\n\n", index+1, mg)
		}	
	}
}

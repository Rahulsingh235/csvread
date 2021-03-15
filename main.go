package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	//db, _ := GetDB()
	//print(db)
	//fmt.Println("listnew.csv")
	csvfile, err := os.Open("listnew.csv")
	if err != nil {
		fmt.Println("An error encountered ::", err)
		return
	}

	// Parse the file
	reader := csv.NewReader(csvfile)
	//fmt.Println(reader)
	var results [][]string
	// Iterate through the records
	for {
		// Read each record from csv
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return
		}
		// add record to result set
		results = append(results, record)
	}
	for i := 1; i < len(results); i++ {
		fmt.Println(results[i][0])
	}
}

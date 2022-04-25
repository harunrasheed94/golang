package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	filePath := "/Users/haroon-kalilur/Documents/Concourse_Payment_Status_Report_20220311125026.csv"
	csvFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer csvFile.Close()

	csvReader := csv.NewReader(csvFile)
	csvReader.FieldsPerRecord = -1
	csvReader.Comma = ','
	i := 0
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			fmt.Println("Reached end of file")
			break
		}

		if i == 0 {
			i++
			continue
		}
		for k := 0; k < len(record); k++ {
			if len(record[k]) > 0 {
				fmt.Println(record[k])
			}
		}

	}
}

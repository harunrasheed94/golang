package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Reads JPMC header records and data records
*/

// type JPMCRecord struct {
// 	AllRecords []map[string]string `json:"allRecords"`
// 	RecordKeys []string            `json:"recordKeys"`
// }

const (
	HEADER string = "Header"
	DETAIL string = "Detail"
)

func main() {

	csvfile, err := os.Open("/Users/haroon-kalilur/Documents/financedoc.csv")

	if err != nil {
		fmt.Println("error in opening file. Error = ", err)
		return
	}
	defer csvfile.Close()
	csvReader := csv.NewReader(csvfile)
	csvReader.FieldsPerRecord = -1 //FieldsPerRecords should be set to negative so that each record can have any number of fields (i.e. variable number of columns)
	//recordMap := make(map[string]interface{})
	records, err := csvReader.ReadAll()

	if err != nil {
		fmt.Println("Error in reading record. Error = ", err)
		return
	}

	//var jpmcRecord JPMCRecord
	headerRecords := [][]string{}
	detailRecords := [][]string{}

	headerRecordsCnt := 0
	detailRecordsCnt := 0
	for _, record := range records {
		//fmt.Println(record)
		if strings.HasPrefix(record[0], "HACT") {
			headerRecords = append(headerRecords, record)
			headerRecordsCnt = len(record)
		} else {
			detailRecords = append(detailRecords, record)
			detailRecordsCnt = len(record)
		}
	}

	headerKeys := make([]string, headerRecordsCnt)
	detailKeys := make([]string, detailRecordsCnt)
	for i := 0; i < headerRecordsCnt; i++ {
		key := HEADER + strconv.Itoa(i+1)
		headerKeys[i] = key
	}
	for i := 0; i < detailRecordsCnt; i++ {
		key := DETAIL + strconv.Itoa(i+1)
		detailKeys[i] = key
	}

	fmt.Println(headerKeys)
	fmt.Println(detailKeys)
	fmt.Println(headerRecords)
	fmt.Println(detailRecords)
	// allRecords := make([]map[string]string, 0)
	// for _, header := range headerRecords {
	// 	headerMap := make(map[string]string)
	// 	for k := 0; k < len(header); k++ {
	// 		key := HEADER + strconv.Itoa(k+1)
	// 		headerMap[key] = header[k]
	// 	}
	// 	for _, detail := range detailRecords {
	// 		recordMap := make(map[string]string)
	// 		recordMap = headerMap
	// 		for l := 0; l < len(detail); l++ {
	// 			key := DETAIL + strconv.Itoa(l+1)
	// 			recordMap[key] = detail[l]
	// 		}
	// 		allRecords = append(allRecords, recordMap)
	// 		//	fmt.Println(recordMap)
	// 	}
	// }

	//fmt.Println(allRecords)
	//jpmcRecord.AllRecords = allRecords
	//jpmcRecord.RecordKeys = sortedKeys
	//jpmcRecordJson, err := json.Marshal(jpmcRecord)
	//fmt.Println("JSON Object")
	//fmt.Println(string(jpmcRecordJson))

	// var unmarshalledJPMCRecord JPMCRecord
	// json.Unmarshal(jpmcRecordJson, &unmarshalledJPMCRecord)

	// for i, recordMap := range unmarshalledJPMCRecord.AllRecords {
	// 	fmt.Println("Record ", i)
	// 	for _, key := range unmarshalledJPMCRecord.RecordKeys {
	// 		fmt.Print(key, " : ", recordMap[key])
	// 		fmt.Print("  ")
	// 	}
	// 	fmt.Print("\n\n")
	// }
}

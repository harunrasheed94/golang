package main

import (
	"encoding/json"
	"fmt"
	"time"
	//"time"
)

type Person struct {
	Name      string    `json:"name"`
	Age       int64     `json:"age"`
	Timestamp time.Time `json:"timestamp"`
	Job       string    `json:"job,omitempty"`
}

// type FeedV2 struct {
// 	SubmissionID  string                 `json:"submissionID"`
// 	FulfillerID   string                 `json:"fulfillerID"`
// 	SupplierID    string                 `json:"supplierID"`
// 	FeedV2Details []FeedV2SegmentDetails `json:"segmentDetails"`
// }

// type FeedV2SegmentDetails struct {
// 	SupplierItem string `json:"supplierItem"`
// 	IsAvailable  bool   `json:"isAvailable"`
// 	UOM          string `json:"uom"`
// }

var p = fmt.Println

func main() {
	var pSlice []Person

	p1 := Person{
		Name:      "Haroon",
		Age:       27,
		Timestamp: time.Now(),
	}

	p2 := Person{
		Name:      "Rasheed",
		Age:       27,
		Timestamp: time.Now(),
		Job:       "Developer",
	}

	pSlice = append(pSlice, p1, p2)
	b, err := json.Marshal(pSlice)

	if err != nil {
		fmt.Println("Error in JSON Marshaling ", err)
	}

	fmt.Println("JSON Object = ")
	fmt.Println(string(b))

	var persons []Person

	err = json.Unmarshal(b, &persons)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Persons Struct = ")
	fmt.Println(persons)
	// for i, v := range persons {
	// 	fmt.Println("Person Number", i)
	// 	fmt.Println(v.Age, v.Name)
	// 	fmt.Println(v.Timestamp)
	// 	//fmt.Println(v.Timestamp.UnixNano())
	// }

	// jsonfile, err := os.Open("sample_babel_fish_1.json")

	// if err != nil {
	// 	fmt.Println("error in opening file = ", err.Error())
	// 	return
	// }

	// fileContentBytes, _ := ioutil.ReadAll(jsonfile)
	// fmt.Println("\n File Content =  \n", string(fileContentBytes))

	// fmt.Println(len(fileContentBytes), " bytes read")
	// fmt.Println(string(fileContentBytes))
	// var feed FeedV2

	// escapedStr, err4 := strconv.Unquote(string(fileContentBytes))
	// if err4 != nil {
	// 	fmt.Println("error in unquoting. Error = ", err4.Error())
	// 	return
	// }

	// fmt.Println("Escaped string ", escapedStr)
	// err3 := json.Unmarshal([]byte(escapedStr), &feed)

	// if err3 != nil {
	// 	fmt.Println("error in unmarshalling. Error = ", err3.Error())
	// 	return
	// }
	// json.Unmarshal(fileContentBytes, &feed)
	// fmt.Println("\n Unmarshalled json = \n", feed)

}

package main

import (
	//"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
//	"encoding/json"
)

type covidData struct{
	testperformed string
	testpositive string
	date string
	discharged string
	expired string
	region string
	admitted string
}

func main() {
	// Open the file
	mp := make(map[string]covidData)
	var data covidData
	csvfile, err := os.Open("data.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Parse the file
	r := csv.NewReader(csvfile)
	//r := csv.NewReader(bufio.NewReader(csvfile))

	// Iterate through the records

	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if record[2] == "11/03/2020"{
			//slcB, _ := json.Marshal(record)
			//fmt.Println(string(slcB))
		
			
			//fmt.Println(record[0])
			//fmt.Println("asdfsdf")
			data.testpositive = record[0]
			data.testperformed = record[1]
			data.date = record[2]
			data.discharged = record[3]
			data.expired = record[4]
			data.region = record[5]
			data.admitted = record[6]
			mp[record[6]] = data

		}
	}
	fmt.Println(data.testpositive)
	fmt.Println(data.testperformed)
	fmt.Println(data.date)
	fmt.Println(data.discharged)
	fmt.Println(data.expired)
	fmt.Println(data.region)
	fmt.Println(data.admitted)
	fmt.Println("asdfgh")
}
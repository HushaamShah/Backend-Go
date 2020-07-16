package main

import (
	//"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"encoding/json"
)

type covidData struct{
	Testperformed 	string
	Testpositive 	string
	Date 			string
	Discharged 		string
	Expired 		string
	Region 			string
	Admitted 		string
}

type Query struct{
	Region string `json:"region"`
	Date string `json:"date"`
}

type JasonInput struct {
	Query Query `json:"query"`	
}

func main() {
	// Open the file
	csvfile, err := os.Open("data.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Parse the file
	 r := csv.NewReader(csvfile)

	jsonString := `{"query": {"region": "11/03/2020"}}`

	var reading JasonInput
	json.Unmarshal([]byte(jsonString), &reading)
	fmt.Printf("%+v\n", reading)

	fmt.Println(reading.Query.Region)
	fmt.Println(reading.Query.Date)

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
		if record[2] == reading.Query.Date{

			data := covidData{
				Testperformed:	record[0],
				Testpositive:	record[1],
				Date:			record[2],
				Discharged:		record[3],
				Expired: 		record[4],
				Region:			record[5],
				Admitted: 		record[6],
			}

			var jsonData []byte
			jsonData, err = json.MarshalIndent(data, "", "   ")
			if err != nil {
				log.Println(err)
			}
			fmt.Println(string(jsonData))
		}
	}
}
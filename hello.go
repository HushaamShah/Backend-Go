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



type FruitBasket struct {
    Name    string
    Fruit   []string
    Id      int64  `json:"ref"`
	private string // An unexported field is not encoded.
	Region 			string
}

type covidData struct{
	Testperformed 	string
	Testpositive 	string
	Date 			string
	Discharged 		string
	Expired 		string
	Region 			string
	Admitted 		string
}

func main() {
	// Open the file
	//mp := make(map[int]covidData)
	//var data covidData
	//var count int
	//count = 0

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
			// fmt.Println("asdfsdf")
			// data.testpositive = record[0]
			// data.testperformed = record[1]
			// data.date = record[2]
			// data.discharged = record[3]
			// data.expired = record[4]
			// data.region = record[5]
			// data.admitted = record[6]

			data := covidData{
				Testperformed:	record[0],
				Testpositive:	record[1],
				Date:			record[2],
				Discharged:		record[3],
				Expired: 		record[4],
				Region:			record[5],
				Admitted: 		record[6],
			}
			// basket := FruitBasket{
			// Name:    record[5],
			// Fruit:   []string{"Apple", "Banana", "Orange"},
			// Id:      999,
			// Region:	record[5],
			// private: data.discharged,	
			// }
			var jsonData []byte
			jsonData, err = json.MarshalIndent(data, "", "")
			if err != nil {
				log.Println(err)
			}
			fmt.Println(string(jsonData))
		}
	}
	// fmt.Println(data.testpositive)
	// fmt.Println(data.testperformed)
	// fmt.Println(data.date)
	// fmt.Println(data.discharged)
	// fmt.Println(data.expired)
	// fmt.Println(data.region)
	// fmt.Println(data.admitted)
	// fmt.Println("asdfgh")


	// basket := FruitBasket{
	// 	Name:    data.region,
	// 	Fruit:   []string{"Apple", "Banana", "Orange"},
	// 	Id:      999,
	// 	private: data.discharged,
	// }


	// var jsonData []byte
	// jsonData, err = json.Marshal(data)
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(string(jsonData))
}
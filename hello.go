package main

import (
	//"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"encoding/json"
	"container/list"
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
	//mp := make(map[int]covidData)
	asd := list.New()
	var data covidData
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
			fmt.Println("asdfsdf")
			data.testpositive = record[0]
			data.testperformed = record[1]
			data.date = record[2]
			data.discharged = record[3]
			data.expired = record[4]
			data.region = record[5]
			data.admitted = record[6]
			//mp[count] = data
			//count++
			asd.PushBack(data)
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


	// empData, err := json.Marshal(asd)   
    // if err != nil {
    //     fmt.Println(err.Error())
    //     return
    // }
     
    // jsonStr := string(empData)
    // fmt.Println("The JSON data is:")
	// fmt.Println(jsonStr)
	
	// for e := asd.Front(); e != nil; e = e.Next() {
	// 	fmt.Println(e.Value)
	// 	jsonStr := string(e.Value)
	// 	fmt.Println("The JSON data is:")
	// 	fmt.Println(jsonStr)
	// }

	slcB, _ := json.Marshal(asd)
	fmt.Println(string(slcB))

}
// Code Related to Server Connection was used from:
// https://www.linode.com/docs/development/go/developing-udp-and-tcp-clients-and-servers-in-go/

package main

import (
        "bufio"
        "fmt"
        "net"
        "os"
        "strings"
		"io"
		"log"
		"encoding/csv"
		"encoding/json"
)

type Query struct{
	Region string `json:"region"`
	Date string `json:"date"`
}

type JasonInput struct {
	Query Query `json:"query"`	
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
	arguments := os.Args
	if len(arguments) == 1 {
			fmt.Println("Please provide port number")
			return
	}

	PORT := ":" + arguments[1]
	l, err := net.Listen("tcp", PORT)
	if err != nil {
			fmt.Println(err)
			return
	}
	defer l.Close()

	c, err := l.Accept()
	if err != nil {
			fmt.Println(err)
			return
	}

	for {
			netData, err := bufio.NewReader(c).ReadString('\n')
			if err != nil {
				fmt.Println(err)
				return
			}
			//STOP command exits the Server
			if strings.TrimSpace(string(netData)) == "STOP" {
					fmt.Println("Exiting TCP server!")
					return
			}

			//Checcking if JSON is in valid format
			if !json.Valid([]byte(netData)) {
				c.Write([]byte("Invalid JSON Query \n Try Again \n"))
				continue
			}
			
			//Stores JSON data into JasonInput Structure
			var reading JasonInput
			json.Unmarshal([]byte(netData), &reading)
			fmt.Printf("%+v\n", reading)


			//Opening CSV file
			csvfile, err := os.Open("data.csv")
			if err != nil {
				//Checks for error in opening CSV file
				log.Fatalln("Couldn't open the csv file", err)
			}
			r := csv.NewReader(csvfile)		

			for {
				// Read each record from csv
				record, err := r.Read()
				if err == io.EOF {
					//Checking for End of file
					break
				}
				if err != nil {
					log.Fatal(err)
				}
				//Comparing the Query to their respective column
				//And storing data in covidData Structure
				if (record[2] == reading.Query.Date || record[5] == reading.Query.Region){
					data := covidData{
						Testperformed:	record[0],
						Testpositive:	record[1],
						Date:			record[2],
						Discharged:		record[3],
						Expired: 		record[4],
						Region:			record[5],
						Admitted: 		record[6],
					}
					//Matching data is converted to JSON and sent
					//Loop repeats untill End of file is found
					//Nothing is stored in covidData permanently
					var jsonData []byte
					jsonData, err = json.MarshalIndent(data, "", "   ")
					if err != nil {
						log.Println(err)
					}
					fmt.Println(string(jsonData))
					c.Write([]byte(jsonData))
				}
			}
        }
}
    
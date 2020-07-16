package main

import (
        "bufio"
        "fmt"
        "net"
        "os"
        "strings"
		"time"
		"io"
		"log"
		"encoding/csv"
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

func main() {

	csvfile, err := os.Open("data.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	r := csv.NewReader(csvfile)

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
				text := strings.TrimSuffix(netData, "\n") // removing "\n from netData"
				fmt.Println("netData = ",netData)
				fmt.Println("text = ",text)
                if err != nil {
                        fmt.Println(err)
                        return
                }
                if strings.TrimSpace(string(netData)) == "STOP" {
                        fmt.Println("Exiting TCP server!")
                        return
                }

                fmt.Print("-> ", string(netData))
                t := time.Now()
                myTime := t.Format(time.RFC3339) + "\n"
				c.Write([]byte(string(netData)))
				fmt.Print(myTime)
				
				//var asd string = string(netData)

				for {
					// Read each record from csv
					record, err := r.Read()
					if err == io.EOF {
						break
					}
					if err != nil {
						log.Fatal(err)
					}
					if (record[2] == text || record[5] == text){
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
						c.Write([]byte(jsonData))
					}
				}
        }
}
    
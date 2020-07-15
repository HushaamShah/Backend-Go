package main

import (
        "bufio"
        "fmt"
        "net"
        "os"
        "strings"
		"time"
//		"io"
		"log"
		"encoding/csv"
)

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
				c.Write([]byte("myTime"))
				fmt.Print(myTime)

				// for {
				// 	// Read each record from csv
				// 	record, err := r.Read()
				// 	if err == io.EOF {
				// 		break
				// 	}
				// 	if err != nil {
				// 		log.Fatal(err)
				// 	}
				// 	if record[5] == string(netData){
				// 		//slcB, _ := json.Marshal(record)
				// 		//fmt.Println(string(slcB))
					
						
				// 		fmt.Println(record)
				// 		c.Write([]byte(record))
				// 		// fmt.Println("asdfsdf")
				// 		// data.testpositive = record[0]
				// 		// data.testperformed = record[1]
				// 		// data.date = record[2]
				// 		// data.discharged = record[3]
				// 		// data.expired = record[4]
				// 		// data.region = record[5]
				// 		// data.admitted = record[6]
				// 		//mp[count] = data
				// 		//count++
				// 		// asd.PushBack(data)
				// 	}
				// }
        }
}
    
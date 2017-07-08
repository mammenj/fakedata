package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) == 3 {

		fmt.Println("number of records -> " + os.Args[1])
		fmt.Println("idpattern -> " + os.Args[2])
	} else {
		log.Fatal("usage: appName noOfRecords idPattern")
		// fakeleads.exe 10000 xdfsd
	}

	numRecords, interr := strconv.Atoi(os.Args[1])
	if interr != nil {
		log.Fatalln("error in integer ", interr)
	}
	idPattern := os.Args[2]

	file, err := os.Create("fakeleads.csv")
	checkError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()
	value := []string{"rNLID", "firstName", "lastName", "email", "address1", "address2", "address3"}
	err = writer.Write(value)
	for i := 0; i < numRecords; i++ {
		rnlid := idPattern + strconv.Itoa(i)
		email := idPattern + strconv.Itoa(i) + "@email.com"
		address1 := "address" + idPattern + strconv.Itoa(i) + strconv.Itoa(rand.Intn(100))
		address2 := "address" + idPattern + strconv.Itoa(i) + strconv.Itoa(rand.Intn(100))
		address3 := "address" + idPattern + strconv.Itoa(i) + strconv.Itoa(rand.Intn(100))
		value := []string{rnlid, "john", "doe", email, address1, address2, address3}
		err := writer.Write(value)
		checkError("Cannot write to file", err)
	}
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

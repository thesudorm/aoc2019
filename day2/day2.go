package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"log"
	"strconv"
	//"io/ioutil"
)

func main(){
	var data []int64

	file, err := os.Open("day2.txt")
	if err != nil {
		log.Fatal(err);
	}

	/*
	contents, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err);
	}

	fmt.Println(contents)
	fmt.Println()
	*/

	input, err := csv.NewReader(file).ReadAll()

	fmt.Println(input)
	fmt.Println(len(input[0]))
	test := input[0]

	for _, item := range test {
		value, err := strconv.ParseInt(item, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, value)
	}

	for i := 0; i < len(data) - 3; i += 4 {
		firstValue := data[data[i + 1]]
		secondValue := data[data[i + 2]]
		position := data[i + 3]

		if data[i] == 1 { // Add
			data[position] = firstValue + secondValue
		} else if data[i] == 2 { // Multiply
			data[position] = firstValue * secondValue
		} else if data[i] == 99 {
			fmt.Println("Done")
			break
		} else {
			fmt.Println(data[i])
			log.Fatal("Something went wrong")
		}
	}

	fmt.Println(data)
}

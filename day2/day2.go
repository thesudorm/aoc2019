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
	fmt.Println(len(input))
	test := input[0]

	for _, item := range test {
		value, err := strconv.ParseInt(item, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, value)
	}

	fmt.Println(data)
}

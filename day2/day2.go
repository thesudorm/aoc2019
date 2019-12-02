package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"log"
	//"strconv"
)

func main(){

	file, err := os.Open("day2.txt")
	if err != nil {
		log.Fatal(err);
	}

	//contents, err := ioutil.ReadAll(file)

	//var data []int64
	//var value int64

	input, err := csv.NewReader(file).ReadAll()

	for _, item := range input {
		//value = strconv.Atoi(item)
		fmt.Println(item)
	}
}

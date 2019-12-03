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

	originalData := data

	for j := 0; j < 100; j++ {
		for k := 0; k < 100; k++ {
			data = originalData
			data[1] = int64(j)
			data[2] = int64(k)

			for i := 0; i < len(data) - 3; i += 4 {
				firstValue := data[data[i + 1]]
				secondValue := data[data[i + 2]]
				position := data[i + 3]

				if data[i] == 1 { // Add
					data[position] = firstValue + secondValue
				} else if data[i] == 2 { // Multiply
					data[position] = firstValue * secondValue
				} else if data[i] == 99 {
					break
				} else {
					break;
				}
			}

			fmt.Println("Noun: ", j, "Verb: ", k, "Answer: ", data[0])
			fmt.Println(data)

			if data[0] == 19690720 {
				fmt.Println(j," ", k)
				break;
			}

		}
	}
}

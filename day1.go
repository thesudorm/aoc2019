package main

import (
    "fmt"
    "os"
    "log"
    "bufio"
    "strconv"
)

func main(){
    var total int64 = 0


    file, err := os.Open("input/day1.txt")
    if err != nil {
        log.Fatal(err)
    }

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        value, err := strconv.ParseInt(scanner.Text(), 10, 64)
        if err != nil {
            panic(err)
        }
        
        for (value > 0){
            value = value / 3 - 2
            if(value > 0){
                total += value
            }
        }
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    fmt.Println(total)
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func opOne(op *[]int, index int) {
	sum := (*op)[(*op)[index + 1]] + (*op)[(*op)[index + 2]]
	(*op)[(*op)[index + 3]] = sum
}

func opTwo(op *[]int, index int) {
	product := (*op)[(*op)[index + 1]] * (*op)[(*op)[index + 2]]
	(*op)[(*op)[index + 3]] = product
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	line, err := reader.ReadString('\n')
	line = strings.TrimSuffix(line, "\n")
	if err != nil {
		log.Fatal(err)
	}
	stringSplit := strings.Split(line, ",")
	var op []int
	for _, e := range stringSplit {
		v, err := strconv.Atoi(e)
		if (err != nil) {
		    log.Fatal(err)
		}
		op = append(op, v)
	}

	for i := 0; i <= len(op); i += 4 {
		if op[i] == 99 {
			fmt.Println("99 found.. exiting")
			break
		}
		if op[i] == 1 {
			opOne(&op, i)
		} else if op[i] == 2 {
			opTwo(&op, i)
		}
	}
	fmt.Printf("%v", op)
}

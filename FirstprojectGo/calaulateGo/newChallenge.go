package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin) //ใช้ทำอะไร ??

func getInput(promt string) float64 {
	fmt.Printf("%v", promt)
	input, _ := reader.ReadString('\n')
	value, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		massage, _ := fmt.Scanf("%v must number only", promt)
		panic(massage)
	}
	return value
}
func getOperator() string {
	fmt.Print("Operator is (+ - * /) : ")
	op, _ := reader.ReadString('\n')
	return strings.TrimSpace(op)
}
func add(value1, value2 float64) float64 {
	return value2 + value1
}
func minus(value1, value2 float64) float64 {
	return value2 - value1
}
func multi(value1, value2 float64) float64 {
	return value2 * value1
}
func div(value1, value2 float64) float64 {
	return value2 / value1
}
func main() {
	num := getInput("จำนวนที่ต้องการรับ = ")
	var count float64
	var result float64 = 0
	result = getInput("Input Value  = ")
	for {
		if num-1 > count {
			operator := getOperator()
			value := getInput("Input Value  = ")
			// ส่งค่าไป func เพื่อคำนวณ
			if operator == "+" {
				result = add(value, result)
			} else if operator == "-" {
				result = minus(value, result)
			} else if operator == "*" {
				result = multi(value, result)
			} else if operator == "/" {
				result = div(value, result)
			} else {
				fmt.Printf("Don't have operator yet")
			}
			count++

		} else {
			fmt.Println("Result : ", result)
			break
		}
	}
}

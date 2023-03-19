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
	return value1 + value2
}
func minus(value1, value2 float64) float64 {
	return value1 - value2
}
func multi(value1, value2 float64) float64 {
	return value1 * value2
}
func div(value1, value2 float64) float64 {
	return value1 / value2
}

func main() {
	value1 := getInput("Input Value1 = ")
	operator := getOperator()
	value2 := getInput("Input Value2 = ")

	fmt.Println("Value 1 : ", value1, "Value 2 :", value2)
	var sum float64

	if operator == "+" {
		sum = add(value1, value2)
	} else if operator == "-" {
		sum = minus(value1, value2)
	} else if operator == "*" {
		sum = multi(value1, value2)
	} else if operator == "/" {
		sum = div(value1, value2)
	} else {
		fmt.Printf("Don't have operator yet")
	}
	fmt.Println(sum)
}

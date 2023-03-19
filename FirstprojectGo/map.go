package main

import (
	"fmt"
)

var product = make(map[string]float32)

func main() {

	// map เก็บค่าตัวแปรในรูปของ Key และ value
	// Add data ลงใน map
	// VariableName["Key"] = value
	product["Anime"] = 90
	fmt.Println("Profuct = ", product)
	product["Manga"] = 120
	product["Manhwa"] = 85.2
	fmt.Println("Profuct = ", product)

	// delete data from map
	// delete(mapVariable, Key)
	delete(product, "Anime")
	fmt.Println("Profuct = ", product)

	// Upadate data
	product["Manhwa"] = 185.2
	fmt.Println("Profuct = ", product)

	// access data
	value1 := product["Manga"]
	fmt.Println("Value = ", value1)

	// กำหนดค่า default ใน mapping
	CourseName := map[string]string{"101": "Java", "102": "Go land"}
	fmt.Println("Course = ", CourseName)

}

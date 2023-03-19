package main

import "fmt"

func main() {
	// same array but can add new data
	var book []string
	book = []string{"Java"}
	fmt.Println(book)
	// add new data to array
	book = append(book, "Go Lang")
	fmt.Println(book)
	book = append(book, "Html", "Css", "Vue", "Javascript")
	fmt.Println(book)
	// Choose data from array to show
	// and
	// add data that choosed to new array
	// bookShow := book[1:4]
	// fmt.Println(bookShow)
	bookShow := book[1:5]
	fmt.Println(bookShow)
	bookShow = bookShow[1:3]
	fmt.Println(bookShow)
}

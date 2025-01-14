package main

import (
	"fmt"
	"os"
	"path"
	c "phoneBook/crud"
)

func main() {
	var data = []c.Entry{
		{Name: "AKIE", Surname: "EAFGT", Tel: "768"},
		{Name: "John", Surname: "Doe", Tel: "123"},
		{Name: "Jane", Surname: "Smith", Tel: "654"},
		{Name: "Alice", Surname: "Brown", Tel: "987"},
		{Name: "Bob", Surname: "Johnson", Tel: "456"},
		{Name: "Charlie", Surname: "Lee", Tel: "321"},
		{Name: "Diana", Surname: "Taylor", Tel: "789"},
		{Name: "Eve", Surname: "Harris", Tel: "159"},
		{Name: "Frank", Surname: "Clark", Tel: "753"},
		{Name: "Grace", Surname: "Walker", Tel: "852"}}
		
	data = append(data, c.Entry{Name: "Matthew", Surname: "Mahlyshev", Tel: "111"})
	c.Populate(99, &data)
	arguments := os.Args

	if len(arguments) == 1 {
		exe := path.Base(arguments[0])
		fmt.Printf("Usage: %s search|list <arguments>\n", exe)
		return
	}

	switch arguments[1] {
	case "search":
		if len(arguments) != 3 {
			fmt.Println("Usage: search number")
			return
		}
		result := c.Search(arguments[2], data)
		if result == nil {
			fmt.Println("Number not found:", arguments[2])
			return
		}
		fmt.Println(*result)
	case "list":
		c.List(data)
	default:
		fmt.Println("Not a valid option")
	}
}

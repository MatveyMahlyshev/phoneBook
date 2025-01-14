package main

import (
	// "phoneBook/csv"
	"fmt"
	"os"
	"path"
	c "phoneBook/users"
)

func main() {
	// csv.Run()
	var data []c.Entry	
	c.Populate(5, &data)
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

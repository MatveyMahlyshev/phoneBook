package main

import (
	// "phoneBook/csv"
	"fmt"
	"os"
	c "phoneBook/users"
	"strings"
)

func main() {
	arguments := os.Args
	// arguments := []string{"asd", "insert", "Matvey", "Mahlyshev", "89672775456"}
	if len(arguments) == 1 {
		fmt.Println("Usage: insert|delete|search|list <arguments>")
		return
	}

	_, err := os.Stat(c.CSVDATA)
	if err != nil {
		fmt.Println("Creating", c.CSVDATA)
		f, err := os.Create(c.CSVDATA)
		if err != nil {
			f.Close()
			fmt.Println(err)
			return
		}
		f.Close()
	}

	fileInfo, _ := os.Stat(c.CSVDATA)
	mode := fileInfo.Mode()
	if !mode.IsRegular() {
		fmt.Println(c.CSVDATA, "not a regular file!")

	}
	_, err = c.ReadCSVFile(c.CSVDATA)
	if err != nil {
		fmt.Println(err)
	}

	err = c.CreateIndex()
	if err != nil {
		fmt.Println("cannot create index")
	}

	switch arguments[1] {
	case "insert":
		if len(arguments) != 5 {
			fmt.Println("Usage: insert Name Surname Telephone")
			return
		}
		t := strings.ReplaceAll(arguments[4], "-", "")

		if !c.MatchTel(t) {
			fmt.Println("Not a valid telephone number:", t)
			return
		}

		temp := c.Inits(arguments[2], arguments[3], t)
		if temp != nil {
			err := c.Insert(temp)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	case "delete":
		if len(arguments) != 3 {
			fmt.Println("Usage: delete Number")
			return
		}

		t := strings.ReplaceAll(arguments[2], "-", "")
		if !c.MatchTel(t) {
			fmt.Println("Not a valid telephone number:", t)
			return
		}

		err := c.DeleteEntry(t)
		if err != nil {
			fmt.Println(err)
		}
	case "search":
		if len(arguments) != 3 {
			fmt.Println("Usage: search Number")
			return
		}
		t := strings.ReplaceAll(arguments[2], "-", "")
		if !c.MatchTel(t) {
			fmt.Println("Not a valid telephone number:", t)
			return
		}
		temp := c.Search(t)
		if temp == nil {
			fmt.Println("Number not found:", t)
			return
		}
		fmt.Println(*temp)
	case "list":
		c.List()
	default:
		fmt.Println("Not a valid option")
	}
}

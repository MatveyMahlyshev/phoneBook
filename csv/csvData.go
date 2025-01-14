package csv

import (
	"fmt"
	"os"
)

func Run() {
	if len(os.Args) != 3 {
		fmt.Println("csvData input output")
		return
	}

	input := os.Args[1]
	output := os.Args[2]
	lines, err := ReadCSVFile(input)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, line := range lines {
		temp := Record{
			Name:       line[0],
			Surname:    line[1],
			Number:     line[2],
			LastAccess: line[3],
		}
		MyData = append(MyData, temp)
		fmt.Println(temp)
	}

	err = SaveCSVFile(output)
	if err != nil {
		fmt.Println(err)
		return
	}
}

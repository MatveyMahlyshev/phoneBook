package crud

import (
	"fmt"
	"strconv"
)

type Entry struct {
	Name    string
	Surname string
	Tel     string
}

func List(data []Entry) {
	for _, v := range data {
		fmt.Println(v)
	}
}

func Search(key string, data []Entry) *Entry {
	for i, v := range data {
		if v.Tel == key {
			return &data[i]
		}
	}
	return nil
}

func Populate(n int, data *[]Entry) {
	for i := 0; i < n; i++ {
		*data = append(*data, Entry{
			Name:    generateData(4),
			Surname: generateData(5),
			Tel:     strconv.Itoa(random(100, 999)),
		})
	}
}

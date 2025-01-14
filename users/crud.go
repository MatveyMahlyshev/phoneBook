package users

import (
	"fmt"
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

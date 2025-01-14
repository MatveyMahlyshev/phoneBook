package users

import (
	"strconv"
	"strings"
)

var Index map[string]int
var CSVDATA = "csv.data"

func CreateIndex() error {
	Index = make(map[string]int)
	for i, k := range Data {
		key := k.Tel
		Index[key] = i
	}
	return nil
}

func MatchTel(t string) bool {
	for _, v := range t {
		_, err := strconv.Atoi(string(v))
		if err != nil {
			return false
		}
	}
	return true
}

func Inits(name string, surname string, telephone string) *Entry {
	name = strings.ReplaceAll(name, " ", "")
	surname = strings.ReplaceAll(surname, " ", "")
	if name != "" && surname != "" {
		return &Entry{Name: name, Surname: surname, Tel: telephone}
	}
	return nil
}

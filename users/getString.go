package users

import (
	"math/rand"
	"strconv"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func generateData(len int64) string {
	temp := ""
	var i int64 = 1
	for {
		genCode := random(65, 90)
		genChar := string(byte(genCode))
		temp += genChar
		if i == len {
			break
		}
		i++
	}
	return temp
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

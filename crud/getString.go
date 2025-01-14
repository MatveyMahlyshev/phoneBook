package crud

import "math/rand"

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

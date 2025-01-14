package users

import (
	"encoding/csv"
	"os"
	"strconv"
	"time"
)

func ReadCSVFile(filepath string) ([][]string, error) {
	_, err := os.Stat(filepath)
	if err != nil {
		return nil, err
	}

	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	for _, l := range lines {
		Data = append(Data, Entry{Name: l[0], Surname: l[1], Tel: l[2], LastAccess: strconv.FormatInt(time.Now().Unix(), 10)})
	}

	return lines, nil
}

func SaveCSVFile(filepath string) error {
	csvfile, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer csvfile.Close()

	csvwriter := csv.NewWriter(csvfile)
	csvwriter.Comma = ','
	for _, row := range Data {
		temp := []string{row.Name, row.Surname, row.Tel, row.LastAccess}
		_ = csvwriter.Write(temp)
	}
	csvwriter.Flush()
	return nil

}

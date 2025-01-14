package csv

import (
	"encoding/csv"
	"os"
)

type Record struct {
	Name       string
	Surname    string
	Number     string
	LastAccess string
}

var MyData = []Record{}

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

	return lines, nil
}

func SaveCSVFile(filepath string) error {
	csvfile, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer csvfile.Close()

	csvwriter := csv.NewWriter(csvfile)
	csvwriter.Comma = '\t'
	for _, row := range MyData {
		temp := []string{row.Name, row.Surname, row.Surname, row.LastAccess}
		_ = csvwriter.Write(temp)
	}
	csvwriter.Flush()
	return nil

}

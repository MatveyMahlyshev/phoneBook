package users

import (
	"fmt"
	"strconv"
	"time"
)

type Entry struct {
	Name       string
	Surname    string
	Tel        string
	LastAccess string
}

var Data []Entry

func List() {
	for _, v := range Data {
		fmt.Println(v)
	}
}

func Search(key string) *Entry {
	i, ok := Index[key]
	if !ok {
		return nil
	}
	Data[i].LastAccess = strconv.FormatInt(time.Now().Unix(), 10)
	return &Data[i]
}

func Insert(pS *Entry) error {
	// если уже существует, то не добавлять
	_, ok := Index[(*pS).Tel]
	if ok {
		return fmt.Errorf("%s already exists", pS.Tel)
	}
	Data = append(Data, *pS)
	// обновить индекс
	_ = CreateIndex()
	err := SaveCSVFile(CSVDATA)
	if err != nil {
		return err
	}
	return nil
}

func DeleteEntry(key string) error {
	i, ok := Index[key]
	if !ok {
		return fmt.Errorf("%s cannot be found!", key)
	}
	Data = append(Data[:i], Data[i+1:]...)
	// обновить индекс — ключ больше не существует
	delete(Index, key)
	err := SaveCSVFile(CSVDATA)
	if err != nil {
		return err
	}
	return nil
}

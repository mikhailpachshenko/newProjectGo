package main

import (
	"errors"
	"fmt"
	"strconv"
)

type Storage struct {
	Name  string
	Age   int
	Grade int
}

func (s Storage) Info() string {
	return fmt.Sprintf("%s, %s, %s.\n", s.Name, strconv.Itoa(s.Age), strconv.Itoa(s.Grade))
}

func (s *Storage) Put(list map[string]*Storage, input *Storage) (int, error) {
	list[input.Name] = input

	if list[input.Name] == nil {
		return -1, errors.New("Ошибка при добавлении в базу данных")
	} else {
		return 0, nil
	}
}

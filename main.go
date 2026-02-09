package main

import (
	"fmt"
	"os"
	"encoding/json"
)

type Student struct {
	Id 		int 		`json:"Id"`
	Name 	string	`json:"Name"`
	Entry int 		`json:"Entry"`
	Time 	int 		`json:"Time"`
}

var students []Student

func Load() {
	if file, err := os.Open("students.json"); err != nil {
		fmt.Println("Ошибка загрузки: ", err)
		return
	} else {
		defer file.Close()
		if err := json.NewDecoder(file).Decode(&students); err != nil {
			fmt.Println("Ошибка декодера: ", err)
		} 
	}
}

func isAnomaly(s Student, ent, time int) bool {
	return (s.Entry*2 < ent || s.Entry*2 > ent*3) && (s.Time*2 < time || s.Time*2 > time*3)
}

func main() {
	Load()

	if len(students) == 0 {
		fmt.Println("Нет данных")
		return
	}

	var sumEnt, sumTim int

	for _, s := range students {
		sumEnt += s.Entry
		sumTim += s.Time
	}

	avgEnt := sumEnt / len(students)
	avgTim := sumTim / len(students)

	fmt.Println("Средние значения:")
	fmt.Println("Входы:", avgEnt)
	fmt.Println("Время:", avgTim)

	fmt.Println("\nАномальные пользователи:")

	for _, s := range students {
		if isAnomaly(s, avgEnt, avgTim) {
			fmt.Printf("ID: %d, Name: %s, Входы: %d, Время: %d\n", s.Id, s.Name, s.Entry, s.Time)
		}
	}
}


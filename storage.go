package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func Load() {
	if file, err := os.Open("students.json"); err != nil {
		fmt.Println("Ошибка загрузки: ", err)
		return
	}
	defer file.Close()
	json.NewDecoder(file).Decode(&students)
}

func Save() {
	if file, err := os.Create("results.json"); err != nil {
		fmt.Println("Ошибка сохранения: ", err)
		return
	}
	defer file.Close()
	json.NewEncoder(file).Encode(results)
}

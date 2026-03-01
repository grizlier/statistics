package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"sort"
)

type Student struct {
	Id    int    `json:"Id"`
	Name  string `json:"Name"`
	Entry int    `json:"Entry"`
	Time  int    `json:"Time"`
}

type Anomaly struct {
	Student Student
	Score   float64
}

var students []Student

func Load() {
	file, err := os.Open("students.json")
	if err != nil {
		fmt.Println("Ошибка загрузки:", err)
		return
	}
	defer file.Close()

	if err := json.NewDecoder(file).Decode(&students); err != nil {
		fmt.Println("Ошибка декодера:", err)
	}
}

func stdDev(values []int, mean float64) float64 {
	var sum float64
	for _, v := range values {
		diff := float64(v) - mean
		sum += diff * diff
	}
	return math.Sqrt(sum / float64(len(values)-1))
}

func median(values []int) float64 {
	sort.Ints(values)
	n := len(values)

	if n%2 == 0 {
		return float64(values[n/2-1]+values[n/2]) / 2
	}
	return float64(values[n/2])
}

func main() {
	Load()

	if len(students) == 0 {
		fmt.Println("Нет данных")
		return
	}

	var entries []int
	var times []int
	var sumEnt, sumTim int

	for _, s := range students {
		entries = append(entries, s.Entry)
		times = append(times, s.Time)
		sumEnt += s.Entry
		sumTim += s.Time
	}

	avgEnt := float64(sumEnt) / float64(len(students))
	avgTim := float64(sumTim) / float64(len(students))

	stdEnt := stdDev(entries, avgEnt)
	stdTim := stdDev(times, avgTim)

	medEnt := median(entries)
	medTim := median(times)

	fmt.Println("Средние значения:")
	fmt.Println("Входы:", avgEnt)
	fmt.Println("Время:", avgTim)

	fmt.Println("\nМедиана:")
	fmt.Println("Входы:", medEnt)
	fmt.Println("Время:", medTim)

	var anomalies []Anomaly

	if stdEnt == 0 || stdTim == 0 {
		fmt.Println("Невозможно вычислить аномалии: стандартное отклонение равно 0")
		return
	}
	
	for _, s := range students {
		score := math.Abs((float64(s.Entry)-avgEnt)/stdEnt) +
			math.Abs((float64(s.Time)-avgTim)/stdTim)

		if score > 2 {
			anomalies = append(anomalies, Anomaly{
				Student: s,
				Score:   score,
			})
		}
	}

	sort.Slice(anomalies, func(i, j int) bool {
		return anomalies[i].Score > anomalies[j].Score
	})

	percent := float64(len(anomalies)) / float64(len(students)) * 100
	fmt.Printf("\nПроцент аномальных пользователей: %.2f%%\n", percent)

	fmt.Println("\nАномальные пользователи (по степени отклонения):")
	for _, a := range anomalies {
		fmt.Printf("ID: %d, Name: %s, Score: %.2f\n",
			a.Student.Id,
			a.Student.Name,
			a.Score)
	}
}
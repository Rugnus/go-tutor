package main

import (
	"fmt"
	"math"
)

func main() {
	const IMTPower = 2
	var userHeight float64
	var userKg float64
	fmt.Print(`__ Калькулятор индекса массы тела __
Введите свой рост в сантиметрах: `)

	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес: ")
	fmt.Scan(&userKg)
	IMT := userKg / math.Pow(userHeight/100, IMTPower)
	result := fmt.Sprintf("Ваш инедекс массы тела: %.1f", IMT)
	fmt.Println(result)
}

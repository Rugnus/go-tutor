package main

import (
	"fmt"
	"math"
)

func main() {

	userHeight, userKg := getUserInput()
	IMT := calculateIMT(userHeight, userKg)
	outputResult(IMT)

	switch {
	case IMT < 16:
		fmt.Println("У вас выраженный дефицит массы")
	case IMT < 18.5:
		fmt.Println("У вас недостаточная масса тела")
	case IMT < 25:
		fmt.Println("У вас нормальный вес")
	case IMT < 30:
		fmt.Println("У вас избыточная масса тела")
	default:
		fmt.Println("У вас ожирение")
	}
}

func getUserInput() (float64, float64) {

	var userHeight float64
	var userKg float64
	fmt.Print(`__ Калькулятор индекса массы тела __
Введите свой рост в сантиметрах: `)

	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес: ")
	fmt.Scan(&userKg)
	return userHeight, userKg
}

func calculateIMT(userHeight float64, userKg float64) float64 {
	IMT := userKg / math.Pow(userHeight/100, 2)
	return IMT
}

func outputResult(IMT float64) {
	result := fmt.Sprintf("Ваш инедекс массы тела: %.1f", IMT)
	fmt.Println(result)
}

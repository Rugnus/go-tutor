package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	for {

		userHeight, userKg := getUserInput()
		IMT, err := calculateIMT(userHeight, userKg)
		if err != nil {
			// fmt.Println("Не правильно указан вес или рост")
			// continue
			panic("Не правильно указан вес или рост")
		}
		outputResult(IMT)

		isRepeatCalculation := checkRepeatCalculation()

		if !isRepeatCalculation {
			break
		}
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

func calculateIMT(userHeight float64, userKg float64) (float64, error) {
	if userKg <= 0 || userHeight <= 0 {
		return 0, errors.New("NO_PARAMS_ERROR")
	}
	IMT := userKg / math.Pow(userHeight/100, 2)
	return IMT, nil
}

func outputResult(IMT float64) {
	result := fmt.Sprintf("Ваш инедекс массы тела: %.1f", IMT)
	fmt.Println(result)
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

func checkRepeatCalculation() bool {
	var confirmation string
	fmt.Print("Хотите ли вы прошложить калькулятор? (y/n): ")
	fmt.Scan(&confirmation)

	if confirmation == "y" || confirmation == "Y" {
		return true
	}
	return false

}

package main

import "fmt"

var rates = map[string]float64{
	"EUR": 90.48,
	"USD": 76.01,
	"RUB": 1,
}

func askUser()(string, string, float64) {
	var currencyChange string
	var currencyTarget string
	var summ float64

	for {
		fmt.Print("Какую валюту вы хотите поменять (EUR/USD/RUB): ")
		fmt.Scan(&currencyChange)
		if _, ok := rates[currencyChange]; ok {
			break
		}
	}

	for {
		fmt.Print("На какую валюту вы хотите сконвертировать (EUR/USD/RUB): ")
		fmt.Scan(&currencyTarget)
		if _, ok := rates[currencyTarget]; ok && currencyTarget != currencyChange {
			break
		}

	}

	for {
		fmt.Print("Введите сумму для обмена: ")
		fmt.Scan(&summ)
		if summ > 0 {
			break
		}
	}
	return currencyChange, currencyTarget, summ
}

func countCurrency(currencyChange string, currencyTarget string, summ float64) float64 {
	return summ * rates[currencyChange] / rates[currencyTarget]
}

func main() {
	currencyChange, currencyTarget, summ := askUser()
	result := countCurrency(currencyChange, currencyTarget, summ)

	fmt.Printf("Сумма сконвертированной валюты из %s в %s составит %.2f\n", currencyChange, currencyTarget, result)

}

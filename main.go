package main

import "fmt"

func askUser ()(string, string, float64) {
	var currencyChange string
	var currencyTarget string
	var summ float64

	fmt.Print("Какую валюту вы хотите поменять?")
	fmt.Scan(&currencyChange)
	fmt.Print("На какую валюту вы хотите сконвертировать?")
	fmt.Scan(&currencyTarget)
	fmt.Print("Какую сумму?")
	fmt.Scan(&summ)

	return currencyChange, currencyTarget, summ
}

func countCurrency (currencyChange string, currencyTarget string, summ float64) {

}

func main () {
	const exchangeUSDtoEUR float64 = 0.84
	const exchangeUSDtoRUB float64 = 76.01
	const exchangeEURtoRUB float64 = exchangeUSDtoRUB / exchangeUSDtoEUR

	// currencyChange, currencyTarget, summ := askUser()

	fmt.Println("Конвертация евро в рубль= ", exchangeEURtoRUB)
}
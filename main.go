package main

import "fmt"

func main () {
	const exchangeUSDtoEUR float64 = 0.84
	const exchangeUSDtoRUB float64 = 76.01
	const exchangeEURtoRUB float64 = exchangeUSDtoRUB / exchangeUSDtoEUR

	fmt.Println("Конвертация евро в рубль= ", exchangeEURtoRUB)
}
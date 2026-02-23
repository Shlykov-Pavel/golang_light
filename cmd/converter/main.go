package main

import "fmt"

func askUser ()(string, string, float64) {
	var currencyChange string
	var currencyTarget string
	var summ float64

	for {
		fmt.Print("Какую валюту вы хотите поменять (EUR/USD/RUB): ")
		fmt.Scan(&currencyChange)
		if currencyChange == "EUR" || currencyChange == "USD" || currencyChange == "RUB" {
			break
		}
	}

    for {
	    switch currencyChange {
	    case "EUR":
	        fmt.Print("На какую валюту вы хотите сконвертировать (USD/RUB): ")
		case "USD":
			fmt.Print("На какую валюту вы хотите сконвертировать (EUR/RUB): ")
		case "RUB":
			fmt.Print("На какую валюту вы хотите сконвертировать (EUR/USD): ")
		}
		fmt.Scan(&currencyTarget)
		if  currencyTarget=="EUR" && currencyTarget !=currencyChange || 
	   		currencyTarget=="USD" && currencyTarget !=currencyChange || 
	   		currencyTarget=="RUB" && currencyTarget !=currencyChange {
		break
		}
	}

	for {
		fmt.Print("Введите сумму для обмена: ")
		fmt.Scan(&summ)
		if summ>0 {
			break
		}
	}
return currencyChange, currencyTarget, summ
}

func countCurrency (currencyChange string, currencyTarget string, summ float64) float64 {
	const exchangeUSDtoEUR float64 = 0.84
	const exchangeUSDtoRUB float64 = 76.01
	const exchangeEURtoRUB float64 = exchangeUSDtoRUB / exchangeUSDtoEUR
    var result float64
	 
    switch {
	case currencyChange=="EUR" && currencyTarget=="USD":
		result=summ/exchangeUSDtoEUR
	case currencyChange=="EUR" && currencyTarget=="RUB":
		result=summ/exchangeUSDtoEUR*exchangeUSDtoRUB
	case currencyChange=="USD" && currencyTarget=="EUR":
		result=summ*exchangeUSDtoEUR
	case currencyChange=="USD" && currencyTarget=="RUB":
		result=summ*exchangeUSDtoRUB
	case currencyChange=="RUB" && currencyTarget=="EUR":
		result=summ/exchangeUSDtoRUB*exchangeUSDtoEUR
	case currencyChange=="RUB" && currencyTarget=="USD":
		result=summ/exchangeUSDtoRUB
	}
	return result
}

func main () {
	currencyChange, currencyTarget, summ:= askUser()
	result := countCurrency(currencyChange, currencyTarget, summ)

	fmt.Printf("Сумма сконвертированной валюты из %s в %s составит %.2f/n", currencyChange, currencyTarget, result)

}

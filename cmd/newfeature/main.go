package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var operation string
	fmt.Print("Укажите операцию (AVG, SUM, MED): ")
	fmt.Scan(&operation)
	operation = strings.ToUpper(strings.TrimSpace(operation))

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите числа через запятую (например: 20, 40, 40, 60): ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	parts := strings.Split(input, ",")
	numbers := make([]float64, 0, len(parts))

	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}

		n, err := strconv.ParseFloat(p, 64)
		if err != nil {
			fmt.Println("Ошибка: в списке есть не число ->", p)
			return
		}
		numbers = append(numbers, n)
	}

	if len(numbers) == 0 {
		fmt.Println("Ошибка: список чисел пустой.")
		return
	}

	var sum float64
	for _, n := range numbers {
		sum += n
	}

	switch operation {
	case "SUM":
		fmt.Printf("Результат SUM: %.2f\n", sum)
	case "AVG":
		fmt.Printf("Результат AVG: %.2f\n", sum/float64(len(numbers)))
	case "MED":
		sort.Float64s(numbers)
		mid := len(numbers) / 2
		var med float64
		if len(numbers)%2 == 0 {
			med = (numbers[mid-1] + numbers[mid]) / 2
		} else {
			med = numbers[mid]
		}
		fmt.Printf("Результат MED: %.2f\n", med)
	default:
		fmt.Println("Ошибка: операция должна быть SUM, AVG или MED.")
	}
}

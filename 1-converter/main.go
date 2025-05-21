package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	USDtoRUB = 92.5
	EURtoRUB = 100.2
	EURtoUSD = 0.92
)

func main() {
	fmt.Println("Конвертер валют")
	fmt.Println("Доступные валюты: EUR, USD, RUB")

	from := readCurrency("Введите исходную валюту:")
	sum := readSum("Введите сумму для конвертации:")
	to := readCurrency("Введите целевую валюту:")

	result := calculationMoney(sum, from, to)
	fmt.Printf("%.2f %s = %.2f %s\n", sum, from, result, to)
}

func readCurrency(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt + " ")
		input, _ := reader.ReadString('\n')
		curr := strings.ToUpper(strings.TrimSpace(input))
		if isValidCurrency(curr) {
			return curr
		}
		fmt.Println("Неверная валюта. Попробуйте ещё раз.")
	}
}

func readSum(prompt string) float64 {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt + " ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		value, err := strconv.ParseFloat(input, 64)
		if err == nil && value >= 0 {
			return value
		}
		fmt.Println("Неверная сумма. Введите положительное число.")
	}
}

func isValidCurrency(curr string) bool {
	switch curr {
	case "EUR", "USD", "RUB":
		return true
	default:
		return false
	}
}

func calculationMoney(sum float64, from, to string) float64 {
	switch {
	case from == "RUB" && to == "USD":
		return float64(sum) / USDtoRUB
	case from == "RUB" && to == "EUR":
		return float64(sum) / EURtoRUB
	case from == "USD" && to == "RUB":
		return float64(sum) * USDtoRUB
	case from == "USD" && to == "EUR":
		return float64(sum) * EURtoUSD
	case from == "EUR" && to == "RUB":
		return float64(sum) * EURtoRUB
	case from == "EUR" && to == "USD":
		return float64(sum) / EURtoUSD
	default:
		return -1
	}
}

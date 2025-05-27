package main

// проверка
import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Введите числа через пробел (можно с запятыми):")

	arr := readArr()
	operator := chekUserOperation("Введите операцию (AVG, SUM, MED):")
	result := calculateOperation(operator, arr)

	fmt.Printf("Результат: %.2f\n", result)
}

func calculateOperation(operator string, array []float64) float64 {
	var value float64

	switch operator {
	case "AVG":
		for _, v := range array {
			value += v
		}
		return value / float64(len(array))
	case "SUM":
		for _, v := range array {
			value += v
		}
		return value
	case "MED":
		sort.Float64s(array)
		len := len(array)
		if len%2 == 0 {
			value = array[len/2-1] + array[len/2]/2
		} else {
			value = array[len/2]
		}
		return value
	default:
		return 0
	}
}

func chekUserOperation(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt + " ")
		input, _ := reader.ReadString('\n')
		curr := strings.ToUpper(strings.TrimSpace(input))
		if isValidCurrency(curr) {
			return curr
		}
		fmt.Println("Пока нет такой операции, попробуйте ее ввести повторно.")
	}
}

func isValidCurrency(curr string) bool {
	switch curr {
	case "AVG", "SUM", "MED":
		return true
	default:
		return false
	}
}

func readArr() []float64 {

	var arr []float64

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	input = strings.TrimSpace(input)
	input = strings.ReplaceAll(input, ",", "")
	parts := strings.Fields(input)

	for _, part := range parts {
		num, _ := strconv.ParseFloat(part, 64)
		arr = append(arr, num)

	}

	return arr
}

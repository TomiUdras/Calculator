package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Calculator struct {
	romanNumerals  map[string]int
	arabicNumerals map[int]string
}

func NewCalculator() *Calculator {
	romanNumerals := map[string]int{"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10, "XI": 11, "XII": 12, "XIII": 13, "XIV": 14, "XV": 15, "XVI": 16, "XVII": 17, "XVIII": 18, "XIX": 19, "XX": 20, "XXI": 21}
	arabicNumerals := make(map[int]string)
	for key, value := range romanNumerals {
		arabicNumerals[value] = key
	}

	return &Calculator{romanNumerals: romanNumerals, arabicNumerals: arabicNumerals}
}

func (calc *Calculator) IsRoman(num string) bool {
	_, exists := calc.romanNumerals[num]
	return exists
}

func (calc *Calculator) IsArabic(num string) bool {
	arabic, err := strconv.Atoi(num)
	return err == nil && 1 <= arabic && arabic <= 10
}

func (calc *Calculator) ToArabic(num string) int {
	if calc.IsRoman(num) {
		return calc.romanNumerals[num]
	}
	arabic, _ := strconv.Atoi(num)
	return arabic
}

func (calc *Calculator) ToRoman(num int) string {
	return calc.arabicNumerals[num]
}

func (calc *Calculator) Calculate(num1, operator, num2 string) {
	arabicNum1 := calc.ToArabic(num1)
	arabicNum2 := calc.ToArabic(num2)

	var result int

	switch operator {
	case "+":
		result = arabicNum1 + arabicNum2
	case "-":
		result = arabicNum1 - arabicNum2
	case "*":
		result = arabicNum1 * arabicNum2
	case "/":
		if arabicNum2 == 0 {
			fmt.Println("Ошибка: деление на ноль")
			return
		}
		result = arabicNum1 / arabicNum2
	default:
		fmt.Println("Ошибка: неверная арифметическая операция")
		return
	}

	if calc.IsArabic(num1) && calc.IsArabic(num2) {
		fmt.Println(result)
	} else if calc.IsRoman(num1) && calc.IsRoman(num2) {
		if result <= 0 {
			fmt.Println("Ошибка: результат не может быть отрицательным или нулевым при использовании римских чисел")
		} else {
			fmt.Println(calc.ToRoman(result))
		}
	} else {
		fmt.Println("Ошибка: использование римских и арабских чисел в одной операции")
	}
}

func main() {
	calculator := NewCalculator()

	fmt.Print("Введите выражение: ")

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		inputStr := scanner.Text()

		parts := strings.Fields(inputStr)
		if len(parts) != 3 {
			fmt.Println("Ошибка: неверный формат ввода")
			return
		}

		num1, operator, num2 := parts[0], parts[1], parts[2]
		calculator.Calculate(num1, operator, num2)
	} else {
		fmt.Println("Ошибка при считывании ввода")
	}
}

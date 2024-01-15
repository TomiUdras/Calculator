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
  romanNumerals := map[string]int{"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10, "XI": 11, "XII": 12, "XIII": 13, "XIV": 14, "XV": 15, "XVI": 16, "XVII": 17, "XVIII": 18, "XIX": 19, "XX": 20, "XXI": 21, "XXII": 22, "XXIII": 23, "XXIV": 24, "XXV": 25, "XXVI": 26, "XXVII": 27, "XXVIII": 28, "XXIX": 29, "XXX": 30, "XXXI": 31, "XXXII": 32, "XXXIII": 33, "XXXIV": 34, "XXXV": 35, "XXXVI": 36, "XXXVII": 37, "XXXVIII": 38, "XXXIX": 39, "XL": 40, "XLI": 41, "XLII": 42, "XLIII": 43, "XLIV": 44, "XLV": 45, "XLVI":46,  "XLVII": 47, "XLVIII": 48, "XLIX": 49, "L": 50, "LI": 51, "LII": 52, "LIII": 53, "LIV": 54, "LV": 55, "LVI": 56, "LVII": 57, "LVIII": 58, "LIX": 59, "LX": 60, "LXI": 61, "LXII": 62, "LXIII": 63, "LXIV": 64, "LXV": 65, "LXVI": 66, "LXVII": 67, "LXVIII": 68, "LXIX": 69, "LXX": 70, "LXXI": 71, "LXXII": 72, "LXXIII": 73, "LXXIV": 74, "LXXV": 75, "LXXVI": 76, "LXXVII": 77, "LXXVIII": 78, "LXXIX": 79, "LXXX": 80, "LXXXI": 81, "LXXXII": 82, "LXXXIII": 83, "LXXXIV": 84, "LXXXV": 85, "LXXXVI": 86, "LXXXVII": 87, "LXXXVIII": 88, "LXXXIX": 89, "XC": 90, "XCI": 91, "XCII": 92, "XCIII": 93, "XCIV": 94, "XCV": 95, "XCVI": 96, "XCVII": 97, "XCVIII": 98, "XCIX": 99, "C": 100}
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
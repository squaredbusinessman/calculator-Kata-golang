package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Введите выражение (например, 3 + 5): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			fmt.Println("Программа завершена.")
			break
		}

		result := calculateInput(input)

		fmt.Println("Результат:", result)
		break
	}
}

// Функция проверки вводимых значений и
func calculateInput(input string) string {
	parts := strings.Fields(input)
	if len(parts) < 3 {
		panic("Выдача паники, так как строка не является математической операцией.")
	}
	if len(parts) > 3 {
		panic("Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
	}

	romanNumbers := map[string]int{
		"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
		"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
	}

	romanINput := false

	for key, value := range romanNumbers {
		if parts[0] == key {
			if _, ok := romanNumbers[parts[0]]; !ok {
				fmt.Println(parts[0])
				panic("Выдача паники, так как используются одновременно разные системы счисления.")
			}
			parts[0] = strconv.Itoa(value)
			romanINput = true
		}
		if parts[2] == key {
			if _, ok := romanNumbers[parts[2]]; !ok {
				panic("Выдача паники, так как используются одновременно разные системы счисления.")
			}
			parts[2] = strconv.Itoa(value)
			romanINput = true
		}
	}

	operand1, err := strconv.Atoi(parts[0])
	if operand1 < 1 || operand1 > 10 {
		return "Неверное число, программа работает только с натуральными числами от 0 до 10"
	}
	if err != nil {
		fmt.Errorf("некорректное число: %s", parts[0])
	}

	operator := parts[1]

	operand2, err := strconv.Atoi(parts[2])
	if operand2 < 1 || operand2 > 10 {
		return "Неверное число, программа работает только с натуральными числами от 0 до 10"
	}
	if err != nil {
		fmt.Errorf("некорректное число: %s", parts[2])
	}

	if !isValidInput(operand1, operator, operand2) {
		fmt.Errorf("некорректные данные")
	}

	resultCalculate, err := calculate(operand1, operator, operand2, romanINput)

	return resultCalculate
}

// Функция для преобразования арабского числа в римское
func arabicToRoman(num int) string {
	// Определение римских цифр и их значения
	romanNumerals := map[int]string{
		1: "I", 2: "II", 3: "III", 4: "IV", 5: "V",
		6: "VI", 7: "VII", 8: "VIII", 9: "IX", 10: "X",
		11: "XI", 12: "XII", 13: "XIII", 14: "XIV", 15: "XV",
		16: "XVI", 17: "XVII", 18: "XVIII", 19: "XIX", 20: "XX",
		21: "XXI", 22: "XXII", 23: "XXIII", 24: "XXIV", 25: "XXV",
		26: "XXVI", 27: "XXVII", 28: "XXVIII", 29: "XXIX", 30: "XXX",
		31: "XXXI", 32: "XXXII", 33: "XXXIII", 34: "XXXIV", 35: "XXXV",
		36: "XXXVI", 37: "XXXVII", 38: "XXXVIII", 39: "XXXIX", 40: "XL",
		41: "XLI", 42: "XLII", 43: "XLIII", 44: " XLIV", 45: "XLV",
		46: "XLVI", 47: "XLVII", 48: "XLVIII", 49: "XLIX", 50: "L",
		51: "LI", 52: "LII", 53: "LIII", 54: "LIV", 55: "LV",
		56: "LVI", 57: "LVII", 58: "LVIII", 59: "LIX", 60: "LX",
		61: "LXI", 62: "LXII", 63: "LXIII", 64: "LXIV", 65: "LXV",
		66: "LXVI", 67: "LXVII", 68: "LXVIII", 69: "LXIX", 70: "LXX",
		71: "LXXI", 72: "LXXII", 73: "LXXIII", 74: "LXXIV", 75: "LXXV",
		76: "LXXVI", 77: "LXXVII", 78: "LXXVIII", 79: "LXXIX", 80: "LXXX",
		81: "LXXXI", 82: "LXXXII", 83: "LXXXIII", 84: "LXXXIV", 85: "LXXXV",
		86: "LXXXVI", 87: "LXXXVII", 88: "LXXXVIII", 89: "LXXXIX", 90: "XC",
		91: "XCI", 92: "XCII", 93: "XCIII", 94: "XCIV", 95: "XCV",
		96: "XCVI", 97: "XCVII", 98: "XCVIII", 99: "XCIX", 100: "C",
	}

	// Создаем строку для хранения римского числа
	result := ""

	// Проходим по римским цифрам
	for value, numeral := range romanNumerals {
		if num == value {
			result = numeral
		}
	}

	return result
}

// Функция проверки валидности оператора
func isValidInput(operand1 int, operator string, operand2 int) bool {
	if operand1 < 1 || operand1 > 10 {
		panic("Выдача паники, так как разрешены только операнды от 1 до 10 включительно")
	}

	if operand2 < 1 || operand2 > 10 {
		panic("Выдача паники, так как разрешены только операнды от 1 до 10 включительно")
	}

	operators := map[string]bool{"+": true, "-": true, "*": true, "/": true}

	_, operatorIsValid := operators[operator]

	return operatorIsValid
}

// Функция калькуляции операндов с условленным оператором, фиксирующая операнды и результат в формате римских чисел.
func calculate(num1 int, operator string, num2 int, isRoman bool) (string, error) {
	var result int
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			return "", fmt.Errorf("деление на ноль")
		}
		result = num1 / num2
	default:
		return "", fmt.Errorf("недопустимый оператор")
	}

	if isRoman {
		if result < 1 {
			panic("Выдача паники, так как в римской системе нет отрицательных чисел и числа 0.")
		} else {
			return arabicToRoman(result), nil
		}
	}

	return strconv.Itoa(result), nil
}

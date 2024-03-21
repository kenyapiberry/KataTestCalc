package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Введите математическую операцию через пробел: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	words := strings.Fields(input)

	if len(words) < 3 {
		panic("Строка не является математической операцией")
	} else if len(words) > 3 {
		panic("Формат математической операции не удовлетворияет условию")
	}

	num1, err1 := strconv.Atoi(words[0])
	num2, err2 := strconv.Atoi(words[2])

	if err1 == nil && err2 == nil {
		if isInRange(num1) && isInRange(num2) {
			fmt.Println(calculate(num1, num2, words[1]))
		} else {
			panic("Числа должны быть от 1 до 10")
		}
	} else if err1 != nil && err2 != nil {
		numRoman1 := convToArabic(words[0])
		numRoman2 := convToArabic(words[2])
		result := calculate(numRoman1, numRoman2, words[1])
		if result > 0 {
			fmt.Println(arabicToRoman(result))
		} else {
			panic("В римской системе нет отрицательных чисел")
		}
	} else {
		panic("Калькулятор умеет работать только с арабскими или римскими цифрами одновременно")
	}

}

func convToArabic(roman string) int {
	switch roman {
	case "I":
		return 1
	case "II":
		return 2
	case "III":
		return 3
	case "IV":
		return 4
	case "V":
		return 5
	case "VI":
		return 6
	case "VII":
		return 7
	case "VIII":
		return 8
	case "IX":
		return 9
	case "X":
		return 10
	case "XL":
		return 40
	case "L":
		return 50
	case "C":
		return 100

	default:
		panic("Числа должны быть от I до X")
	}
}

func arabicToRoman(arabic int) string {
	roman := ""
	values := []int{50, 10, 9, 5, 4, 1}
	symbols := []string{"L", "X", "IX", "V", "IV", "I"}

	for i := 0; i < 6; i++ {
		for arabic >= values[i] {
			arabic -= values[i]
			roman += symbols[i]
		}
	}

	return roman
}

func calculate(num1, num2 int, sign string) int {
	switch sign {
	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	case "*":
		return num1 * num2
	case "/":
		if num2 == 0 {
			panic("Деление на ноль невозможно")
		}
		return num1 / num2
	default:
		panic("Ошибка")
	}
}

func isInRange(num int) bool {
	return num > 0 && num <= 10
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var romanToArabic = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
	"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
}

var arabicToRoman = []string{
	"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X",
	"XI", "XII", "XIII", "XIV", "XV", "XVI", "XVII", "XVIII", "XIX", "XX",
	"XXI", "XXII", "XXIII", "XXIV", "XXV", "XXVI", "XXVII", "XXVIII", "XXIX", "XXX",
	"XXXI", "XXXII", "XXXIII", "XXXIV", "XXXV", "XXXVI", "XXXVII", "XXXVIII", "XXXIX", "XL",
	"XLI", "XLII", "XLIII", "XLIV", "XLV", "XLVI", "XLVII", "XLVIII", "XLIX", "L",
	"LI", "LII", "LIII", "LIV", "LV", "LVI", "LVII", "LVIII", "LIX", "LX",
	"LXI", "LXII", "LXIII", "LXIV", "LXV", "LXVI", "LXVII", "LXVIII", "LXIX", "LXX",
	"LXXI", "LXXII", "LXXIII", "LXXIV", "LXXV", "LXXVI", "LXXVII", "LXXVIII", "LXXIX", "LXXX",
	"LXXXI", "LXXXII", "LXXXIII", "LXXXIV", "LXXXV", "LXXXVI", "LXXXVII", "LXXXVIII", "LXXXIX", "XC",
	"XCI", "XCII", "XCIII", "XCIV", "XCV", "XCVI", "XCVII", "XCVIII", "XCIX", "C",
}

func isArabic(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func isRoman(s string) bool {
	_, exists := romanToArabic[s]
	return exists
}

func toArabic(s string) int {
	if val, err := strconv.Atoi(s); err == nil {
		return val
	}
	return romanToArabic[s]
}

func toRoman(n int) string {
	if n <= 0 || n >= len(arabicToRoman) {
		panic("Result out of range for Roman numerals")
	}
	return arabicToRoman[n]
}

func evaluate(a, b int, operator string) int {
	switch operator {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		if b == 0 {
			panic("Division by zero")
		}
		return a / b
	default:
		panic("Invalid operator")
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Simple Calculator")
	fmt.Println("Enter expressions in the format: operand1 operator operand2")
	fmt.Println("Example: 3 + 4 or IV * II")

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}

		expression := scanner.Text()
		if strings.ToLower(expression) == "exit" {
			break
		}

		tokens := strings.Fields(expression)
		if len(tokens) != 3 {
			panic("Invalid input format")
		}

		operand1, operator, operand2 := tokens[0], tokens[1], tokens[2]

		if isArabic(operand1) && isArabic(operand2) {
			a := toArabic(operand1)
			b := toArabic(operand2)
			if a < 1 || a > 10 || b < 1 || b > 10 {
				panic("Arabic numbers out of range (1-10)")
			}
			result := evaluate(a, b, operator)
			fmt.Printf("Result: %d\n", result)
		} else if isRoman(operand1) && isRoman(operand2) {
			a := toArabic(operand1)
			b := toArabic(operand2)
			if a < 1 || a > 10 || b < 1 || b > 10 {
				panic("Roman numbers out of range (I-X)")
			}
			result := evaluate(a, b, operator)
			if result <= 0 {
				panic("Result out of range for Roman numerals")
			}
			fmt.Printf("Result: %s\n", toRoman(result))
		} else {
			panic("Mixed or invalid numeral systems")
		}
	}

	if scanner.Err() != nil {
		fmt.Println("Error reading input:", scanner.Err())
	}
}

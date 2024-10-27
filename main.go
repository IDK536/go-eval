package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type addition struct {
	numbers []float64
	znak    []string
}

type staples struct {
	staples []string
	znak    []string
}

func parseExpression(expression string) (addition, error) {
	a := addition{}
	znakinx := 0

	// Добавляем дополнительный пробел, чтобы захватить последнее число.

	// fmt.Println(expression)
	for i := range expression {
		if strings.Contains("+-*/", string(expression[i])) {
			intnumber, err := strconv.ParseFloat(strings.TrimSpace(expression[znakinx:i]), 64)
			if err != nil {
				return a, err
			}
			a.numbers = append(a.numbers, intnumber)
			a.znak = append(a.znak, string(expression[i]))
			znakinx = i + 1
		}
	}
	// Обрабатываем последнее число
	if znakinx < len(expression) {
		intnumber, err := strconv.ParseFloat(strings.TrimSpace(expression[znakinx:]), 64)
		if err != nil {
			return a, err
		}
		a.numbers = append(a.numbers, intnumber)
	}
	return a, nil
}

func calculateMulDiv(a *addition) {
	for i := 0; i < len(a.znak); i++ {
		if a.znak[i] == "*" {
			a.numbers[i] = a.numbers[i] * a.numbers[i+1]
			a.numbers = append(a.numbers[:i+1], a.numbers[i+2:]...)
			a.znak = append(a.znak[:i], a.znak[i+1:]...)
			i-- // Уменьшаем индекс, чтобы повторно проверить текущую позицию
		} else if a.znak[i] == "/" {
			a.numbers[i] = a.numbers[i] / a.numbers[i+1]
			a.numbers = append(a.numbers[:i+1], a.numbers[i+2:]...)
			a.znak = append(a.znak[:i], a.znak[i+1:]...)
			i-- // Уменьшаем индекс, чтобы повторно проверить текущую позицию
		}
	}
}

func calculateAddSub(a *addition) float64 {
	calc := a.numbers[0]
	for i := 0; i < len(a.znak); i++ {
		if a.znak[i] == "+" {
			calc += a.numbers[i+1]
		} else if a.znak[i] == "-" {
			calc -= a.numbers[i+1]
		}
	}
	return calc
}

func Calc_without_brackets(expression string) (float64, error) {
	a, err := parseExpression(expression)
	if err != nil {
		return 0, err
	}

	calculateMulDiv(&a) // Вычисляем умножение и деление
	// fmt.Println(a)
	result := calculateAddSub(&a) // Вычисляем сложение и вычитание
	return result, nil
}

// func Staples(expression string) staples {
// 	esh := 0
// 	s := staples{}
// 	bufindex := 0
// 	for i := range expression {
// 		if expression[i] == '(' && esh != 1 {
// 			esh = 0
// 			// fmt.Println(bufindex, i-1, len(expression))
// 			s.staples = append(s.staples, expression[bufindex:i-1])
// 			s.znak = append(s.znak, string(expression[i-1]))
// 			if bufindex != 0 {
// 				if string(expression[bufindex-1]) == "-" {
// 					s.staples[len(s.staples)-1] = strings.Replace(s.staples[len(s.staples)-1], "+", "?", -1)
// 					s.staples[len(s.staples)-1] = strings.Replace(s.staples[len(s.staples)-1], "-", "+", -1)
// 					s.staples[len(s.staples)-1] = strings.Replace(s.staples[len(s.staples)-1], "?", "-", -1)
// 				}
// 			}
// 			bufschet := 0
// 			for j := range expression[i:] {
// 				if expression[j] == '(' {
// 					bufschet++
// 				}
// 				if expression[j] == ')' {
// 					// fmt.Println(expression[j])
// 					bufschet--
// 					if bufschet == 0 {

// 						s.staples = append(s.staples, expression[i+1:j])
// 						s.znak = append(s.znak, string(expression[j+1]))
// 						// fmt.Println(s)
// 						bufindex = j
// 						// fmt.Println(i+1, j, j+1)
// 						esh = 1
// 						break
// 					}
// 				}
// 			}
// 		}

// 	}
// 	// fmt.Println(1, 1, 2)
// 	if bufindex < len(expression) {
// 		// fmt.Println(s)
// 		s.staples = append(s.staples, expression[bufindex+2:])
// if string(expression[bufindex+1]) == "-" {
// 	s.staples[len(s.staples)-1] = strings.Replace(s.staples[len(s.staples)-1], "+", "?", -1)
// 	s.staples[len(s.staples)-1] = strings.Replace(s.staples[len(s.staples)-1], "-", "+", -1)
// 	s.staples[len(s.staples)-1] = strings.Replace(s.staples[len(s.staples)-1], "?", "-", -1)
// }

// 	}
// 	return s
// }

func splitByParentheses(s string) ([]string, error) {
	var result []string
	var current string
	var stack []rune // Стек для отслеживания уровней скобок

	for _, char := range s {
		if char == '(' {
			// Если текущая строка не пустая, добавляем её в результат
			if current != "" {
				result = append(result, current) // Добавляем часть до скобок
				// fmt.Println(char)
				current = ""
			}
			stack = append(stack, char) // Добавляем открывающую скобку в стек
		} else if char == ')' {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1] // Убираем последнюю открытую скобку из стека
				if len(stack) == 0 {
					if current != "" {
						result = append(result, current) // Добавляем полное выражение между скобками в результат
						current = ""                     // Обнуляем текущую строку
					}
				}
			} else {
				// Обработка случая, когда закрывающая скобка не находит соответствующей открывающей
				return []string{}, errors.New("Staples")
			}
		} else {
			// Добавляем обычный символ в текущую строку
			current += string(char)
		}
	}

	// Добавляем оставшуюся часть
	if current != "" {
		result = append(result, current)
	}
	// fmt.Println(result)
	return result, nil
}

func hasConsecutiveOperators(s string) bool {
	operators := "+-*/"

	for i := 0; i < len(s)-1; i++ {
		// Проверяем, что текущий символ и следующий являются операторами
		if strings.ContainsRune(operators, rune(s[i])) && strings.ContainsRune(operators, rune(s[i+1])) {
			return true
		}
	}
	return false
}

func Calc(expression string) (float64, error) {
	expression = strings.Replace(expression, " ", "", -1)
	if len(expression) == 0 {
		return 0, errors.New("empty expression")
	}
	if hasConsecutiveOperators(expression) {
		return 0, errors.New("nenene")
	}
	s := staples{}
	buf, err1 := splitByParentheses(expression)
	if err1 != nil {
		return 0, err1
	}
	if string(expression[0]) == "+" || string(expression[0]) == "-" || string(expression[0]) == "/" || string(expression[0]) == "*" || string(expression[len(expression)-1]) == "+" || string(expression[len(expression)-1]) == "-" || string(expression[len(expression)-1]) == "*" || string(expression[len(expression)-1]) == "/" {
		return 0, errors.ErrUnsupported
	}
	for part := range buf {
		if string(buf[part][0]) == "+" || string(buf[part][0]) == "-" || string(buf[part][0]) == "*" || string(buf[part][0]) == "/" {
			s.staples = append(s.staples, buf[part][1:])
			s.znak = append(s.znak, string(buf[part][0]))
			// fmt.Println(s.staples)
			if string(buf[part][0]) == "-" {
				s.staples[len(s.staples)-1] = strings.Replace(s.staples[len(s.staples)-1], "+", "?", -1)
				s.staples[len(s.staples)-1] = strings.Replace(s.staples[len(s.staples)-1], "-", "+", -1)
				s.staples[len(s.staples)-1] = strings.Replace(s.staples[len(s.staples)-1], "?", "-", -1)
			}
		} else if string(buf[part][len(string(buf[part]))-1]) == "+" || string(buf[part][len(string(buf[part]))-1]) == "-" || string(buf[part][len(string(buf[part]))-1]) == "*" || string(buf[part][len(string(buf[part]))-1]) == "/" {
			s.staples = append(s.staples, buf[part][:len(string(buf[part]))-1])
			s.znak = append(s.znak, string(buf[part][len(string(buf[part]))-1]))
			if string(buf[part][len(string(buf[part]))-1]) == "-" {
				s.staples[len(s.staples)-1] = strings.Replace(s.staples[len(s.staples)-1], "+", "?", -1)
				s.staples[len(s.staples)-1] = strings.Replace(s.staples[len(s.staples)-1], "-", "+", -1)
				s.staples[len(s.staples)-1] = strings.Replace(s.staples[len(s.staples)-1], "?", "-", -1)
			}
		} else {
			s.staples = append(s.staples, buf[part])
		}
	}
	// fmt.Println(s)
	// fmt.Println(s)
	// for i := range s.staples {
	// 	fmt.Println(s.staples[i])
	// 	if strings.Contains(string(s.staples[i]), "(") {
	// 		bufs1 := Staples(s.staples[i])
	// 		fmt.Println(bufs1)
	// 		bufss := append(s.staples[:i], bufs1.staples...)
	// 		s.staples = append(bufss, s.staples[i+1:]...)
	// 		fmt.Println(s, i)
	// 	}
	// }

	// fmt.Println(s)
	a := addition{}
	a.znak = s.znak
	for i := range s.staples {
		// fmt.Println(s.staples[i])
		bufs, _ := Calc_without_brackets(s.staples[i])
		a.numbers = append(a.numbers, bufs)
		// fmt.Println(a.numbers)
	}
	// fmt.Println(s, a)
	calculateMulDiv(&a)
	result1 := calculateAddSub(&a)
	if result1 == math.Inf(1) || result1 == math.Inf(-1) {
		return 0, errors.New("mnogo")
	}
	return result1, nil
}

func main() {
	result, err := Calc("1/1/0")
	fmt.Println(result)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Result:", result)
}

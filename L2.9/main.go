package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
)

var errInvalidString = errors.New("invalid string")

func parse(data string) (string, error) {
	// Обрабатываем кейс с пустой строкой
	if data == "" {
		return "", nil
	}

	runes := []rune(data)

	// Обрабатываем кейс, если строка начинается с числа
	if unicode.IsDigit(runes[0]) {
		return "", errInvalidString
	}

	var builder strings.Builder
	l := len(runes)

	var prev rune
	escaped := false

	for i := 0; i < l; i++ {
		ch := runes[i]

		// Если символ экранирован символом "\"
		if escaped {
			escaped = false

			// Если после экранированной цифры идет 0, удаляем ее
			if i+1 < l {
				if runes[i+1] == '0' {
					i++
					continue
				}
			}
			prev = ch
			builder.WriteRune(ch)
			continue
		}

		// Сигнализируем, если цифра экранирована
		if ch == '\\' {
			if i == l-1 {
				return "", errInvalidString
			}
			escaped = true
			continue
		}

		// Строим последовательность определенной длины
		if unicode.IsDigit(ch) && ch != '0' {
			j := i
			for j < l && unicode.IsDigit(runes[j]) {
				j++
			}

			num, err := strconv.Atoi(string(runes[i:j]))
			if err != nil {
				return "", err
			}
			fmt.Println(num)
			builder.WriteString(strings.Repeat(string(prev), num-1))

			i = j - 1
			continue
		}

		// Если после символа идет 0, удаляем его
		if i+1 < l {
			if runes[i+1] == '0' {
				i++
				continue
			}
		}

		builder.WriteRune(ch)
		prev = ch
	}

	return builder.String(), nil
}

func main() {
	result, err := parse("\\ra\\5")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}

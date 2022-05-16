package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	var builder strings.Builder
	var flagEscape bool
	for i, char := range input {
		switch {
		case string(char) == `\` && !flagEscape:
			flagEscape = true
		case flagEscape:
			builder.WriteRune(char)
			flagEscape = false
		case unicode.IsDigit(char) && len(builder.String()) > 0:
			number, _ := strconv.Atoi(string(char))

			if number == 0 {
				if !unicode.IsDigit(rune(input[i-1])) {
					str := []rune(builder.String())
					str = str[:len(str)-1]
					builder.Reset()
					builder.WriteString(string(str))
				} else {
					return "", ErrInvalidString
				}
			} else {
				w := []rune(builder.String())
				w = w[len(w)-1:]
				builder.WriteString(strings.Repeat(string(w), number-1))
			}
		case !unicode.IsDigit(char):
			builder.WriteRune(char)
		default:
			return "", ErrInvalidString
		}
	}
	return builder.String(), nil
}

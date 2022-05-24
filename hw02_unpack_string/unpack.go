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
	var runeInput = []rune(input)
	var stringLength = len(runeInput)
	for i, char := range runeInput {
		switch {
		case string(char) == `\` && !flagEscape:
			flagEscape = true
		case flagEscape:
			builder.WriteRune(char)
			flagEscape = false
		case unicode.IsDigit(char) && builder.Len() > 0:
			number, _ := strconv.Atoi(string(char))

			if number == 0 {
				if unicode.IsDigit(runeInput[i-1]) {
					return "", ErrInvalidString
				} else {
					continue
				}

			} else {
				w := runeInput[i-1]
				builder.WriteString(strings.Repeat(string(w), number-1))
			}
		case !unicode.IsDigit(char):
			if i+1 <= stringLength {
				w := string(char)
				if i+1 < stringLength {
					w = string(runeInput[i+1])
				}
				if w != "0" {
					builder.WriteRune(char)
				}
			}
		default:
			return "", ErrInvalidString
		}
	}
	return builder.String(), nil
}

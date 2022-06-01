package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

type Pair struct {
	word  string
	count int
}

var reg = regexp.MustCompile(`[^\r\n\t\f\v\s]+`)

func Top10(input string) []string {
	if len(input) == 0 {
		return nil
	}
	words := make(map[string]int)

	res := reg.FindAllString(input, -1)
	for i := range res {
		word := strings.ToLower(res[i])
		word = strings.Trim(word, ".,!?()'`»«:;")
		if word != "-" {
			words[word]++
		}
	}

	top := make([]Pair, 0, len(words))
	for w, c := range words {
		top = append(top, Pair{w, c})
	}

	sort.Slice(top, func(i int, j int) bool {
		if top[i].count == top[j].count {
			return top[i].word < top[j].word
		}
		return top[i].count > top[j].count
	})

	lenList := len(top)
	if lenList > 10 {
		lenList = 10
	}

	topSlice := make([]string, 0, lenList)
	for i, v := range top {
		topSlice = append(topSlice, v.word)
		if i == 9 || i > len(top) {
			break
		}
	}

	return topSlice
}

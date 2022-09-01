package hw10programoptimization

import (
	"bufio"
	"errors"
	"io"
	"regexp"
	"strings"

	json "github.com/json-iterator/go"
)

type User struct {
	Email string
}

type DomainStat map[string]int

var (
	reg                  = regexp.MustCompile(`@(?P<emailPart>.+)`)
	ErrJSONUnmarshalling = errors.New("error json unmarshalling")
)

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	result := make(DomainStat)
	scanner := bufio.NewScanner(r)
	var user User
	for scanner.Scan() {
		if err := json.Unmarshal(scanner.Bytes(), &user); err != nil {
			return nil, err
		}

		if strings.HasSuffix(user.Email, domain) {
			matches := reg.FindStringSubmatch(strings.ToLower(user.Email))
			if len(matches) > 0 {
				partIndex := reg.SubexpIndex("emailPart")
				result[matches[partIndex]]++
			}
		}
	}

	return result, nil
}

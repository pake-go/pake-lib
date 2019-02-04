package argutil

import (
	"encoding/csv"
	"strings"
)

func GetTokens(str string) ([]string, error) {
	r := csv.NewReader(strings.NewReader(str))
	r.Comma = ' '
	return r.Read()
}

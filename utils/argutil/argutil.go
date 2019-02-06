// Package argutil provides functions for working with shell arguments
package argutil

import (
	"encoding/csv"
	"strings"
)

// GetTokens takes a string, parses it for tokens as evaluated by a shell, and
// return a list of tokens along with any errors encountered.
func GetTokens(str string) ([]string, error) {
	r := csv.NewReader(strings.NewReader(str))
	r.Comma = ' '
	return r.Read()
}

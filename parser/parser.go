package parser

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"strings"

	pakelib "github.com/pake-go/pake-lib"
)

type parser struct {
	commandCandidates []pakelib.CommandCandidate
	commentValidator  pakelib.CommentValidator
}

func New(cmdCandidates []pakelib.CommandCandidate, commentValidator pakelib.CommentValidator) *parser {
	return &parser{
		commandCandidates: cmdCandidates,
		commentValidator:  commentValidator,
	}
}

func (p *parser) ParseFile(filename string) ([]pakelib.Command, error) {
	fileContent, err := ioutil.ReadFile(filename)
	if err != nil {
		return []pakelib.Command{}, err
	}
	return p.ParseString(string(fileContent))
}

func (p *parser) ParseString(str string) ([]pakelib.Command, error) {
	var commands []pakelib.Command
	lines := strings.Split(str, "\n")
	for _, line := range lines {
		command, err := p.ParseLine(line)
		if err != nil {
			return []pakelib.Command{}, err
		}
		commands = append(commands, command)
	}
	return commands, nil
}

func (p *parser) ParseLine(line string) (pakelib.Command, error) {
	if p.commentValidator.IsValid(line) {
		return &pakelib.Comment{}, nil
	}
	for _, cmdCandidate := range p.commandCandidates {
		validator := cmdCandidate.Validator
		if validator.CanHandle(line) {
			tokens, err := GetTokens(line)
			if err != nil {
				return nil, err
			}
			args := tokens[1:len(tokens)]

			if validator.ValidateArgs(args) {
				constructor := cmdCandidate.Constructor
				return constructor(args), nil
			}
			return nil, fmt.Errorf("At least one of the given argument is not valid")
		}
	}
	return nil, fmt.Errorf("%s is not valid syntax", line)
}

func GetTokens(str string) ([]string, error) {
	r := csv.NewReader(strings.NewReader(str))
	r.Comma = ' '
	return r.Read()
}

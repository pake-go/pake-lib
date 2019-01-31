package parser

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type parser struct {
	commandCandidates []CommandCandidate
	commentValidator  CommentValidator
}

func New(cmdCandidates []CommandCandidate, commentValidator CommentValidator) *parser {
	return &parser{
		commandCandidates: cmdCandidates,
		commentValidator:  commentValidator,
	}
}

func (p *parser) ParseFile(filename string) ([]Command, error) {
	fileContent, err := ioutil.ReadFile(filename)
	if err != nil {
		return []Command{}, err
	}
	return p.ParseString(string(fileContent))
}

func (p *parser) ParseString(str string) ([]Command, error) {
	var commands []Command
	lines := strings.Split(str, "\n")
	for _, line := range lines {
		command, err := p.ParseLine(line)
		if err != nil {
			return []Command{}, err
		}
		commands = append(commands, command)
	}
	return commands, nil
}

func (p *parser) ParseLine(line string) (Command, error) {
	if p.commentValidator.IsValid(line) {
		return &Comment{}, nil
	}
	for _, cmdCandidate := range p.commandCandidates {
		validator := cmdCandidate.validator
		if validator.CanHandle(line) {
			tokens := strings.Split(line, " ")
			args := tokens[1:len(tokens)]
			if validator.ValidateArgs(args) {
				constructor := cmdCandidate.constructor
				return constructor(args), nil
			}
			return nil, fmt.Errorf("At least one of the given argument is not valid")
		}
	}
	return nil, fmt.Errorf("%s is not valid syntax", line)
}

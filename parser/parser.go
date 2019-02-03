package parser

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	pakelib "github.com/pake-go/pake-lib"
)

type parser struct {
	commandCandidates []pakelib.CommandCandidate
	commentValidator  pakelib.CommentValidator
}

func New(cmdCandidates []pakelib.CommandCandidate, cv pakelib.CommentValidator) *parser {
	return &parser{
		commandCandidates: cmdCandidates,
		commentValidator:  cv,
	}
}

func (p *parser) ParseFile(filename string, logger *log.Logger) ([]pakelib.Command, error) {
	fileContent, err := ioutil.ReadFile(filename)
	if err != nil {
		logger.Println(err.Error())
		return []pakelib.Command{}, err
	}
	return p.ParseString(string(fileContent), logger)
}

func (p *parser) ParseString(str string, logger *log.Logger) ([]pakelib.Command, error) {
	var commands []pakelib.Command
	silentLogger := log.New(ioutil.Discard, "", log.LstdFlags)
	lines := strings.Split(str, "\n")
	for linenum, line := range lines {
		command, err := p.ParseLine(line, silentLogger)
		if err != nil {
			errMsg := fmt.Errorf("An error occured on line %d: %s", linenum+1, err.Error())
			logger.Println(errMsg.Error())
			return []pakelib.Command{}, errMsg
		}
		commands = append(commands, command)
	}
	return commands, nil
}

func (p *parser) ParseLine(line string, logger *log.Logger) (pakelib.Command, error) {
	if p.commentValidator.IsValid(line) {
		return &pakelib.Comment{}, nil
	}

	tokens, err := GetTokens(line)
	if err != nil {
		logger.Println(err.Error())
		return nil, err
	}
	args := tokens[1:len(tokens)]
	for _, cmdCandidate := range p.commandCandidates {
		validator := cmdCandidate.Validator
		if validator.CanHandle(line) {
			if validator.ValidateArgs(args) {
				constructor := cmdCandidate.Constructor
				return constructor(args), nil
			}
			errMsg := fmt.Errorf("At least one of the arguments is not valid: %+q", args)
			logger.Println(errMsg.Error())
			return nil, errMsg
		}
	}
	errMsg := fmt.Errorf("%s is not a valid command", tokens[0])
	logger.Println(errMsg.Error())
	return nil, errMsg
}

func GetTokens(str string) ([]string, error) {
	r := csv.NewReader(strings.NewReader(str))
	r.Comma = ' '
	return r.Read()
}

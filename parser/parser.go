// Package parser provides functions for parsing strings and source files and converting it
// into a list of commands that can be executed by executor.Run.
package parser

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	pakelib "github.com/pake-go/pake-lib"
	"github.com/pake-go/pake-lib/utils/argutil"
)

// Parser implements a parser for converting strings and source files into a list of commands.
type Parser struct {
	// Represents a list of CommandCandidate, a bundle of the command's validator and the
	// command's constructor.
	commandCandidates []pakelib.CommandCandidate
	// Represents the function used to check if a string is a valid comment.
	commentValidator pakelib.CommentValidator
}

// New returns a parser for converting source files and strings into a list of commands.
func New(cmdCandidates []pakelib.CommandCandidate, cv pakelib.CommentValidator) *Parser {
	return &Parser{
		commandCandidates: cmdCandidates,
		commentValidator:  cv,
	}
}

// ParseFile takes in a filename and parses the content of the file to return a list of commands
// that can be run by executor.Run along with any errors that were encountered.
func (p *Parser) ParseFile(filename string, logger *log.Logger) ([]pakelib.Command, error) {
	fileContent, err := ioutil.ReadFile(filename)
	if err != nil {
		logger.Println(err.Error())
		return []pakelib.Command{}, err
	}
	return p.ParseString(string(fileContent), logger)
}

// ParseString takes in a string and parses it to return a list of commands that can be run by
// executor.Run along with any errors that was encountered.
func (p *Parser) ParseString(str string, logger *log.Logger) ([]pakelib.Command, error) {
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

// ParseLine takes a string that represent one line of code in the language and parses it to
// return a list of commands that can be run by executor.Run along with any errors that were
// encountered.
func (p *Parser) ParseLine(line string, logger *log.Logger) (pakelib.Command, error) {
	if p.commentValidator.IsValid(line) {
		return &pakelib.Comment{}, nil
	}

	tokens, err := argutil.GetTokens(line)
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

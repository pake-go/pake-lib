package pakelib

import "github.com/pake-go/pake-lib/config"

type CommandCandidate struct {
	Validator   CommandValidator
	Constructor func([]string) Command
}

type Command interface {
	Execute(*config.Config) error
}

type CommandValidator interface {
	CanHandle(string) bool
	ValidateArgs([]string) bool
}

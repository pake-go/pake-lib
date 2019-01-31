package pakelib

import "github.com/pake-go/pake-lib/config"

type CommandCandidate struct {
	validator   CommandValidator
	constructor func([]string) Command
}

type Command interface {
	Execute(*config.Config) (bool, error)
}

type CommandValidator interface {
	CanHandle(string) bool
	ValidateArgs([]string) bool
}

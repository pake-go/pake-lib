package pakelib

type CommandCandidate struct {
	validator   CommandValidator
	constructor func([]string) Command
}

type Command interface {
	Execute() (bool, error)
}

type CommandValidator interface {
	IsValid(string) bool
	ValidateArgs([]string) bool
}

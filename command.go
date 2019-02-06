// Package pakelib provides the framework for writing a simple language
// whose source code consists solely of commands followed by their
// respective arguments.
package pakelib

import (
	"log"

	"github.com/pake-go/pake-lib/config"
)

// CommandCandidate is used to bundle a command's validator with its
// constructor.
type CommandCandidate struct {
	// Validator is a struct that must satisfy the ComamnValdidator interface.
	Validator CommandValidator
	// Constructor will construct the command and return it.
	Constructor func([]string) Command
}

// Command is an interface that all commands of the language must satisfy.
type Command interface {
	// Execute would perform the action behind the command.
	Execute(*config.Config, *log.Logger) error
}

// CommandValidator is an interface that all commands' validators must satisfy.
type CommandValidator interface {
	// CanHandle checks to see if the command can handle the given string.
	CanHandle(string) bool
	// ValidateArgs checks to see if the list of strings given are valid
	// arguments for the command.
	ValidateArgs([]string) bool
}

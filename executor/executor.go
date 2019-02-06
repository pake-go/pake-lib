// Package executor provides a default function for executing commands returned by the parser.
package executor

import (
	"fmt"
	"log"

	"github.com/PGo-Projects/output"
	pakelib "github.com/pake-go/pake-lib"
	"github.com/pake-go/pake-lib/config"
)

// Run iterates through the list of commands passed to it and calls the Execute() function for
// each of them.
func Run(commands []pakelib.Command, logger *log.Logger) {
	cfg := config.New()
	for line, command := range commands {
		err := command.Execute(cfg, logger)
		if err != nil {
			errMsg := fmt.Errorf("There was an error at line %d: %s", line+1, err.Error())
			logger.Println(errMsg.Error())
			output.Error(errMsg)
		}
		cfg.SmartReset()
	}
}

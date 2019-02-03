package executor

import (
	"fmt"
	"log"

	"github.com/PGo-Projects/output"
	pakelib "github.com/pake-go/pake-lib"
	"github.com/pake-go/pake-lib/config"
)

func Run(commands []pakelib.Command, logger *log.Logger) {
	cfg := config.New()
	for line, command := range commands {
		err := command.Execute(cfg, logger)
		if err != nil {
			errMsg := fmt.Errorf("There was an error at line %d: %s", line, err.Error())
			output.Error(errMsg)
		}
		cfg.SmartReset()
	}
}

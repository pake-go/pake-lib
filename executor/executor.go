package executor

import (
	pakelib "github.com/pake-go/pake-lib"
	"github.com/pake-go/pake-lib/config"
)

func Run(commands []pakelib.Command) {
	cfg := config.New()
	for _, command := range commands {
		command.Execute(cfg)
		cfg.SmartReset()
	}
}

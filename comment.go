package pakelib

import (
	"log"

	"github.com/pake-go/pake-lib/config"
)

type Comment struct {
}

func (c *Comment) Execute(cfg *config.Config, logger *log.Logger) error {
	return nil
}

type CommentValidator interface {
	IsValid(string) bool
}

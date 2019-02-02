package pakelib

import "github.com/pake-go/pake-lib/config"

type Comment struct {
}

func (c *Comment) Execute(cfg *config.Config) error {
	return nil
}

type CommentValidator interface {
	IsValid(string) bool
}

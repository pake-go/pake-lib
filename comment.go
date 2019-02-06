package pakelib

import (
	"log"

	"github.com/pake-go/pake-lib/config"
)

// Comment is used to represent the comment type in the language.
type Comment struct {
}

// Execute simply returns nil because there is nothing to be done for comments.
func (c *Comment) Execute(cfg *config.Config, logger *log.Logger) error {
	return nil
}

// CommentValidator is an interface that the comment validator for the language
// must satisfy.
type CommentValidator interface {
	// IsValid checks to see if the given string is a valid comment.
	IsValid(string) bool
}

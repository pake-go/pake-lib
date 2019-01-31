package pakelib

type Comment struct {
}

func (c Comment) Execute() (bool, error) {
	return true, nil
}

type CommentValidator interface {
	IsValid(string) bool
}

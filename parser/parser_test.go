package parser

import (
	"io/ioutil"
	"log"
	"reflect"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	pakelib "github.com/pake-go/pake-lib"
	"github.com/pake-go/pake-lib/config"
)

func TestNew(t *testing.T) {
	commandCandidates := []pakelib.CommandCandidate{
		helloCandidate,
		byeCandidate,
	}
	cv := &commentValidator{}

	parser := New(commandCandidates, cv)

	if !reflect.DeepEqual(commandCandidates, parser.commandCandidates) {
		t.Errorf("Expected %+v but got %+v", commandCandidates, parser.commandCandidates)
	}
	if !reflect.DeepEqual(cv, parser.commentValidator) {
		t.Errorf("Expected %+v but got %+v", cv, parser.commentValidator)
	}
}

func TestParseFile_noerror(t *testing.T) {
	commandCandidates := []pakelib.CommandCandidate{
		helloCandidate,
		byeCandidate,
	}
	cv := &commentValidator{}
	logger := log.New(ioutil.Discard, "", 0)

	parser := New(commandCandidates, cv)
	command, err := parser.ParseFile("validpakefile", logger)
	if err != nil {
		t.Error(err)
	}
	expected := []pakelib.Command{
		&hello{Args: []string{""}},
		&bye{Args: []string{""}},
	}
	if !cmp.Equal(command, expected) {
		t.Errorf("Expected %+v but got %+v", expected, command)
	}
}

func TestParseFile_witherror(t *testing.T) {
	commandCandidates := []pakelib.CommandCandidate{
		helloCandidate,
		byeWithErrorCandidate,
	}
	cv := &commentValidator{}
	logger := log.New(ioutil.Discard, "", 0)

	parser := New(commandCandidates, cv)
	_, err := parser.ParseFile("invalidpakefile", logger)
	if err == nil {
		t.Errorf("There should be an error parsing the given!")
	}
}

func TestParseString_noerror(t *testing.T) {
	commandCandidates := []pakelib.CommandCandidate{
		helloCandidate,
		byeCandidate,
	}
	cv := &commentValidator{}
	str := "hello \nbye "
	logger := log.New(ioutil.Discard, "", 0)

	parser := New(commandCandidates, cv)
	command, err := parser.ParseString(str, logger)
	if err != nil {
		t.Error(err)
	}
	expected := []pakelib.Command{
		&hello{Args: []string{""}},
		&bye{Args: []string{""}},
	}
	if !cmp.Equal(command, expected) {
		t.Errorf("Expected %+v but got %+v", expected, command)
	}
}

func TestParseString_witherror(t *testing.T) {
	commandCandidates := []pakelib.CommandCandidate{
		helloCandidate,
		byeWithErrorCandidate,
	}
	cv := &commentValidator{}
	str := "hello \nbyeWithError "
	logger := log.New(ioutil.Discard, "", 0)

	parser := New(commandCandidates, cv)
	_, err := parser.ParseString(str, logger)
	if err == nil {
		t.Errorf("There should be an error parsing the given!")
	}
}

func TestParseLine_noerror(t *testing.T) {
	commandCandidates := []pakelib.CommandCandidate{
		helloCandidate,
		byeCandidate,
	}
	cv := &commentValidator{}
	line := "hello "
	logger := log.New(ioutil.Discard, "", 0)

	parser := New(commandCandidates, cv)
	command, err := parser.ParseLine(line, logger)
	if err != nil {
		t.Error(err)
	}
	expected := &hello{Args: []string{""}}
	if !cmp.Equal(command, expected) {
		t.Errorf("Expected %+v but got %+v", expected, command)
	}
}

func TestParseLine_witherror(t *testing.T) {
	commandCandidates := []pakelib.CommandCandidate{
		helloCandidate,
		byeWithErrorCandidate,
	}
	cv := &commentValidator{}
	line := "byeWithError "
	logger := log.New(ioutil.Discard, "", 0)

	parser := New(commandCandidates, cv)
	_, err := parser.ParseLine(line, logger)
	if err == nil {
		t.Errorf("There should be an error parsing the given!")
	}
}

type commentValidator struct {
}

func (cv *commentValidator) IsValid(line string) bool {
	return strings.HasPrefix(line, "# ")
}

type hello struct {
	Args []string
}

func newHello(args []string) pakelib.Command {
	return &hello{
		Args: args,
	}
}

func (h *hello) Execute(cfg *config.Config, logger *log.Logger) error {
	return nil
}

type helloValidator struct {
}

func (hv *helloValidator) CanHandle(line string) bool {
	return strings.HasPrefix(line, "hello ")
}

func (hv *helloValidator) ValidateArgs(args []string) bool {
	return true
}

type bye struct {
	Args []string
}

func newBye(args []string) pakelib.Command {
	return &bye{
		Args: args,
	}
}

func (b *bye) Execute(cfg *config.Config, logger *log.Logger) error {
	return nil
}

type byeValidator struct {
}

func (bv *byeValidator) CanHandle(line string) bool {
	return strings.HasPrefix(line, "bye ")
}

func (bv *byeValidator) ValidateArgs(args []string) bool {
	return true
}

type byeWithError struct {
	Args []string
}

func newByeWithError(args []string) pakelib.Command {
	return &byeWithError{
		Args: args,
	}
}

func (bwe *byeWithError) Execute(cfg *config.Config, logger *log.Logger) error {
	return nil
}

type byeWithErrorValidator struct {
}

func (bwev *byeWithErrorValidator) CanHandle(line string) bool {
	return strings.HasPrefix(line, "byeWithError ")
}

func (bwev *byeWithErrorValidator) ValidateArgs(args []string) bool {
	return false
}

var helloCandidate = pakelib.CommandCandidate{
	Validator:   &helloValidator{},
	Constructor: newHello,
}

var byeCandidate = pakelib.CommandCandidate{
	Validator:   &byeValidator{},
	Constructor: newBye,
}

var byeWithErrorCandidate = pakelib.CommandCandidate{
	Validator:   &byeWithErrorValidator{},
	Constructor: newByeWithError,
}

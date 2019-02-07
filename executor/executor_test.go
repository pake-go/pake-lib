package executor

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"testing"

	capturer "github.com/kami-zh/go-capturer"
	pakelib "github.com/pake-go/pake-lib"
	"github.com/pake-go/pake-lib/config"
)

func TestRun_noerror(t *testing.T) {
	logOutput := bytes.Buffer{}
	logger := log.New(&logOutput, "", 0)

	output := capturer.CaptureOutput(func() {
		Run([]pakelib.Command{&hello{}, &bye{}}, logger)
	})

	expectedOutput := "Hello\nBye\n"
	if output != expectedOutput {
		t.Errorf("Expected %s but got %s", expectedOutput, output)
	}
	expectedLogOutput := ""
	if logOutput.String() != expectedLogOutput {
		t.Errorf("Expected %s but got %s", expectedLogOutput, logOutput.String())
	}
}

func TestRun_expectederror(t *testing.T) {
	logOutput := bytes.Buffer{}
	logger := log.New(&logOutput, "", 0)

	output := capturer.CaptureOutput(func() {
		Run([]pakelib.Command{&hello{}, &byeError{}}, logger)
	})

	expectedOutput := "Hello\n"
	if output != expectedOutput {
		t.Errorf("Expected %s but got %s", expectedOutput, output)
	}
	expectedLogOutput := "There was an error at line 2: Error from bye\n"
	if logOutput.String() != expectedLogOutput {
		t.Errorf("Expected %s but got %s", expectedLogOutput, logOutput.String())
	}
}

type hello struct {
	args []string
}

func (h *hello) Execute(cfg *config.Config, logger *log.Logger) error {
	fmt.Println("Hello")
	return nil
}

type bye struct {
	args []string
}

func (b *bye) Execute(cfg *config.Config, logger *log.Logger) error {
	fmt.Println("Bye")
	return nil
}

type byeError struct {
	args []string
}

func (be *byeError) Execute(cfg *config.Config, logger *log.Logger) error {
	return errors.New("Error from bye")
}

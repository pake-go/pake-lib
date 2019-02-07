package argutil

import (
	"reflect"
	"testing"
)

func TestGetTokens_noquotes(t *testing.T) {
	testString := "hello world bye"
	tokens, err := GetTokens(testString)
	if err != nil {
		t.Error(err)
	}

	expected := []string{"hello", "world", "bye"}
	// While reflect is very slow, we are concerned more about correctness
	// than performance
	if !reflect.DeepEqual(tokens, expected) {
		t.Errorf("Expected %+q, but got %+q", expected, tokens)
	}
}

func TestGetTokens_withquotes(t *testing.T) {
	testString := `hello "hello world" bye`
	tokens, err := GetTokens(testString)
	if err != nil {
		t.Error(err)
	}

	expected := []string{"hello", "hello world", "bye"}
	// While reflect is very slow, we are concerned more about correctness
	// than performance
	if !reflect.DeepEqual(tokens, expected) {
		t.Errorf("Expected %+q, but got %+q", expected, tokens)
	}
}

func TestGetTokens_expectederr(t *testing.T) {
	testString := `hello "bye die`
	_, err := GetTokens(testString)
	if err == nil {
		t.Errorf("Expected an error parsing due to missing end quote")
	}
}

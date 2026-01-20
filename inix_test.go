package inix

import (
	"strings"
	"testing"
)

func TestDuplicatedKey(t *testing.T) {
	data := `[data]
	test=hello
	test=hello`

	_, err := Parse(data)
	if err == nil {
		t.Fatal("No errors received")
	} else {
		if !strings.Contains(err.Error(), "already defined") {
			t.Fatal("Error is not about duplicated key")
		}
	}
}

func TestSectionNameWithSpaces(t *testing.T) {
	data := `[test data]
	test=hello`

	_, err := Parse(data)
	if err == nil {
		t.Fatal("No errors received")
	} else {
		if !strings.Contains(err.Error(), "section name cannot use spaces") {
			t.Fatal("Error is not about spaces in section name")
		}
	}
}

func TestKeyShouldBeDefinedInSection(t *testing.T) {
	data := "test=hello"

	_, err := Parse(data)
	if err == nil {
		t.Fatal("No errors received")
	} else {
		if !strings.Contains(err.Error(), "key should be defined in section") {
			t.Fatal("Error is not about key should be defined in section")
		}
	}
}

func TestInvalidKeySyntax(t *testing.T) {
	data := `[testdata]
	test hello`

	_, err := Parse(data)
	if err == nil {
		t.Fatal("No errors received")
	} else {
		if !strings.Contains(err.Error(), "invalid syntax") {
			t.Fatal("Error is not about invalid syntax")
		}
	}
}

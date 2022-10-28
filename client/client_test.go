package main

import (
	"testing"
)

// Parsing with a positive and negative int.
func TestParseInts_OK(t *testing.T) {
	a, b, err := parseInts("1", "-4")
	if err != nil || a != 1 || b != -4 {
		t.Errorf("parseInt with a=1, b=-4")
	}
}

// Cannot parse a float.
func TestParseInt_Err_Float(t *testing.T) {
	a, _, err := parseInts("1.5", "4")
	if err == nil {
		t.Errorf("parsed a float 1.5 to %d", a)
	}
}

// Cannot parse a string.
func TestParseInt_Err_String(t *testing.T) {
	a, _, err := parseInts("1string", "4")
	if err == nil {
		t.Errorf("parsed a string 1string to %d", a)
	}
}

// The input cannot be empty.
func TestParseInt_Err_Empty(t *testing.T) {
	a, _, err := parseInts("", "4")
	if err == nil {
		t.Errorf("parsed an empty string to %d", a)
	}
}

package strs

import (
	"testing"
)

// create a test function for SplitLast.
func TestSplitLast(t *testing.T) {
	s := "a/b/c"
	sep := "/"
	last := SplitLast(s, sep)
	if last != "c" {
		t.Errorf("SplitLast(%s, %s) = %s; want c", s, sep, last)
	}
}

// create a test function for SplitFirst
func TestSplitFirst(t *testing.T) {
	s := "a/b/c"
	sep := "/"
	first := SplitFirst(s, sep)
	if first != "a" {
		t.Errorf("SplitFirst(%s, %s) = %s; want a", s, sep, first)
	}
}

// create a test function for Reverse
func TestReverse(t *testing.T) {
	s := "abc"
	reversed := Reverse(s)
	if reversed != "cba" {
		t.Errorf("Reverse(%s) = %s; want cba", s, reversed)
	}
}

// create a test function for TrimMultiSpace
func TestTrimMultiSpace(t *testing.T) {
	s := "  a  b  c  "
	trimmed := TrimMultiSpace(s)
	if trimmed != "a  b  c" {
		t.Errorf("TrimMultiSpace(%s) = %s; want a b c", s, trimmed)
	}
}

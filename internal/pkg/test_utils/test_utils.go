// Package test_utils provides lightweight abstractions for working with vanilla tests
package test_utils

import (
	"strconv"
	"testing"
)

func Assert(sutName, description string, caseNum int, t *testing.T, expected, actual interface{}, e error) *testing.T {
	if e != nil {
		// Handle Error and return
		t.Errorf(sutName+" - "+description+" > case "+strconv.Itoa(caseNum)+": encountered an exception: %q", e)

		return t
	}

	if actual != expected {
		t.Errorf(sutName+" - "+description+" > case "+strconv.Itoa(caseNum)+":\nEXPECTED:\n%q\nRECEIVED:\n%q\n", expected, actual)

		return t
	}

	return t
}

func TestShouldThrow(sutName, description string, caseNum int, t *testing.T, err error) {
	if err == nil {
		t.Errorf(sutName + " - " + description + " > case " + strconv.Itoa(caseNum) + ": expected " + sutName + " to throw an error.")
	}
}

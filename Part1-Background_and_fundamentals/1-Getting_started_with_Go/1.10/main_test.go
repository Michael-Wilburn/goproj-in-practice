package main

import (
	"strings"
	"testing"
)

func TestName(t *testing.T) {
	name1 := reverseName(strings.TrimSpace("William"))
	expected1 := "mailliW"

	if !strings.EqualFold(name1, expected1) {
		t.Errorf("Response from reverseName is unexpected value. got [%s], expected [%s]", name1, expected1)
	}

	name2 := reverseName(strings.TrimSpace("Mister😎"))
	expected2 := "😎retsiM"

	if !strings.EqualFold(name2, expected2) {
		t.Errorf("Response from reverseName is unexpected value. got [%s], expected [%s]", name2, expected2)
	}
}

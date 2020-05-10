package maths_test

import (
	"github.com/tkech17/jack_compiler/utils/maths"
	"github.com/tkech17/jack_compiler/utils/tests"
	"testing"
)

func TestIsNumber(t *testing.T) {
	tests.AssertTrue(t, maths.IsNumber("123"), "129 Is Number")
	tests.AssertFalse(t, maths.IsNumber("a"), "a Is Not Number")
}

package functional

import (
	"github.com/tkech17/hach_virtual_machine_translator/utils/tests"
	"testing"
)

func TestFilter(t *testing.T) {
	var slice = []string{"1", "2", "22"}
	var filtered = Filter(slice, func(elem string) bool {
		return len(elem)%2 == 0
	})
	tests.AssertEqualsInt(t, 1, len(filtered))
	tests.AssertEqualsString(t, "22", filtered[0])
}

func TestMapString(t *testing.T) {
	var slice = []string{"1", "2"}
	var filtered = MapString(slice, func(elem string) string {
		return elem + elem
	})
	tests.AssertEqualsInt(t, 2, len(filtered))
	tests.AssertEqualsString(t, "11", filtered[0])
	tests.AssertEqualsString(t, "22", filtered[1])
}
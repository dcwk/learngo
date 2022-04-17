package prose

import (
	"fmt"
	"testing"
)

func TestTwoElements(t *testing.T) {
	list := []string{"apple", "banana"}
	want := "apple and banana"

	got := JoinWithCommas(list)

	if got != want {
		t.Error(errorString(list, got, want))
	}
}

func TestThreeElements(t *testing.T) {
	list := []string{"apple", "orange", "pear"}
	want := "apple, orange and pear"

	got := JoinWithCommas(list)

	if got != want {
		t.Error(errorString(list, got, want))
	}
}

func errorString(list []string, got string, want string) string {
	return fmt.Sprintf("JoinWithComma(%#v) = \"%s\", want \"%s\"", list, got, want)
}

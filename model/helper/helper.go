package helper

import (
	"reflect"
	"testing"
)

// ExpectError is ...
func ExpectError(name string, t *testing.T, f func()) {
	defer func(t *testing.T) {
		err := recover()
		if err == nil {
			t.Fatalf("Error Test: [%s] did not return error", name)
		}
		t.Logf("Error Test:Success! [%s]\n%s", name, err)
	}(t)
	f()
}

// DiffTest is ...
func DiffTest(want, got interface{}, t *testing.T) {
	if !reflect.DeepEqual(want, got) {
		t.Errorf(
			"Incorrect!\ngot: %v\nwant: %v.\n",
			got,
			want,
		)
	}
}

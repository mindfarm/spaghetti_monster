package foo_test

import (
	"github.com/mindfarm/spaghetti_god/foo"

	"testing"
)

var f *foo.Foo

func TestMain(m *testing.M) {
	f = &foo.Foo{}
}

func TestOuterFoo(t *testing.T) {
	testcases := map[string]struct {
		inputA int
		inputB int
		output int
	}{
		"Odd":  {1, 2, 3},
		"Even": {2, 4, 6},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {

			output := f.Adder(tc.inputA, tc.inputB)
			if output != tc.output {
				t.Fatal("Wut")
			}
		})
	}
}

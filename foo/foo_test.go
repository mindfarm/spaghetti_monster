package foo

import "testing"

func TestInnerFoo(t *testing.T) {
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
			f := &Foo{}
			output := f.Adder(tc.inputA, tc.inputB)
			if output != tc.output {
				t.Fatal("Wut")
			}
		})
	}
}

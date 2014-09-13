package space

import (
		"bytes"
		"testing"
)

func TestSaySpace(t *testing.T) {
		out := new(bytes.Buffer)
		SaySpace(out)
		checkString(t, "unexpected greeting", "hello from space\n", out.String())
}

func checkString(t *testing.T, message, expected, actual string) {
		if expected != actual {
				t.Errorf("%s: '%s' != '%s'", message, expected, actual)
		}
}


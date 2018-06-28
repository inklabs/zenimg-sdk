package zenimg_test

import (
	"testing"
)

func assertEqual(t *testing.T, expected string, actual string) {
	t.Helper()

	if actual != expected {
		t.Errorf("Expected [%s] got [%s]", expected, actual)
	}
}

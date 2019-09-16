package err

import (
	"gopkg.in/qamarian-dtp/err.v0"
	"gopkg.in/qamarian-lib/str.v2"
	"testing"
)

func TestFup (t *testing.T) {
	str.PrintEtr ("Testing function 'Fup ()' ...", "std", "TestFup ()")

	errA := err.New ("A", 1, 1)
	errB := err.New ("B", 1, 1, errA)
	errC := err.New ("C", 1, 1, errB)
	if Fup (errC) != "C [B [A]]" {
		str.PrintEtr ("Test failed!", "err", "TestFup ()")
		t.FailNow ()
	}

	str.PrintEtr ("Test passed!", "std", "TestFup ()")
}

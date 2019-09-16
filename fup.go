package err

import (
	"errors"
	"fmt"
)

// Fup () fully unwraps an error, then parses the result into a string. Take for example,
// err A (Foo) wraps err B (Bar), and err B (Bar) wraps err C (Tar). If we supply err A as
// an argument to this function, it would return a string like this: "Foo [Bar [Tar]]"
func Fup (err error) (string) {
	return fullUnwrapParse (err)
}

	// fullUnwrapParse () assists Fup (), that is the reason it was indented.
	func fullUnwrapParse (err error) (string) {
		secondary := errors.Unwrap (err)
		if secondary == nil {
			return err.Error ()
		} else {
			sec := fullUnwrapParse (secondary)
			return fmt.Sprintf ("%s [%s]", err.Error (), sec)
		}
	}

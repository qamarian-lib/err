package err

import (
	"errors"
	"fmt"
)

func Fup (err error) (string) {
	return fullUnwrapParse (err)
}

func fullUnwrapParse (err error) (string) {
	secondary := errors.Unwrap (err)
	if secondary == nil {
		return err.Error ()
	} else {
		sec := fullUnwrapParse (secondary)
		return fmt.Sprintf ("%s [%s]", err.Error (), sec)
	}
}

package errfmt

import "fmt"

// Wrap return wrapping error with message.
func Wrap(e error, msg string) error {
	return fmt.Errorf("%w: %s", e, msg)
}

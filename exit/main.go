
package exit

import (
	"os"
	"fmt"
	"errors"
)

func PrintError(errs ...error) error {
	e := errors.Join(errs...)
	if e != nil {
		// fmt.Fprintf(os.Stderr, "%v\n", e)
		panic(e)
	}
	return e
}

func Print_Msg(e error, msg ...any) error {
	if e != nil {
		fmt.Fprintln(os.Stderr, msg...)
		PrintError(e)
	}
	return e
}

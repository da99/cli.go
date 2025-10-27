
package exit

import (
	"os"
	"fmt"
)

func PrintError(e error) {
	if e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}
}

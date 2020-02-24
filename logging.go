package golo

import (
	"fmt"
	"os"
)

var (
	// LogChan takes output objects and, one by one, spits them out.
	// This operation is sequential, or at least more sequential than
	// before, to avoid crashing on `panic: too many concurrent operations on a single file or socket (max 1048575)`
	// errors
	LogChan = make(chan Output)
)

// logLoop iterates over the logChan, turns Outputs to json,
// and prints that json to STDOUT
func logLoop() {
	// Don't cause panics when logs can't be written
	defer func() {
		err := recover()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%+v\n", err)
		}
	}()

	for o := range LogChan {
		fmt.Println(o.String())
	}
}

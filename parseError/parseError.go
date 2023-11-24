package parseError

import (
	"fmt"
	"os"
)

var HadError = false

func RaiseError(line int, message string) {
	report(line, "", message)

}
func report(line int, where string, message string) {
	fmt.Fprintf(os.Stderr, "[line  %v ] Error %v :  %v", line, where, message)
	HadError = true
}

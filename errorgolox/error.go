package errorgolox

import "fmt"

var HadError = false

func LogError(line int, message string) {
	Report(line, "", message)
}

func Report(line int, where, message string) {
	fmt.Errorf("[line %d] Error %s: %s", line, where, message)
	HadError = true
}

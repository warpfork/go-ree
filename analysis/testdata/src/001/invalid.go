package main

func ErrorNotLast() (error, int) { // want "error should be returned as the last argument"
	return &Error{"hello-error"}, 0
}

// CallToInvalidFunction calls a function which has an error as non-last return argument.
//
// Errors:
//
//    - zonk-error -- is always returned
func CallToInvalidFunction() error { // want CallToInvalidFunction:"ErrorCodes: zonk-error"
	e, _ := ErrorNotLast() // want "unsupported: tracking error codes for function call with error as non-last return argument"
	if false {
		return e
	}
	return &Error{"zonk-error"}
}

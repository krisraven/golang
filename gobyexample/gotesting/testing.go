// example of unit testing
// testing package contains tools to write unit tests
// go test runs the tests

package main

/* usually the code that we're testing would be in a file with a name that is
*  named something like intmin.go and the test file would be intmin_test.go
*/
func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

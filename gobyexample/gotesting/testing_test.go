// example of unit testing
// testing package contains tools to write unit tests
// go test runs the tests

package main

import (
	"fmt"
	"testing"
)

func TestIntMinBasic(t *testing.T) { // prefix of Test is a convention for test
	ans := IntMin(2, -2)
	if ans != -2 {
		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}
} // t.Errorf reports the failures but continue the test. t.Fatal will report failures and stop the test

func TestIntMinTableDriven(t *testing.T) {
	var tests = []struct {
		a, b int // example. from {0, 1, -1}, 0 = a 1 = b
		want int // example: from {0, 1, -1}, -1 = want
	} {
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
		{-1, -1, -1},
	} // it's easier to write repeated tests with test data in a table

	for _, tt := range tests {
		testname := fmt.Sprintf("%d, %d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := IntMin(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

func BenchmarkIntMin(b *testing.B) {
	for i:= 0; i < b.N; i++ {
		IntMin(1, 2)
	}
}
// can run a single test with go test -v testing_test.go testing.go
// or all tests with go test -v
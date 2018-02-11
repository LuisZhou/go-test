package try_test

import (
	"os"
	"testing"
)

// if "?   	command-line-arguments	[no test files]"
// means name is not postfix by _test.go

// further reading: http://cs-guy.com/blog/2015/01/test-main/

func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags
	// os.Exit(m.Run())

	setup()
	ret := m.Run()
	if ret == 0 {
		teardown()
	}
	os.Exit(ret)
}

func TestA(t *testing.T) {
}

func TestB(t *testing.T) {
}

func setup() {
	println("setup")
}

func teardown() {
	println("teardown")
}

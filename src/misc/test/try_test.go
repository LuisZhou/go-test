package try_test

import (
	"bytes"
	"fmt"
	"html/template"
	"testing"
)

// Doc
// 1. https://golang.org/pkg/testing/
// 2. go test help (For more detail, run “go help test” and “go help testflag”)

// use -run to select which test function, the name using regular expression, means, you don't need to copy total name of the
// test.

// 1. To write a new test suite, create a file whose name ends _test.go that contains the TestXxx functions as described
//		here.
// 2. Put the file in the same package as the one being tested. The file will be excluded from regular package builds
//		but will be included when the “go test” command is run.
// 3. For more detail, run “go help test” and “go help testflag”

func TestTimeConsuming(t *testing.T) {
	// Tests and benchmarks may be skipped if not applicable with a call to the Skip method of *T and *B:
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	// Within these functions, use the Error, Fail or related methods to signal failure.
}

// Benchmarks

// 1. form: func BenchmarkXxx(*testing.B)
// 2. go test -bench=. misc/test/try_test.go
// 3. During benchmark execution, b.N is adjusted until the benchmark function lasts long enough to be timed reliably.
// 4. Output:
//		BenchmarkHello    10000000    282 ns/op
//		means that the loop ran 10000000 times at a speed of 282 ns per loop.
// 5. If a benchmark needs to test performance in a parallel setting, it may use the RunParallel helper function;
//		such benchmarks are intended to be used with the go test -cpu flag.

// carefule: if you do not use -bench, bench test will omitted.

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("hello")
	}
}

// If a benchmark needs some expensive setup before running, the timer may be reset
func BenchmarkBigLen(b *testing.B) {
	// big := NewBig()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// big.Len()
	}
}

// go test -cpu=4 -bench=. misc/test/try_test.go
func BenchmarkTemplateParallel(b *testing.B) {
	templ := template.Must(template.New("test").Parse("Hello, {{.}}!"))
	b.RunParallel(func(pb *testing.PB) {
		var buf bytes.Buffer
		for pb.Next() {
			buf.Reset()
			templ.Execute(&buf, "World")
		}
	})
}

// Example

// 1. Example functions may include a concluding line comment that begins with "Output:" and is compared with the standard
// 		output of the function when the tests are run. (The comparison ignores leading and trailing space.)
// 2. Can match unorder
// 3. Example functions without output comments are compiled but not executed
// 4. Naming convention. The suffix must start with a lower-case letter.
// 		func Example_suffix() 	{ ... }	// for package
// 		func ExampleF_suffix() 	{ ... }	// for function
// 		func ExampleT_suffix() 	{ ... }	// for type
// 		func ExampleT_M_suffix(){ ... }	// for method

func ExampleHello() {
	fmt.Println("hello")
	// Output: hello
}

func ExampleSalutations() {
	fmt.Println("hello, and")
	fmt.Println("goodbye")
	// Output:
	// hello, and
	// goodbye
}

// Unordered

// func ExamplePerm() {
// 	for _, value := range Perm(4) {
// 		fmt.Println(value)
// 	}
// 	// Unordered output: 4
// 	// 2
// 	// 1
// 	// 3
// 	// 0
// }

// Subtests and Sub-benchmarks

// 1. Only for type T and B
// 2. This enables uses like table-driven benchmarks and creating hierarchical tests.
// 3. It also provides a way to share common setup and tear-down code.
// 4. Each subtest and sub-benchmark has a unique name: the combination of the name of the top-level test and the
// 		sequence of names passed to Run, separated by slashes, with an optional trailing sequence number for disambiguation.
// 5. The argument to the -run and -bench command-line flags is an unanchored regular expression that matches the test's name.
//		(allow multi match, an empty expression matches any string)
// 6. Subtests can also be used to control parallelism. A parent test will only complete once all of its subtests complete.

// go test -run ''      # Run all tests.
// go test -run Foo     # Run top-level tests matching "Foo", such as "TestFooBar".
// go test -run Foo/A=  # For top-level tests matching "Foo", run subtests matching "A=".
// go test -run /A=1    # For all top-level tests, run subtests matching "A=1".

func TestFoo(t *testing.T) {
	// <setup code>
	t.Run("A=1", func(t *testing.T) { t.Log("A=1") })
	t.Run("A=2", func(t *testing.T) { t.Log("A=2") })
	t.Run("B=1", func(t *testing.T) { t.Log("B=2") })
	// <tear-down code>
}

var tests = []struct {
	Name string
}{{"A"}}

func TestGroupedParallel(t *testing.T) {
	for _, tc := range tests {
		tc := tc // capture range variable
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			t.Log("I'm here")
		})
	}
}

// Run does not return until parallel subtests have completed, providing a way to clean up after a group of parallel tests

func TestTeardownParallel(t *testing.T) {
	// This Run will not return until the parallel tests finish.
	t.Run("group", func(t *testing.T) {
		t.Run("Test1", func(t *testing.T) {})
		t.Run("Test2", func(t *testing.T) {})
		t.Run("Test3", func(t *testing.T) {})
	})
	// <tear-down code>
}

// -run and -bench using regular expression.

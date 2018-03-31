package main

import (
	"flag"
	"runtime/pprof"
)

// try this one:
// To start tuning the Go program, we have to enable profiling. If the code used the Go testing package's benchmarking
// support, we could use gotest's standard -cpuprofile and -memprofile flags.

// After adding that code, we can run the program with the new -cpuprofile flag and then run `go tool pprof` to interpret the profile.
// ./profile -cpuprofile=profile.prof
// go tool pprof profile profile.prof

// [Google's pprof C++ profiler]: https://github.com/gperftools/gperftools
// topN

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()

		// When CPU profiling is enabled, the Go program stops about 100 times per second and records a sample consisting of
		// the program counters on the currently executing goroutine's stack.

		// So understand the stack is useful.

		// The first two columns show the number of samples in which the function was running (as opposed to waiting for a
		// called function to return)

		// The third column shows the running total during the listing

		// The fourth and fifth columns show the number of samples in which the function appeared (either running or waiting
		// for a called function to return)

		// was running in 10.6% of the samples
		// on the call stack (it or functions it called were running) in 84.1% of the samples
	}
}

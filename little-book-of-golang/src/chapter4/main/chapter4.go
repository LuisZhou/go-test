// Chapter 4 - Code Organization and Interfaces

// Packages

// In Go, package names follow the directory structure of your Go workspace.

// package shopping <--> $GOPATH/src/shopping/ . In other words, when you name a package, via the package keyword,
// you provide a single value, not a complete hierarchy.

// Question:
// 1. same package, can't name conflict?
// 2. go get url, will it get all other dependence of the download repo?

// __important__ go compiler compile all source to the binary, so no anyother dependence.

// you can also rename this file to shopping/main.go, and run go run main/main.go, the result is same.

package main

import (
	"fmt"
	"shopping"
)

type Logger interface {
	Log(message string)
}

// interface{} empty interface.

type ConsoleLogger struct{}

func (l ConsoleLogger) Log(message string) {
	fmt.Println(message)
}

func process(logger Logger) {
	logger.Log("hello!")
}

func main() {
	fmt.Println(shopping.PriceCheck(4343))

	// __Section__: Cyclical Imports
	// We solve this by introducing another package which contains shared structures.
	// In a few sections, we'll look at interfaces which can help us untangle these types of dependencies.

	// __Section__: Visibility
	// Go uses a simple rule to define what types and functions are visible outside of a package.
	// If the name of the __type or function__ starts with an uppercase letter, it's visible.
	// If it starts with a lowercase letter, it isn't.

	// This also applies to structure fields. If a structure field name starts with a lowercase letter,
	// only code within the same package will be able to access them.

	// __Section__: Package Management
	// go get github.com/mattn/go-sqlite3 -> $GOPATH/src

	// __Section__: Dependency Management (revert operation)
	// we go get within a project, it'll scan all the files,
	// looking for imports to third-party libraries and will download them.
	// go get :install. In a way, our own source code becomes a __Gemfile__ or __package.json__
	// go get -u :update.
	// go get -u FULL_PACKAGE_NAME

	// go get without parameter, should run in the main package of the project.

	// For one thing, there's no way to specify a revision, it always points to the master/head/trunk/default.
	// This is an even larger problem if you have two projects needing different versions of the same library.

	// __Section__: Interfaces
	// Interfaces are types that define a contract but not an implementation.
	// Interfaces help decouple your code from specific implementations.

	// In Go, this happens implicitly.

	consoleLogger := ConsoleLogger{}
	// consoleLogger.Log("I'm a struct implement interface.")
	process(consoleLogger)

	// interfaces value:
	// + participate in composition.
	// + interfaces are commonly used to avoid cyclical imports.
}

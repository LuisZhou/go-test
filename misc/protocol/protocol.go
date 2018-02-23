package main

// ref
// https://developers.google.com/protocol-buffers/docs/gotutorial

// why
// other option:
// + gobs: only Go-specific environment.
// + ad-hoc way: should write your own code to encode and decode, parse runtime cost(protocol is not a parse process)
// + XML/json: human readable, support by a lot of language, space intensive, performance penalty when encode and decode,
//	 navigating is more complicated than a class.
// + protocol: flexible, efficient, automated solution, extendable.

// work flow:
// + Define message formats in a .proto file.
// + Use the protocol buffer compiler.
// + Use the Go protocol buffer API to write and read messages.
//	 From that, the protocol buffer compiler creates a class that implements automatic encoding and parsing of the
//	 protocol buffer data with an efficient binary format. The generated class provides getters and setters for the
//	 fields that make up a protocol buffer and takes care of the details of reading and writing the protocol buffer as a
//	 unit.

// important && careful:
// + read the readme in language plugin. the zip include bin and include. The include file is very very often import by
//	 by other protocol file.
//	 https://developers.google.com/protocol-buffers/docs/downloads.html
//
// + The compiler plugin protoc-gen-go will be installed in $GOBIN, defaulting to $GOPATH/bin. It must be in your $PATH
//	 for the protocol compiler protoc to find it.
//	 So you should type 'which protoc-gen-go' to test if the protoc-gen-go is in your path.
//	 And once 'go get -u github.com/golang/protobuf/protoc-gen-go' it will locate in $GOPATH/bin.
//	 So the compile process is: protoc -> protoc-gen-go
//
// + read careful about the help of the protoc and protoc-gen-go.

// how to write protocol
// + package declaration
//	 In Go, the package name is used as the Go package, unless you have specified a go_package. Even if you do provide a
//	 go_package, you should still define a normal package as well to avoid name collisions in the Protocol Buffers name
//	 space as well as in non-Go languages.
// + have your message definitions

func main() {

}

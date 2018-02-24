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
//
// + if you get:
//	 google/protobuf/timestamp.proto: File not found. just do:
//	 sudo chmod 777 -R /usr/local/include/google/

// how to write protocol
// + package declaration
//	 In Go, the package name is used as the Go package, unless you have specified a go_package. Even if you do provide a
//	 go_package, you should still define a normal package as well to avoid name collisions in the Protocol Buffers name
//	 space as well as in non-Go languages.
// + have your message definitions

//	 simple data types, using other message types

//   You can even define message types nested inside other messages
//	 also define enum types if you want one of your fields to have one of a predefined list of values

//	 The " = 1", " = 2" markers on each element identify the unique "tag" that field uses in the binary encoding.
//	 Tag numbers 1-15 require one less byte to encode than higher numbers, so as an optimization you can decide to use
//	 those tags for the commonly used or repeated elements, leaving tags 16 and higher for less-commonly used optional
//	 elements. Each element in a repeated field requires re-encoding the tag number, so repeated fields are particularly
//	 good candidates for this optimization.

//	 If a field value isn't set, a __default value__ is used: zero for numeric types, the empty string for strings, false
//	 for bools. For embedded messages, the default value is always the "default instance" or "prototype" of the message,
//	 which has none of its fields set. Calling the accessor to get the value of a field which has not been explicitly
//	 set always returns that field's default value.

// 	 __repeat__: If a field is repeated, the field may be repeated any number of times (including zero). The order of
//	 the repeated values will be preserved in the protocol buffer. Think of repeated fields as dynamically sized arrays.

//	 complete guide: https://developers.google.com/protocol-buffers/docs/proto3

// question
// + why 'go get -u github.com/golang/protobuf/protoc-gen-go' will install protoc-gen-go to $GOBIN?
func main() {

}

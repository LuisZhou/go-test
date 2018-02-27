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
// + why we need to tell the compile the source code location: protoc -I=. --go_out=. ./addressbook.proto
// + think about, why protocol use the pointer all the time.

// The Protocol Buffer API

// Generating addressbook.pb.go gives you the following useful types:
//     An AddressBook structure with a People field.
//     A Person structure with fields for Name, Id, Email and Phones.
//     A Person_PhoneNumber structure, with fields for Number and Type.
//     The type Person_PhoneType and a value defined for each value in the Person.PhoneType enum.

// careful about enum, it will become const with name prefix by parent's name and underscope

// c/c++ runtime normal use the option optimize_for = LITE_RUNTIME. for the lite version runtime is always smaller.

import (
	proto "github.com/golang/protobuf/proto"
	"io/ioutil"
	"log"
	pb "misc/protocol/protocol"
	"reflect"
)

func main() {
	book := &pb.AddressBook{[]*pb.Person{ // repeat will become slice of pointer. why use pointer?
		{
			Name:  "John Doe",
			Id:    101,
			Email: "john@example.com",
		},
		{
			Name: "Jane Doe",
			Id:   102,
		},
		{
			Name:  "Jack Doe",
			Id:    201,
			Email: "jack@example.com",
			Phones: []*pb.Person_PhoneNumber{
				{Number: "555-555-5555", Type: pb.Person_WORK},
			},
		},
		{
			Name:  "Jack Buck",
			Id:    301,
			Email: "buck@example.com",
			Phones: []*pb.Person_PhoneNumber{
				{Number: "555-555-0000", Type: pb.Person_HOME},
				{Number: "555-555-0001", Type: pb.Person_MOBILE},
				{Number: "555-555-0002", Type: pb.Person_WORK},
			},
		},
		{
			Name:  "Janet Doe",
			Id:    1001,
			Email: "janet@example.com",
			Phones: []*pb.Person_PhoneNumber{
				{Number: "555-777-0000"},
				{Number: "555-777-0001", Type: pb.Person_HOME},
			},
		},
	}}

	type test struct{}

	//var tmp *test = test{}
	tmp := []*test{
		{},
	}

	for _, c := range tmp {
		log.Println(reflect.TypeOf(c))
	}

	//
	// []byte{1, 2}
	// []byte("abc")

	out, err := proto.Marshal(book)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}
	if err := ioutil.WriteFile("./bin/test", out, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}

	// read

	in, err := ioutil.ReadFile("./bin/test")
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	book = &pb.AddressBook{}
	if err := proto.Unmarshal(in, book); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}

	for _, c := range book.People {
		log.Println(c)
	}

	log.Println("ok")

	// ref
	// https://godoc.org/github.com/golang/protobuf/proto
}

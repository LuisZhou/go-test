package main

// Communication between a client and a service requires the exchange of data. This data may be highly structured, but
// has to be serialised for transport.

// words
// + protocols in two process

// ## Introduce

// problem:
// Messages are sent across the network as a sequence of bytes, which has no structure except for a linear stream of bytes.
// an application has to serialise any data into a stream of bytes in order to write it, and deserialise the stream of
// bytes back into suitable data structures on reading it.
// These two operations are known as marshalling and unmarshalling respectively.

// consider sending the following variable length table of two columns of variable length strings:
// fred			programmer
// liping		analyst
// sureerat	manager

// 1. length method:
// yes, you have assumed about the datastruct, and tell to counterpart length of the following char.
// but this method is not portable(global use by other program).

// what is the size of len, and is it big endian or little endian.

// 3                	// 3 	rows, 2 columns assumed
// 4	fred           	// 4 	char string,	col 1
// 10 programmer    	// 10 char string,	col 2
// 6 	liping         	// 6 	char string, 	col 1
// 7 	analyst        	// 7 	char string, 	col 2
// 8 	sureerat       	// 8 	char string, 	col 1
// 7 	manager        	// 7 	char string, 	col 2

// 2. control-char method:
// read as much as you can, until the control-char, meaning end of the  char.

// 3
// fred\0
// programmer\0
// liping\0
// analyst\0
// sureerat\0
// manager\0

// 3. fix size methdo:
// 3-row fixed table of two columns of strings of length 8 and 10 respectively

// fred\0\0\0\0		// 8
// programmer
// liping\0\0			// 8
// analyst\0\0\0
// sureerat				// 8
// manager\0\0\0

// this methods only solve serilize the chars(string), not solve the problem of any other types, such as integer,
// what is the size of it, and is it big endian or little endian.

// ## Mutual agreement

// question:
// + how many rows are possible for the table - that is, how big an integer do we need to describe the row size. A
//	 similar problem occurs for the length of each string.
// + which character set do chars belong.

// The above serialisation is opaque or implicit. If data is marshalled using the above format, then there is nothing in
// the serialised data to say how it should be unmarshalled. The unmarshalling side has to know exactly how the data is
// serialised in order to unmarshal it correctly.

// early well-known serialisation method
// XDR (external data representation) used by Sun's RPC, later known as ONC (Open Network Computing)
// + RFC 1832
// + precise
// + inherently type-unsafe
// + use: compilers generating code for both marshalling and unmarshalling

// about go:
// go no explicit support for __opaque serialised data__
// go use uses "gob" serialisation for package RPC.

// ## Self-describing data
// Self-describing data carries type information along with the data.

// table (should have some method to present the type, for example, little integer)
//     uint8 3
//     uint 2
// string
//     uint8 4
//     []byte fred
// string
//     uint8 10
//     []byte programmer
// string
//     uint8 6
//     []byte liping
// string
//     uint8 7
//     []byte analyst
// string
//     uint8 8
//     []byte sureerat
// string
//     uint8 7
//     []byte manager

// the principle is that the marshaller will generate such type information in the serialised data.
// The unmarshaller will know the type-generation rules and will be able to use this to reconstruct the correct data structure.

func main() {

}

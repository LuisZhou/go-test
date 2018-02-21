// support:

// + JSON serialises objects, arrays and basic values
// + The basic values include string, number, boolean values and the null value.
// + Arrays are a comma-separated list of values that can represent arrays, vectors, lists or sequences of various
//	 programming languages. They are delimited by square brackets  [ ... ] .
// + Objects are represented by a list of "field: value" pairs enclosed in curly braces  { ... } .
// + There is no special support for complex data types such as dates, no distinction between number types, no recursive
//	 types, etc.

// property:

// + text-based format
// + overheads of string handling

// specification about golang json marshalling.

// ...
//
// + Struct values encode as JSON objects. Each struct field becomes a member of the object. By default the object's key
//	 name is the struct field name __converted to lower case__. If the struct field has a tag, that tag will be used as the
//	 name instead.
//
// + Map values encode as JSON objects. The map's key type must be string; the object keys are used directly as map keys.
//
// + Pointer values encode as the value pointed to. (Note: this allows trees, but not graphs!). A nil pointer encodes as
//	 the null JSON object.
//
// + Interface values encode as the value contained in the interface. A nil interface value encodes as the null JSON object.
//
// + Channel, complex, and function values cannot be encoded in JSON. Attempting to encode such a value cause Marshal to
//	 return an  InvalidTypeError .
//
// + JSON cannot represent cyclic data structures and Marshal does not handle them. Passing cyclic structures to Marshal
//	 will result in an infinite recursion.
//

/* SaveJSON
 */

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name  Name
	Email []Email
}

type Name struct {
	Family   string
	Personal string
}

type Email struct {
	Kind    string
	Address string
}

func main() {
	person := Person{
		Name: Name{Family: "Newmarch", Personal: "Jan"},
		Email: []Email{Email{Kind: "home", Address: "jan@newmarch.name"},
			Email{Kind: "work", Address: "j.newmarch@boxhill.edu.au"}}}

	saveJSON("./bin/person.json", person)
}

func saveJSON(fileName string, key interface{}) {
	outFile, err := os.Create(fileName)
	checkError(err)
	encoder := json.NewEncoder(outFile)
	err = encoder.Encode(key)
	checkError(err)
	outFile.Close()
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}

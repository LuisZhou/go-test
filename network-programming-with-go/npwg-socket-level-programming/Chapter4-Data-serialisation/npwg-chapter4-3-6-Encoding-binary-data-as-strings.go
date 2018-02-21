package main

// For example, in an "even parity" scheme, the check digit would be set to one or zero to make an even number of 1-s
// in a byte. This allows detection of errors of a __single bit__ in each byte. (can't check one more error, ether can
// know which position.)

/**
 * Base64
 */

// ref:
// https://en.wikipedia.org/wiki/Base64
// https://en.wikipedia.org/wiki/MIME#Content-Transfer-Encoding
// https://en.wikipedia.org/wiki/Binary_data
// https://en.wikipedia.org/wiki/Binary-to-text_encoding
// https://en.wikipedia.org/wiki/Parity_bit
// https://en.wikipedia.org/wiki/Error_detection_and_correction

// question:
// what benifi of convert 8-bit binary data into 7-bit ASCII, will network can know what kind data, and make usage of ASCII?
// How ascII use parity, in which part?

// important:
// No matter how parity, base64 is useful for making the data transfer safy, make the server/router/client not to modify
// or detect the content.
// ref: https://stackoverflow.com/questions/3538021/why-do-we-use-base64

import (
	"bytes"
	"encoding/base64"
	"fmt"
)

func main() {
	eightBitData := []byte{1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Println(eightBitData)
	// 00000001 00000010 0000011
	// 000000 010000 001000 000011
	// A Q I D
	bb := &bytes.Buffer{}
	encoder := base64.NewEncoder(base64.StdEncoding, bb)
	encoder.Write(eightBitData)
	encoder.Close()
	fmt.Println(bb)

	dbuf := make([]byte, 12)
	decoder := base64.NewDecoder(base64.StdEncoding, bb)
	decoder.Read(dbuf)
	for _, ch := range dbuf {
		fmt.Print(ch)
	}

	fmt.Println()
}

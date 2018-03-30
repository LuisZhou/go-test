/* Blowfish
 */

package main

import (
	"bytes"
	"fmt"
	"golang.org/x/crypto/blowfish"
)

// uses a single key that is the same for both encryption and decryption.
// there are many encryption algorithms.

// Go has support for several symmetric key algorithms such as Blowfish and DES.

// The algorithms are block algorithms. That is they work on blocks of data. If your data is not aligned to the block
// size, then you will have to pad it with extra blanks at the end.

// Each algorithm is represented by a Cipher object.

// The blocks have to be 8-byte blocks for Blowfish.

func main() {
	key := []byte("my key")
	cipher, err := blowfish.NewCipher(key)
	if err != nil {
		fmt.Println(err.Error())
	}
	src := []byte("hello\n\n\n")
	var enc [512]byte

	cipher.Encrypt(enc[0:], src)

	var decrypt [8]byte
	cipher.Decrypt(decrypt[0:], enc[0:])
	result := bytes.NewBuffer(nil)
	result.Write(decrypt[0:8])
	fmt.Println(string(result.Bytes()))
}

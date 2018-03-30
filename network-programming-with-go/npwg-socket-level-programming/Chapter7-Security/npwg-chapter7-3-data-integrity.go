/* MD5Hash
 */

package main

import (
	"crypto/md5"
	"fmt"
)

// Ensuring data integrity means supplying a means of testing that the data has not been tampered with. Usually this is
// done by forming a simple number out of the bytes in the data. This process is called hashing and the resulting number
// is called a hash or hash value.

// naive hashing algorithm: sum up all the bytes in the data
// attacker could just swap two bytes

// Hashing algorithms used for security purposes have to be "strong"
// cryptographic hashing algorithms
// MD4, MD5, RIPEMD-160, SHA1, SHA224, SHA256, SHA384 and SHA512

// a function  New  (or similar) in the appropriate package returns a  Hash  object from the  hash  package.

func main() {
	hash := md5.New()
	// func NewMD5(key []byte) hash.Hash
	bytes := []byte("hello\n")
	hash.Write(bytes)
	hashValue := hash.Sum(nil)
	hashSize := hash.Size()
	for n := 0; n < hashSize; n += 4 {
		var val uint32
		val = uint32(hashValue[n])<<24 +
			uint32(hashValue[n+1])<<16 +
			uint32(hashValue[n+2])<<8 +
			uint32(hashValue[n+3])
		fmt.Printf("%x ", val)
	}
	fmt.Println()

	fmt.Printf("%x", md5.Sum([]byte("hello\n")))
	fmt.Println()

	// no this function now.
	// hash2 := md5.NewMD5("hello\n")
	// _ := hash2
}

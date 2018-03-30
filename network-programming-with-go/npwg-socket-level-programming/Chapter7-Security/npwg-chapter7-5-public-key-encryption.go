/* GenRSAKeys
 */

package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/gob"
	"encoding/pem"
	"fmt"
	"os"
)

// this chapter just tech how to create and load public key, and private key.

// Public key encryption and decryption requires two keys: one to encrypt and a second one to decrypt. The encryption
// key is usually made public in some way so that anyone can encrypt messages to you. The decryption key must stay
// private, otherwise everyone would be able to decrypt those messages! Public key systems are asymmetric, with
// different keys for different uses.

// There are many public key __encryption systems__ supported by Go. A typical one is the RSA scheme.

func main() {
	reader := rand.Reader
	bitSize := 512
	key, err := rsa.GenerateKey(reader, bitSize)
	checkError(err)

	fmt.Println("Private key primes", key.Primes[0].String(), key.Primes[1].String())
	fmt.Println("Private key exponent", key.D.String())

	publicKey := key.PublicKey
	fmt.Println("Public key modulus", publicKey.N.String())
	fmt.Println("Public key exponent", publicKey.E)

	saveGobKey("private.key", key)
	saveGobKey("public.key", publicKey)

	savePEMKey("private.pem", key)

	// load
	{
		var key rsa.PrivateKey
		loadKey("private.key", &key)

		fmt.Println("Private key primes", key.Primes[0].String(), key.Primes[1].String())
		fmt.Println("Private key exponent", key.D.String())

		var publicKey rsa.PublicKey
		loadKey("public.key", &publicKey)

		fmt.Println("Public key modulus", publicKey.N.String())
		fmt.Println("Public key exponent", publicKey.E)
	}
}

func saveGobKey(fileName string, key interface{}) {
	outFile, err := os.Create(fileName)
	checkError(err)
	encoder := gob.NewEncoder(outFile)
	err = encoder.Encode(key)
	checkError(err)
	outFile.Close()
}

func savePEMKey(fileName string, key *rsa.PrivateKey) {
	outFile, err := os.Create(fileName)
	checkError(err)

	var privateKey = &pem.Block{Type: "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key)}

	pem.Encode(outFile, privateKey)

	outFile.Close()
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}

func loadKey(fileName string, key interface{}) {
	inFile, err := os.Open(fileName)
	checkError(err)
	decoder := gob.NewDecoder(inFile)
	err = decoder.Decode(key)
	checkError(err)
	inFile.Close()
}

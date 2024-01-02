/*
https://www.systutorials.com/how-to-generate-rsa-private-and-public-key-pair-in-go-lang/

https://www.systutorials.com/tag/golang/

The Crypto standard libraries
crypto/rsa - implements RSA encryption as specifed in PKCS#1 and RFC 8017

func GenerateKey(random io.Reader, bits int) (*PrivateKey, error)


crypto/rand - The var Reader io.Reader struct is a global, shared instance of a
cryptographically secure random number generator.


crypto/x509
func MarshalPKCS1PrivateKey(key *rsa.PrivateKey) []byte
	converts an RSA private key to PKCS #1, ASN.1 DER form.
	This kind of key is commonly encoded in PEM blocks
	of type "RSA PRIVATE KEY".


func MarshalPKCS1PublicKey(key *rsa.PublicKey) []byte
	converts an RSA public key to PKCS #1, ASN.1 DER form,
	This kind of key is commonly encoded in PEM blocks of
	type "RSA PUBLIC KEY".


encoding/pem
func Encode(out io.Writer, b *Block) error
	writes the PEM encoding of b to out
*/

package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func main() {
	privatekey, err := rsa.GenerateKey(rand.Reader, 2048)

	if err != nil {
		fmt.Printf("Cannot generate RSA key\n")
		os.Exit(1)
	}

	publickey := &privatekey.PublicKey

	var privateKeyBytes []byte = x509.MarshalPKCS1PrivateKey(privatekey)
	privateKeyBlock := pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	}
	privatePem, err := os.Create("private.pem")
	if err != nil {
		fmt.Printf("error when create private.pem: %s \n", err)
		os.Exit(1)
	}
	err = pem.Encode(privatePem, &privateKeyBlock)
	if err != nil {
		fmt.Printf("error when encode priavate pem: %s \n", err)
		os.Exit(1)
	}

	// dump public key to file
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publickey)
	if err != nil {
		fmt.Printf("error when dumping publickey: %s \n", err)
		os.Exit(1)
	}
	publicKeyBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}
	publicPem, err := os.Create("public.pem")
	if err != nil {
		fmt.Printf("error when create public.pem: %s \n", err)
		os.Exit(1)
	}
	err = pem.Encode(publicPem, publicKeyBlock)
	if err != nil {
		fmt.Printf("error when encode public pem: %s \n", err)
		os.Exit(1)
	}
}

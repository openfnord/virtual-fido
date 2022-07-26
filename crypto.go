package main

import (
	"crypto/aes"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
)

func encrypt(key []byte, data []byte) []byte {
	deviceCipher, err := aes.NewCipher(key)
	checkErr(err, "Could not create device cipher")
	encryptedData := make([]byte, len(data))
	deviceCipher.Encrypt(encryptedData, data)
	return encryptedData
}

func sign(key *ecdsa.PrivateKey, data []byte) []byte {
	hash := sha256.Sum256(data)
	signature, err := ecdsa.SignASN1(rand.Reader, key, hash[:])
	checkErr(err, "Could not sign data")
	return signature
}
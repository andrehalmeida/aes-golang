package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
)

func main() {
	fmt.Println("Hello World")
	text := []byte("Hey guys! It's a  nice day for programming.")
	secret := []byte("passphrasewhichneedstobe32bytes!")
	c, err := aes.NewCipher(secret)
	if err != nil {
		fmt.Println(err)
	}

	gcm, err := cipher.NewGCM(c)

	if err != nil {
		fmt.Println(err)
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err)
	}

	fmt.Println(gcm.Seal(nonce, nonce, text, nil))
	err = ioutil.WriteFile("myfile.data", gcm.Seal(nonce, nonce, text, nil), 0777)
	if err != nil {
		fmt.Println(err)
	}
}

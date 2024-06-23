package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

// So basically, we can used the values within the same package level of code...
// I can use the Md5Hashing function without even importing it to above^
func EncryptIt(value []byte, keyPhrase string) []byte {
	// NewCipher method takes a keyPhrase which in this case, you hashed for increased security.
	aesBlock, err := aes.NewCipher([]byte(Md5Hashing(keyPhrase)))
	if err != nil {
		panic(err) // Use panic cause we want to terminate the program if error occurs
	}

	gcmInstance, err := cipher.NewGCM(aesBlock)
	if err != nil {
		panic(err)
	}

	nonce := make([]byte, gcmInstance.NonceSize())
	_, _ = io.ReadFull(rand.Reader, nonce)

	cipheredText := gcmInstance.Seal(nonce, nonce, value, nil)
	return cipheredText
}

func DecryptIt(ciphered []byte, keyPhrase string) []byte {
	// NewCipher method takes a keyPhrase which in this case, you hashed for increased security.
	aesBlock, err := aes.NewCipher([]byte(Md5Hashing(keyPhrase)))
	if err != nil {
		panic(err)
	}

	gcmInstance, err := cipher.NewGCM(aesBlock)
	if err != nil {
		panic(err)
	}

	nonceSize := gcmInstance.NonceSize()
	fmt.Println(ciphered[:nonceSize])
	nonce, cipheredText := ciphered[:nonceSize], ciphered[nonceSize:]

	originalText, err := gcmInstance.Open(nil, nonce, cipheredText, nil)
	if err != nil {
		panic(err)
	}

	return originalText
}

package encryption

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
)

// We can use the md5 hashed to hash the password of a logged in user and compare the input hashed password
// Once logged in and hashed them again to the compared values.
func Md5Hashing(input string) string {
	byteInput := []byte(input)
	md5Hash := md5.Sum(byteInput)
	// [:] This symbol it means it will display all the array elements
	// It was just like spread in JavaScript
	return hex.EncodeToString(md5Hash[:])
}

func ShaHashing(input string) string {
	plainText := []byte(input)
	sha256Hash := sha256.Sum256(plainText)
	return hex.EncodeToString(sha256Hash[:])
}

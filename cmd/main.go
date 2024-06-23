package main

import (
	"fmt"
	encryption "llauderesv/crypto/lib"
)

func main() {
	str := "Vince Llauderes"
	str2 := "Vince Llauderes"

	fmt.Println(str[:])

	hashed := encryption.Md5Hashing(str)
	hashed2 := encryption.Md5Hashing(str2)
	fmt.Printf("Hashed version of this string using MD5 %s - %s\n", str, hashed)

	fmt.Printf("Is Two Values Equal? %t\n", hashed == hashed2)

	s := "Vince Llauderes"

	encryptS := encryption.ShaHashing(s)
	fmt.Printf("Hashed version of this string using SHA256 %s\n", encryptS)

	str3 := "This is a sample value that you wanted to encrypt"

	fmt.Println("===============Encrypt / Decrypt Values=============")
	fmt.Println("Byte Format")
	fmt.Println(encryption.EncryptIt([]byte(str3), "private.pem"))

	fmt.Println("====================================================")
	fmt.Println("String Format")
	fmt.Println(string(encryption.EncryptIt([]byte(str3), "private.pem")))

	fmt.Println("Decrypting values")
	fmt.Println(string(encryption.DecryptIt(encryption.EncryptIt([]byte(str3), "private.pem"), "private.pem")))
}

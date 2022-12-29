package main

import (
	"crypto/aes"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// cache map
var data = make(map[string]string)

var cipher = "thisis32bitlongpassphraseimusing"

func main() {

	// key := "keyyyy1"
  // value := "valueyyyy1"
  // v := fmt.Sprintf("%s # %s \n", key, value)
    // w := EncryptAES([]byte(cipher), v)
	  // fmt.Println(w)

	// DecryptAES([]byte(key), c)
	// fmt.Println(c)

	insert("key1111", "value1111")
	insert("key2222", "value2222")
	 writeEncrypt()
	 erase()
	 readDecrypt()
   display()
}

func insert(key, value string) {
	data[key] = value
}

func deleteByKey(key string) {
	_, exists := data[key]
	if exists {
		delete(data, key)
	}
}

func display() {
	for key, name := range data {
		fmt.Println(key + " " + name)
	}
}

func erase() {
	for k := range data {
		delete(data, k)
	}
}

func writeEncrypt() {
	file, _ := os.Create("data.txt")
	defer file.Close()

	for key, value := range data {
    v := fmt.Sprintf("%s # %s \n", key, value)
     w := EncryptAES([]byte(cipher), v)
		file.WriteString(w)
	}
}

func readDecrypt() {
	v, _ := ioutil.ReadFile("data.txt")
	z := string(v)
  zz := DecryptAES([]byte(cipher), z)
	w := strings.Split(zz, "\n")

	for _, x := range w {
	  y := strings.Split(x, "#")
    if (len(y) ==2) {
		   k := y[0]
	     v := y[1]
       insert(k, v)
      }
    }
}

func EncryptAES(key []byte, plaintext string) string {

	c, err := aes.NewCipher(key)
	CheckError(err)

	out := make([]byte, len(plaintext))

	c.Encrypt(out, []byte(plaintext))

	return hex.EncodeToString(out)
}

func DecryptAES(key []byte, ct string) string {
	ciphertext, _ := hex.DecodeString(ct)

	c, err := aes.NewCipher(key)
	CheckError(err)

	pt := make([]byte, len(ciphertext))
	c.Decrypt(pt, ciphertext)

  return string(pt[:])
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

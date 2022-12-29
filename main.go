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

var on = false

func main() {
	insert("key111", "value1111")
	insert("key222", "value2222")
  insert("key333", "value3333")
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
  fmt.Println("display cache: ")
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
     agg := fmt.Sprintf("%s#%s|", key, value)
    fmt.Println("encrypting: " + agg)
    var w string
if (on) {
      w = EncryptAES([]byte(cipher), agg)
                } else {
      w = agg
    }

  	file.WriteString(w)
	}
}

func readDecrypt() {
	v, _ := ioutil.ReadFile("data.txt")
	z := string(v)
  fmt.Println("recovered from file: " + z)
  var zz string
if (on) {
  zz = DecryptAES([]byte(cipher), z)
  } else {
  zz =z
  }
  fmt.Println("decrypted: " + zz)
	w := strings.Split(zz, "|")
  fmt.Println(len(w))
	for _, x := range w {
	  y := strings.Split(x, "#")
    if (len(y) == 2) {
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

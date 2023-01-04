package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// cache map
var data = make(map[string]string)

var enckey = []byte("this's secret key.enough 32 bits")

var keyStr = hex.EncodeToString(enckey)
var encryptCache = true
var logToFile = false

func main() {

	defer func() {
		if err := recover(); err != nil {
			log.Println("panic ojccurred:", err)
		}
	}()

	if logToFile {
		file, _ := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		log.SetOutput(file)
	}
	insert("key1", "value1111")
	insert("key2", "value2222")
	insert("key3", "value3333")
	writeEncrypt()
	erase()
	readDecrypt()

}

func insert(key, value string) {
	s := fmt.Sprintf("inserting %s, %s", key, value)
	log.Println(s)
	data[key] = value
}

func deleteByKey(key string) {
	_, exists := data[key]
	if exists {
		delete(data, key)
	}
}

func display() {
	log.Println("display cache: ")
	for key, name := range data {
		log.Println(key + " " + name)
	}
}

func erase() {
	log.Println("erasing...")
	for k := range data {
		delete(data, k)
	}
}

func writeEncrypt() {
	file, _ := os.Create("data.txt")
	defer file.Close()
	agg := ""
	for key, value := range data {
		agg += fmt.Sprintf("%s#%s|", key, value)
	}
	log.Println("encrypting: " + agg)
	var w string
	if encryptCache {
		w = encrypt(keyStr, agg)
	} else {
		w = agg
	}

	file.WriteString(w)
}

func readDecrypt() {
	v, _ := ioutil.ReadFile("data.txt")
	z := string(v)
	log.Println("recovered from file: " + z)
	var zz string
	if encryptCache {
		zz = decrypt(keyStr, z)
	} else {
		zz = z
	}
	log.Println("decrypted: " + zz)
	w := strings.Split(zz, "|")
	log.Println("reloading...")
	for _, x := range w {
		y := strings.Split(x, "#")
		if len(y) == 2 {
			k := y[0]
			v := y[1]
			insert(k, v)
		}
	}
}

func encrypt(keyString string, stringToEncrypt string) (encryptedString string) {
	// convert key to bytes
	key, _ := hex.DecodeString(keyString)
	plaintext := []byte(stringToEncrypt)

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	// convert to base64
	return base64.URLEncoding.EncodeToString(ciphertext)
}

// decrypt from base64 to decrypted string
func decrypt(keyString string, stringToDecrypt string) string {
	key, _ := hex.DecodeString(keyString)
	ciphertext, _ := base64.URLEncoding.DecodeString(stringToDecrypt)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)

	return fmt.Sprintf("%s", ciphertext)
}

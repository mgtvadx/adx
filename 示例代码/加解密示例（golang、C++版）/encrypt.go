package encrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
)

var paddingErr error = errors.New("Error Padding")

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte, blockSize int) ([]byte, error) {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	if unpadding > 0 && unpadding <= blockSize {
		return origData[:(length - unpadding)], nil
	}
	return origData[:], paddingErr
}

func AesCBCEncrypte(s []byte, key string) string {
	keybytes := []byte(key)
	plaintext := PKCS5Padding([]byte(s), aes.BlockSize)
	block, _ := aes.NewCipher(keybytes[:aes.BlockSize])

	mode := cipher.NewCBCEncrypter(block, keybytes[:aes.BlockSize])
	crypted := make([]byte, len(plaintext))
	mode.CryptBlocks(crypted, plaintext)
	return string(hex.EncodeToString(crypted))
}

func AesCBCDecrypte(decrypte string, key string) (string, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Fprintf(os.Stderr, "error string:%s key:%s err:%v\n", decrypte, key, err)
		}
	}()

	if (len(decrypte) % (aes.BlockSize * 2)) != 0 {
		return "", errors.New("decrypte data too short")
	}
	keybytes := []byte(key)
	decrypteData, err := hex.DecodeString(decrypte)
	if err != nil {
		return "", errors.New("decrypte data format error")
	}

	block, _ := aes.NewCipher(keybytes[:aes.BlockSize])
	mode := cipher.NewCBCDecrypter(block, keybytes[:aes.BlockSize])

	decryptedData := make([]byte, len(decrypteData))
	mode.CryptBlocks(decryptedData, decrypteData)
	decryptedData, err = PKCS5UnPadding(decryptedData, aes.BlockSize)
	if err != nil {
		return "", err
	}
	return string(decryptedData), nil
}

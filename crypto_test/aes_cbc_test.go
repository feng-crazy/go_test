package crypto_test

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"testing"
)

//
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// AesEncrypt 加密函数
func AesEncrypt(plaintext []byte, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	plaintext = PKCS7Padding(plaintext, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(plaintext))
	blockMode.CryptBlocks(crypted, plaintext)
	return crypted, nil
}

// AesDecrypt 解密函数
func AesDecrypt(ciphertext []byte, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, iv[:blockSize])
	origData := make([]byte, len(ciphertext))
	blockMode.CryptBlocks(origData, ciphertext)
	origData = PKCS7UnPadding(origData)
	return origData, nil
}

func TestAesCBC(t *testing.T) {
	key, _ := hex.DecodeString("a5cc893b0db32f4f48963d516dff25cb")
	plaintext := []byte(`{"User":"system","Verify":"e10adc3949ba59abbe56e057f20f883e"}`)

	// c := make([]byte, aes.BlockSize+len(plaintext))
	// iv := c[:aes.BlockSize]
	iv := []byte("a5cc893b0db32f4f")
	// 加密
	ciphertext, err := AesEncrypt(plaintext, key, iv)
	if err != nil {
		panic(err)
	}

	// 打印加密base64后密码
	fmt.Println(base64.StdEncoding.EncodeToString(ciphertext))

	// 解密
	plaintext, err = AesDecrypt(ciphertext, key, iv)
	if err != nil {
		panic(err)
	}

	// 打印解密明文
	fmt.Println(string(plaintext))

}

func TestAesCBC1(t *testing.T) {
	key := []byte("fimos87ejdusedke")
	plaintext := []byte(`15767634647`)

	// c := make([]byte, aes.BlockSize+len(plaintext))
	// iv := c[:aes.BlockSize]
	iv := []byte("Bce Encryptor IV")
	// 加密
	ciphertext, err := AesEncrypt(plaintext, key, iv)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(ciphertext))

	// 解密
	plaintext, err = AesDecrypt(ciphertext, key, iv)
	if err != nil {
		panic(err)
	}

	// 打印解密明文
	fmt.Println(string(plaintext))

}

func TestAesCBC3(t *testing.T) {
	key := "fimos87ejdusedke"
	cipherText := "012c96c5b51f04282f1d68a99ce4bf293ca53310aacb8b68be4d3b786c2c9390"
	iv := []byte("Bce Encryptor IV")
	//plainTextByte, err := AesDecrypt([]byte(cipherText), []byte(key), iv)
	//if err != nil {
	//	t.Error(err)
	//}
	//t.Log(string(plainTextByte))
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		t.Error(err)
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, iv[:blockSize])
	origData := make([]byte, len(cipherText))
	blockMode.CryptBlocks(origData, []byte(cipherText))
	//origData = PKCS7UnPadding(origData)

	t.Log(string(origData))

}

package crypto_test

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"testing"
)

func TestAesCtr(t *testing.T) {

	// Load your secret key from a safe place and reuse it across multiple
	// NewCipher calls. (Obviously don't use this example key for anything
	// real.) If you want to convert a passphrase to a key, use a suitable
	// package like bcrypt or scrypt.
	key, _ := hex.DecodeString("6368616e676520746869732070617373")

	//key, _ := hex.DecodeString("GROdz4ioKxC590SRWhmSXtwJuMAaOh0p")
	plaintext := []byte("15767634546")

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	// It's important to remember that ciphertexts must be authenticated
	// (i.e. by using crypto/hmac) as well as being encrypted in order to
	// be secure.

	// CTR mode is the same for both encryption and decryption, so we can
	// also decrypt that ciphertext with NewCTR.

	plaintext2 := make([]byte, len(plaintext))
	stream = cipher.NewCTR(block, iv)
	stream.XORKeyStream(plaintext2, ciphertext[aes.BlockSize:])

	fmt.Printf("%s\n", plaintext2)
	// Output: some plaintext
}

func TestAesCtr1(t *testing.T) {
	key := []byte("1234567890123456")
	plaintext := []byte("text can be a random lenght")
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	// The IV needs to be unique,but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	// BTW (only for test purpose) I don't include it
	ciphertext := make([]byte, len(plaintext))
	iv := []byte{'\x0f', '\x0f', '\x0f'}
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(ciphertext, plaintext)
	// CTR mode is the same for both encryption and decryption,so we can
	// also decrypt that ciphertext with NewCTR.
	base := base64.StdEncoding.EncodeToString(ciphertext)
	fmt.Printf("encodedHEX: %x\n", ciphertext)
	fmt.Printf("encodedBASE: %s\n", base)
	plaintext2 := make([]byte, len(plaintext))
	stream = cipher.NewCTR(block, iv)
	stream.XORKeyStream(plaintext2, ciphertext)
	fmt.Printf("decoded: %s\n", plaintext2)
}

func AesCTREncrypt(plainText, key []byte) []byte {
	//判断用户传过来的key是否符合16字节，如果不符合16字节加以处理
	keylen := len(key)
	if keylen == 0 { //如果用户传入的密钥为空那么就用默认密钥
		key = []byte("wumansgygoaescbc") //默认密钥
	} else if keylen > 0 && keylen < 16 { //如果密钥长度在0到16之间，那么用0补齐剩余的
		key = append(key, bytes.Repeat([]byte{0}, 16-keylen)...)
	} else if keylen > 16 {
		key = key[:16]
	}
	//1.指定使用的加密aes算法
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	//2.不需要填充,直接获取ctr分组模式的stream
	// 返回一个计数器模式的、底层采用block生成key流的Stream接口，初始向量iv的长度必须等于block的块尺寸。
	iv := []byte("Bce Encryptor IV")
	stream := cipher.NewCTR(block, iv)

	//3.加密操作
	cipherText := make([]byte, len(plainText))
	stream.XORKeyStream(cipherText, plainText)

	return cipherText
}

func AesCTRDecrypt(cipherText, key []byte) []byte {
	//判断用户传过来的key是否符合16字节，如果不符合16字节加以处理
	keylen := len(key)
	if keylen == 0 { //如果用户传入的密钥为空那么就用默认密钥
		key = []byte("fimos87ejdusedke") //默认密钥
	} else if keylen > 0 && keylen < 16 { //如果密钥长度在0到16之间，那么用0补齐剩余的
		key = append(key, bytes.Repeat([]byte{0}, 16-keylen)...)
	} else if keylen > 16 {
		key = key[:16]
	}
	//1.指定算法:aes
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	//2.返回一个计数器模式的、底层采用block生成key流的Stream接口，初始向量iv的长度必须等于block的块尺寸。
	iv := []byte("Bce Encryptor IV")
	stream := cipher.NewCTR(block, iv)

	//3.解密操作
	plainText := make([]byte, len(cipherText))
	stream.XORKeyStream(plainText, cipherText)

	return plainText
}

func TestAesCtr2(t *testing.T) {
	plainText := "15767634546"
	key := "fimos87ejdusedke"
	cipherText := AesCTREncrypt([]byte(plainText), []byte(key))
	t.Log(string(cipherText))
	plainTextByte := AesCTRDecrypt([]byte(cipherText), []byte(key))
	t.Log(string(plainTextByte))
}

func TestAesCtr3(t *testing.T) {
	key := "fimos87ejdusedke"
	cipherText := "012c96c5b51f04282f1d68a99ce4bf293ca53310aacb8b68be4d3b786c2c9390"
	cipher, err := hex.DecodeString(cipherText)
	if err != nil {
		t.Error(err)
	}
	plainTextByte := AesCTRDecrypt(cipher, []byte(key))
	t.Log(string(plainTextByte))
}

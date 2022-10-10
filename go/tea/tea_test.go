package tea

import (
	"encoding/hex"
	"testing"
)

const k = "1234657890abcdef" // 密钥, 长度必须为16byte

func TestTeaEncode(t *testing.T) {
	c := NewTeaCipher([]byte(k))

	//加密字符串
	str := "hello qq tea go"

	result := c.Encrypt([]byte(str))
	t.Log("加密后的2进制字符串:", result)

	encodedStr := hex.EncodeToString(result)
	t.Log("加密后的16进制字符串:", encodedStr)

	result = c.Decrypt(result)
	t.Log("解密后的字符串:", string(result))
}

func TestTeaDecode(t *testing.T) {
	c := NewTeaCipher([]byte(k))

	str2 := "1d9f0fb8464a1b6e7072cbd4c09c4efb4cae8437562d771beb914fa1faafcd6e"
	result, _ := hex.DecodeString(str2)
	result = c.Decrypt(result)
	t.Log("解密后的字符串:", string(result))
}

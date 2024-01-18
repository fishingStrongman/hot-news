package tools

import (
	"crypto/sha256"
	"encoding/hex"
)

func Sha256Hash(input string) string {
	// 创建一个新的 SHA-256 散列器
	hasher := sha256.New()

	// 将字符串转换为字节数组并写入散列器
	hasher.Write([]byte(input))

	// 计算散列值并转换为十六进制字符串
	hashInBytes := hasher.Sum(nil)
	hashString := hex.EncodeToString(hashInBytes)

	return hashString
}

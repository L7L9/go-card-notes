package utils

import "golang.org/x/crypto/bcrypt"

// HashEncrypt //
// 哈希加密
func HashEncrypt(data string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(data), bcrypt.DefaultCost)
	return string(bytes)
}

// CompanyHash //
// 比较数据加密后的hash是否一样
func CompanyHash(data string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(data))
	return err == nil
}

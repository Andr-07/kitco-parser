package utils

import "crypto/md5"

func HashMD5(s string) []byte {
	hash := md5.Sum([]byte(s))
	return hash[:]
}
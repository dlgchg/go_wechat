package util

import (
	"crypto/sha1"
	"fmt"
	"io"
	"sort"
)

//Signature sha1签名
func Signature(params ...string) string {
	sort.Strings(params)
	hash := sha1.New()
	for _, v := range params {
		io.WriteString(hash, v)
	}
	return fmt.Sprintf("%x", hash.Sum(nil))
}

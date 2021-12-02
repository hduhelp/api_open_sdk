package utils

import (
	"crypto/sha1"
	"fmt"
)

func Sha1Encode(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}

/**
 *通用实验
 */

package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func main() {
	fmt.Print(sessionId())
}

func sessionId() string {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

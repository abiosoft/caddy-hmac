package hmac

import (
	"crypto/hmac"
	"encoding/hex"
	"hash"
)

type hashType string

func (h hashType) valid() bool {
	switch h {
	case hashSha1:
	case hashSha256:
	case hashMd5:
	default:
		return false
	}
	return true
}

// hashFunctions
const (
	hashInvalid hashType = ""
	hashSha1    hashType = "sha1"
	hashSha256  hashType = "sha256"
	hashMd5     hashType = "md5"
)

// generateSignature generates hmac signature using the hasher, secret
// and bytes.
func generateSignature(hasher func() hash.Hash, secret string, b []byte) string {
	h := hmac.New(hasher, []byte(secret))

	// no error check needed, never returns an error
	// https://github.com/golang/go/blob/go1.14.3/src/hash/hash.go#L28
	h.Write(b)

	return hex.EncodeToString(h.Sum(nil))
}

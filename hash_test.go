package hmac

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"testing"
)

func TestGenerateSignature(t *testing.T) {
	type algos struct {
		sha1   string
		sha256 string
		md5    string
	}
	tests := []struct {
		body     string
		secret   string
		expected algos
	}{
		{
			body:   "something",
			secret: "secret",
			expected: algos{
				sha1:   "889b4bc00e8a5d6b05c3cb47db58217cac788526",
				sha256: "f28635b332e79bbb9322cebf268dc2bb7f13c90e52310d8ecb0a56bab8a8b9b7",
				md5:    "a54c3399cb671fe6e4abecf2b078e3ad",
			},
		},
		{
			body:   "something else",
			secret: "secrets",
			expected: algos{
				sha1:   "cd936c4f51ecec358b62846b9c543e6c0bf3e1d7",
				sha256: "9bd480144621173f8b045224ad409aa2a8e8213a7723acdf040868158496e409",
				md5:    "49c6a78f69fd4b317d04ba21abe96ac4",
			},
		},
		{
			body:   `Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum`,
			secret: "C%8@'ikedHT.r@0",
			expected: algos{
				sha1:   "22094f7f74b9424d038567e05058b989ad49fc72",
				sha256: "56923d77e8a4053ac3709150238394f157d68364e05b7daa16455e705f87ccf0",
				md5:    "a85ec23c3b5f82e67ccf15341a82d09f",
			},
		},
	}

	// sha1
	for i, tt := range tests {
		t.Run(fmt.Sprintf("sha1 %v", i), func(t *testing.T) {
			signature := generateSignature(sha1.New, tt.secret, []byte(tt.body))
			if signature != tt.expected.sha1 {
				t.Errorf("got: %v, want: %v", signature, tt.expected.sha1)
			}
		})
	}

	// sha256
	for i, tt := range tests {
		t.Run(fmt.Sprintf("sha256 %v", i), func(t *testing.T) {
			signature := generateSignature(sha256.New, tt.secret, []byte(tt.body))
			if signature != tt.expected.sha256 {
				t.Errorf("got: %v, want: %v", signature, tt.expected.sha256)
			}
		})
	}

	// md5
	for i, tt := range tests {
		t.Run(fmt.Sprintf("md5 %v", i), func(t *testing.T) {
			signature := generateSignature(md5.New, tt.secret, []byte(tt.body))
			if signature != tt.expected.md5 {
				t.Errorf("got: %v, want: %v", signature, tt.expected.md5)
			}
		})
	}

}

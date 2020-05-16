package hmac

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestCopyRequestBody(t *testing.T) {
	tests := []string{
		"lorem ipsun",
		"8b3608e39f0ec2ddbf69467f92e65faf4c2f2c33",
		"some random strings that will be bytes",
	}
	for i, tt := range tests {
		r := &http.Request{
			Body: ioutil.NopCloser(strings.NewReader(tt)),
		}
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			b, err := copyRequestBody(r)
			if err != nil {
				t.Fatal(err)
			}
			if s := string(b); s != tt {
				t.Errorf("got: %v, want: %v", s, tt)
			}
		})
	}
}

package hmac

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
)

// ServeHTTP implements caddyhttp.MiddlewareHandler.
func (m HMAC) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {
	if r.Body == nil {
		// nothing to do
		return next.ServeHTTP(w, r)
	}
	body, err := copyRequestBody(r)
	if err != nil {
		return err
	}

	repl := r.Context().Value(caddy.ReplacerCtxKey).(*caddy.Replacer)

	signature := generateSignature(m.hasher, m.Secret, body)
	if err != nil {
		return err
	}

	repl.Set(m.replacerKey(), signature)
	return next.ServeHTTP(w, r)
}

// copyRequestBody copies the request body while making it reusable.
// It returns the copied []byte.
func copyRequestBody(r *http.Request) ([]byte, error) {
	// drain the body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	// make a copy
	bodyCopy := make([]byte, len(body))
	copy(bodyCopy, body)

	// replace the body
	r.Body = ioutil.NopCloser(bytes.NewReader(body))

	// return the copy
	return bodyCopy, nil
}

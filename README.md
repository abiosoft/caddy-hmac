# caddy-hmac

Caddy v2 hmac middleware

## Usage

`hmac` computes the hash of the request body as a `{hmac.signature}` [placeholder](https://caddyserver.com/docs/caddyfile/concepts#placeholders) for other [matchers](https://caddyserver.com/docs/caddyfile/matchers) and [handlers](https://caddyserver.com/docs/caddyfile/directives).

### Caddyfile

```
hmac [<name>] <algorithm> <secret>
```

* **name** - [optional] if set, names the signature. It can then be referenced with `{hmac.name.signature}` instead.
* **algorithm** - hash algorithm to use. Can be one of `sha1`, `sha256`, `md5`.
* **secret** - the hmac secret key

#### Example

Validate a Github webhook secured with a secret.

```
@github {
    header_regexp X-Hub-Signature "[a-z0-9]+\=([a-z0-9]+)"
}
@hmac {
    expression {hmac.signature} == {http.regexp.1}
}
route @github {
    hmac sha1 {$GITHUB_WEBHOOK_SECRET}
    ...
}
```

### JSON

`hmac` can be part of any route as an handler

```jsonc
{
  ...
  "routes": [
    {
      "handle": [
        {
          // required to indicate the handler
          "handler": "hmac",
          // [optional] if set, names the sigurature to be referenced
          // as {hmac.name.signature} instead.
          "name": "",
          // the algorithm to use. can be sha1, sha256, md5
          "algorithm": "sha1",
          // hmac secret
          "secret": "some secrets"
        }
      ]
    },
    ...
  ]
  ...
}
```

## License

Apache 2
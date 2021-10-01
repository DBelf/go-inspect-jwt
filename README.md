This application can be used to inspect jwt in a similar fashion as done on [jwt.io](www.jwt.io).

You can use this application in the following way:

```bash
$ inspectjwt -t <token>
```

---
Example:
```bash
$ inspectjwt -t eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c
{
  "alg": "HS256",
  "typ": "JWT"
}{
  "iat": 1516239022,
  "name": "John Doe",
  "sub": "1234567890"
}
```

The `-t` parameter is mandatory.

The `-exp` parameter can be added to also determine whether the token is still valid. 

The inpsected token will use colors to distinguish the token header and token claims.

---
Current limitations:
* Token signature cannot be validated
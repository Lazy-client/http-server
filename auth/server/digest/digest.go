package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"strings"
)

var users = map[string]string{
	"admin": "admin", // admin:admin
}

func main() {
	r := http.NewServeMux()
	r.HandleFunc("/", digestHandler)
	r.HandleFunc("/hmac", func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		// content=hello secret=123456 algorithm=sha256
		if auth != "ac28d602c767424d0c809edebf73828bed5ce99ce1556f4df8e223faeec60edd" {
			w.WriteHeader(http.StatusUnauthorized)
		}
	})
	http.ListenAndServe(":8082", r)
}

func digestHandler(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		w.Header().Set("WWW-Authenticate", `Digest realm="example", nonce="1234567890", algorithm="MD5", qop="auth-int"`)
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintln(w, "Authentication required")
		return
	}
	authInfo := parseAuthHeader(authHeader, r.Method)
	if authInfo == nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalid Authorization header")
		return
	}
	password, ok := users[authInfo.username]
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintln(w, "Invalid username")
		return
	}

	var payload []byte
	if r.Body != nil {
		payload, _ = io.ReadAll(r.Body)
	}
	hbody := hash(string(payload))
	ha1 := hash(fmt.Sprintf("%s:%s:%s", authInfo.username, authInfo.realm, password))
	ha2 := hash(fmt.Sprintf("%s:%s:%s", authInfo.method, authInfo.uri, hbody))
	response := hash(fmt.Sprintf("%s:%s:%s:%s:%s:%s", ha1, authInfo.nonce, authInfo.nc, authInfo.cnonce, authInfo.qop, ha2))
	if authInfo.response != response {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintln(w, "Invalid password")
		return
	}
	fmt.Fprintln(w, "Hello, "+authInfo.username+"!")
}

func hash(data string) string {
	sum := md5.Sum([]byte(data))
	return hex.EncodeToString(sum[:])
}

type authInfo struct {
	username string
	realm    string
	nonce    string
	uri      string
	method   string
	qop      string
	nc       string
	cnonce   string
	response string
}

func parseAuthHeader(header, method string) *authInfo {
	if !strings.HasPrefix(header, "Digest ") {
		return nil
	}
	header = strings.TrimPrefix(header, "Digest ")
	parts := strings.Split(header, ", ")
	info := &authInfo{}
	for _, part := range parts {
		kv := strings.SplitN(part, "=", 2)
		if len(kv) != 2 {
			return nil
		}
		key := strings.Trim(kv[0], "\"")
		value := strings.Trim(kv[1], "\"")
		switch key {
		case "username":
			info.username = value
		case "realm":
			info.realm = value
		case "nonce":
			info.nonce = value
		case "uri":
			info.uri = value
		case "algorithm":
			// ignore
		case "qop":
			info.qop = value
		case "nc":
			info.nc = value
		case "cnonce":
			info.cnonce = value
		case "response":
			info.response = value
		default:
			return nil
		}
	}
	if info.username == "" || info.realm == "" || info.nonce == "" || info.uri == "" || info.qop == "" || info.nc == "" || info.cnonce == "" || info.response == "" {
		return nil
	}
	info.method = strings.ToUpper(method)
	return info
}

package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var SigningKey = []byte("secret_key")

func main() {
	http.HandleFunc("/token", loginForToken)
	http.HandleFunc("/protected", protectedHandler)
	http.ListenAndServe("0.0.0.0:8080", nil)
}

func loginForToken(w http.ResponseWriter, r *http.Request) {
	// 验证用户名和密码
	username := r.FormValue("username")
	password := r.FormValue("password")
	if username == "admin" && password == "123456" {
		// 颁发Token
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["username"] = username
		claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
		tokenString, err := token.SignedString(SigningKey)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, "Error while signing the token")
			return
		}
		w.Write([]byte(tokenString))
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintln(w, "Invalid credentials")
	}
}

const (
	BearerToken = "Bearer "
)

const (
	HeaderAuthorization = "Authorization"
	HeaderAuthenticate  = "WWW-Authenticate"
)

func protectedHandler(w http.ResponseWriter, r *http.Request) {
	// 从请求头中获取Token
	auth := r.Header.Get("Authorization")
	if !strings.HasPrefix(auth, BearerToken) {
		w.Header().Set(http.CanonicalHeaderKey(HeaderAuthenticate), `Bearer realm="127.0.0.1", error="invalid_token"`)
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintln(w, "Token not found")
		return
	}
	auth = strings.TrimPrefix(auth, BearerToken)
	// 解析Token
	token, err := jwt.Parse(auth, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return SigningKey, nil
	})
	if err != nil {
		w.Header().Set(http.CanonicalHeaderKey(HeaderAuthenticate), `Bearer realm="127.0.0.1", error="invalid_token"`)
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintln(w, "Invalid token")
		return
	}
	// 检查Token的有效性
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username := claims["username"].(string)
		exp := time.Unix(int64(claims["exp"].(float64)), 0)
		if time.Now().After(exp) {
			w.Header().Set(http.CanonicalHeaderKey(HeaderAuthenticate), `Bearer realm="127.0.0.1", error="invalid_token", error_description="The access token expired"`)
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintln(w, "Token has expired")
			return
		}
		fmt.Fprintf(w, "Hello %s!", username)
		w.WriteHeader(http.StatusOK)
	}
}

package main

import (
	"encoding/base64"
	"fmt"
	"regexp"
	"time"
)

func main() {
	var x time.Duration
	fmt.Println(x == 0)
	//msg := DecodeB64([]byte(base64.StdEncoding.EncodeToString([]byte("hello world\xFF"))))
	//fmt.Println(msg)
	//fmt.Println(EncodeB64("hello world\xFF"))
	//filename, _ := url.PathUnescape("%E5%B9%BF%E5%91%8A%E4%B8%BB%E4%BD%93%E5%BA%93%E5%BA%97%E9%93%BA%E6%9F%A5%E8%AF%A220231205151420.xlsx")
	//println(filename)

}

func isBase64(in string) bool {
	// 定义正则表达式模式
	pattern := `^([A-Za-z0-9+/]{4})*([A-Za-z0-9+/]{3}=|[A-Za-z0-9+/]{2}==)?$`
	return regexp.MustCompile(pattern).MatchString(in)
}

func EncodeB64(message string) (retour string) {
	base64Text := make([]byte, base64.StdEncoding.EncodedLen(len(message)))
	base64.StdEncoding.Encode(base64Text, []byte(message))
	return string(base64Text)
}

func DecodeB64(message []byte) []byte {
	base64Text := make([]byte, base64.StdEncoding.DecodedLen(len(message)))
	base64.StdEncoding.Decode(base64Text, message)
	return base64Text
}

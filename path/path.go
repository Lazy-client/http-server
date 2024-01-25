package main

import (
	"fmt"
	"net/url"
	"sort"
	"strings"
)

func main() {

	fmt.Println(url.PathEscape("%&#$%^&*(){}[]%20"))
	fmt.Println(url.QueryEscape("%&#$%^&*(){}[]%20"))

	uu := &url.URL{
		Path: "x- +9#/你好/y/\t/m%",
	}
	encodedPath := uu.String()
	fmt.Println(encodedPath)
	pathUnescape, err := url.PathUnescape(uu.Path)

	fmt.Println(pathUnescape)
	uri, _ := EncodeURI("/x- +9#/你好/y/\t/m?t=jjj\t\tnn&x=9&ii\t=\n&kk=你 好%2B?")
	u := "http://www.baidu.com" + uri
	baseUrl, _ := url.Parse(u)
	// 合并接口设置的参数以及url的查询参数，
	params, err := ParseQuery(baseUrl.RawQuery)
	params.Add("888", "8888 8888+++")
	decodedQuery, err := url.PathUnescape(EncodeQuery(params))
	if err != nil {
		return
	}
	fmt.Println(decodedQuery)

}

func EncodeURI(uri string) (string, error) {
	// 分隔出路径和查询参数
	parts := strings.SplitN(uri, "?", 2)
	path := parts[0]
	var query string
	if len(parts) > 1 {
		query = parts[1]
	}
	q, err := ParseQuery(query)
	if err != nil {
		return "", err
	}
	encodedQuery := EncodeQuery(q)
	// 构建编码后的URL
	u := &url.URL{
		Path:     path,
		RawQuery: encodedQuery,
	}
	// 返回编码后的URI
	return u.String(), nil
}

func EncodeQuery(v url.Values) string {
	if v == nil {
		return ""
	}
	var buf strings.Builder
	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		vs := v[k]
		keyEscaped := url.PathEscape(k)
		for _, v := range vs {
			if buf.Len() > 0 {
				buf.WriteByte('&')
			}
			buf.WriteString(keyEscaped)
			buf.WriteByte('=')
			buf.WriteString(url.PathEscape(v))
		}
	}
	return buf.String()
}

func ParseQuery(query string) (m url.Values, err error) {
	m = make(url.Values)
	for query != "" {
		var key string
		key, query, _ = strings.Cut(query, "&")
		if strings.Contains(key, ";") {
			err = fmt.Errorf("invalid semicolon separator in query")
			return
		}
		if key == "" {
			continue
		}
		key, value, _ := strings.Cut(key, "=")
		key, err = url.PathUnescape(key)
		if err != nil {
			return
		}
		value, err = url.PathUnescape(value)
		if err != nil {
			return
		}
		m[key] = append(m[key], value)
	}
	return
}

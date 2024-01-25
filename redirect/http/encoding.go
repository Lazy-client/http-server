package http

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/andybalholm/brotli"
	"github.com/klauspost/compress/flate"
	"github.com/klauspost/compress/gzip"
)

func DecodeResponse(resp *http.Response) (*http.Response, error) {
	if resp == nil {
		return nil, fmt.Errorf("invliad http resp value: nil")
	}
	unCompress := false

	// 获取Content-Encoding头部值并按顺序解码
	encodings := strings.Split(resp.Header.Get("Content-Encoding"), ",")
	for i := len(encodings) - 1; i >= 0; i-- {
		encoding := strings.TrimSpace(encodings[i])

		switch encoding {
		case "gzip":
			reader, err := gzip.NewReader(resp.Body)
			if err != nil {
				return nil, fmt.Errorf("gzip解码失败：%v", err)
			}

			resp.Body = reader
			unCompress = true

		case "deflate":
			reader := flate.NewReader(resp.Body)
			resp.Body = reader
			unCompress = true

		case "br":
			reader := io.NopCloser(brotli.NewReader(resp.Body))
			resp.Body = reader
			unCompress = true
		}
	}

	if unCompress {
		// 这里防止客户端进行2次解码，所以要移除此头部。保留了此头部原有值，仅是为了和postman的表现一致，无其它含义。
		// resp.Header.Del("Content-Encoding")
		// 移除此头部是因为，解码后body实际长度已经发生变化，通过设置Content-Length为-1表示Content-Length无效
		resp.Header.Del("Content-Length")
		resp.ContentLength = -1
	}
	return resp, nil
}

package main

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	//encode "awesomeProject/redirect/http"
)

func main() {
	url := "https://14.119.104.189"
	method := http.MethodGet

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("222", "222")
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	bs, err := io.ReadAll(req.Body)
	req.Body = io.NopCloser(bytes.NewBufferString(string(bs)))
	//var configCheckRedirect func(req *http.Request, via []*http.Request) error
	//configCheckRedirect = func(req *http.Request, via []*http.Request) error {
	//	// 判断重定向次数是否超过限制
	//	if len(via) >= 10 {
	//		return errors.New("stopped after 10 redirects")
	//	}
	//
	//	//prevReq := via[len(via)-1]
	//	//req.Method = prevReq.Method
	//	// 复制前一个请求的GetBody方法到新的请求中
	//	req.Body = io.NopCloser(bytes.NewBufferString(string(bs)))
	//	// 自定义重定向逻辑
	//	// ...
	//	return nil
	//}
	client := &http.Client{
		//CheckRedirect: configCheckRedirect,
	}

	resp, err := client.Do(req)

	//resp, err = encode.DecodeResponse(resp)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resp.StatusCode)
}

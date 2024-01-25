package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {

	body := `{"Action":"create_order","MidasAppId":"1450008583","Num":1,"OpenId":"123321","OutTradeNo":"12333121","Pf":"pc","ProductDetail":"122121","ProductId":"12345","ProductName":"123","UnitPrice":1}`
	//req, err := http.NewRequest(http.MethodPost, "http://9.138.93.44/api/common/openOrder", bytes.NewBufferString(body))

	req, err := http.NewRequest(http.MethodGet, "http://127.0.0.1/time", nil)
	req.Header.Set("X-Tdea-Version", "2023122501")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Tdea-Timestamp", "1704793820")
	//req.Header.Set("Cookie", "x-client-ssid=4f11630a:018cd423c2d1:1bb85c; x_host_key_access=5dc9f6f6110f6a7902315231bb99474460620549_s")
	auth := "TDEA-HMAC-SHA256 Credential=1582119908462/2024-01-09/1450008583/tdea_request, SignedHeaders=content-type;host, Signature=dd2926d4fff594c925e4e480cbad3db5ff0867e290ccb31c2cb8e8e8a1c2be50"
	req.Header.Set("Authorization", auth)
	//req.Header.Set("Accept-Encoding", "gzip")

	req.Host = "9.138.93.44"

	now := time.Now()
	resp, err := (&http.Client{
		Timeout: 2 * time.Second,
	}).Do(req)

	fmt.Println(time.Since(now).Seconds())
	// 读取解压缩后的数据
	var data []byte
	if resp != nil {
		data, err = io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}

	//zipBuffer, err := gzip.NewReader(resp.Body)
	//originData, err := io.ReadAll(zipBuffer)

	_ = err
	_ = body
	fmt.Println(string(data))
	//fmt.Println(string(originData))
}

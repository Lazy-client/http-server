package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync/atomic"
	"time"

	"github.com/klauspost/compress/gzip"
)

func main() {
	// 设置logrus的时间格式为自定义格式
	log.Println("hello word")
	redirect302()
	cnt()
	callback()
	parseQuery()
	respFile()
	hmac()
	respTextNoUtf8()
	respJsonNoUtf8()
	http.HandleFunc("/form", handleFormData)
	respTxt()
	respExcel()
	respTextB64()
	chunked()
	timeout()
	// 启动HTTP服务
	err := http.ListenAndServe("0.0.0.0:80", nil)
	if err != nil {
		log.Println("HTTP server error:", err)
	}
}
func parseQuery() {
	http.HandleFunc("/ /你好/你好/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("raw-query", r.URL.RawQuery)
		log.Println("name", r.URL.Query().Get("name"))
		log.Println("path", r.URL.Path)
	})
}

var loadTestCnt int64

func atomicAdd() {
	atomic.AddInt64(&loadTestCnt, 1)
}
func callback() {
	http.HandleFunc("/event-hook/send", func(w http.ResponseWriter, r *http.Request) {
		log.Println("/event-hook/send接口被调用")
		if r.Body != nil {
			payload, err := io.ReadAll(r.Body)
			if err != nil {
				log.Println(err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			log.Println("body=", string(payload))
		}
		log.Println("headers=", r.Header)
		log.Println("Host", r.Header.Get("Host"))
		log.Println("req.Host", r.Host)
		atomicAdd()
		log.Println("=====请求处理完毕======")
	})
}

func cnt() {
	http.HandleFunc("/cnt", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "访问次数", loadTestCnt)
	})
}

// a->b->c 302 并日志打印body和query参数
func redirect302() {
	// 定义接口a的处理函数
	http.HandleFunc("/a", func(w http.ResponseWriter, r *http.Request) {
		// 读取请求的body参数
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		// 打印body参数
		log.Println("body参数 a:", string(body))

		log.Println("query参数 a:", r.URL.Query())

		// 重定向到接口b
		http.Redirect(w, r, "/b", http.StatusFound)
	})

	// 定义接口b的处理函数
	http.HandleFunc("/b", func(w http.ResponseWriter, r *http.Request) {
		// 读取请求的body参数
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		// 打印body参数ƒ
		log.Println("body参数 b:", string(body))

		log.Println("query参数 b:", r.URL.Query())

		http.Redirect(w, r, "/c", http.StatusFound)
	})

	http.HandleFunc("/c", func(w http.ResponseWriter, r *http.Request) {
		// 读取请求的body参数
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusFound)
			return
		}
		defer r.Body.Close()

		// 打印body参数
		log.Println("body参数 c:", string(body))

		log.Println("query参数 c:", r.URL.Query())

		fmt.Fprintln(w, "a->b->c")
	})
}

func respFile() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 读取二进制文件
		data, err := os.ReadFile("test.png")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// 设置响应头
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename*=UTF-8''%E4%BD%A0%E5%A5%BD.png")
		w.Header().Set("Content-Length", fmt.Sprintf("%d", len(data)))

		// 写入响应体
		if _, err := w.Write(data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	})
}
func respTxt() {
	http.HandleFunc("/txt", func(w http.ResponseWriter, r *http.Request) {
		// 读取二进制文件
		data, err := os.ReadFile("hello.txt")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// 设置响应头
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename*=UTF-8''%E4%BD%A0%E5%A5%BD.txt")
		w.Header().Set("Content-Length", fmt.Sprintf("%d", len(data)))

		// 写入响应体
		if _, err := w.Write(data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	})
}
func respExcel() {
	http.HandleFunc("/excel", func(w http.ResponseWriter, r *http.Request) {
		// 读取二进制文件
		data, err := os.ReadFile("广告主体库店铺查询20231206094538.xlsx")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// 设置响应头
		w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
		w.Header().Set("Content-Disposition", "attachment; filename*=UTF-8''%E4%BD%A0%E5%A5%BD.xlsx")
		w.Header().Set("Content-Length", fmt.Sprintf("%d", len(data)))

		// 写入响应体
		if _, err := w.Write(data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	})
}

func respTextNoUtf8() {
	http.HandleFunc("/text", func(w http.ResponseWriter, r *http.Request) {
		// 设置响应头
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		// 写入响应体
		if _, err := w.Write([]byte("\xFF \xFF")); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}
func respTextB64() {
	http.HandleFunc("/text-b64", func(w http.ResponseWriter, r *http.Request) {
		// 设置响应头
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		// 写入响应体
		if _, err := w.Write([]byte("aGVsbG8gd29ybGT/")); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}

func respJsonNoUtf8() {
	http.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request) {
		// 设置响应头
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		// 写入响应体
		if _, err := w.Write([]byte(`{"code":"\\xFF"}`)); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}

func hmac() {
	http.HandleFunc("/hmac", func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		// content=hello secret=123456 algorithm=sha256
		if auth != "ac28d602c767424d0c809edebf73828bed5ce99ce1556f4df8e223faeec60edd" {
			w.WriteHeader(http.StatusUnauthorized)
		}
	})
}
func timeout() {
	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(15 * time.Minute)
		w.WriteHeader(http.StatusOK)
	})
}

func chunked() {
	http.HandleFunc("/chunked-ok", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Connection", "Keep-Alive")
		w.Header().Set("Transfer-Encoding", "chunked")
		w.Header().Set("Content-Encoding", "gzip")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(http.StatusOK)

		data := []byte(`Hello Chunked`)
		var compressedData bytes.Buffer
		writer := gzip.NewWriter(&compressedData)
		writer.Write(data)
		writer.Close()

		// 写入内容编码后的数据 gzip后的，设置chunked自动在respBody中以分块的格式写入
		w.Write(compressedData.Bytes())
	})

	http.HandleFunc("/chunked-nok", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Connection", "Keep-Alive")
		w.Header().Set("Transfer-Encoding", "chunked")
		w.Header().Set("Content-Encoding", "gzip")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(http.StatusOK)
		data := []byte(`Hello Chunked`)
		// 响应头gzip，内容并没有gzip格式
		w.Write(data)
	})
}

func handleFormData(w http.ResponseWriter, r *http.Request) {
	reader, err := r.MultipartReader()
	if err != nil {
		log.Println(err.Error())
		return
	}

	for {
		part, err := reader.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err.Error())
			return
		}
		if part.FileName() != "" {
			file, err := os.Create(part.FileName())
			if err != nil {
				log.Println(err.Error())
				return
			}
			_, err = io.Copy(file, part)
			if err != nil {
				log.Println(err.Error())
				return
			}
			fmt.Fprintf(w, "file %s successfully\n", part.FileName())
		} else {
			var buf bytes.Buffer
			_, err := io.Copy(&buf, part)
			if err != nil {
				log.Println(err.Error())
				return
			}
			fmt.Fprintf(w, "key: %s,value: %s \n", part.FormName(), buf.String())
		}
	}
}

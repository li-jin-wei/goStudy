package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// 处理单个请求的函数
func handleRequest(w http.ResponseWriter, r *http.Request) {
	// 模拟处理请求的耗时操作
	time.Sleep(200 * time.Millisecond)
	fmt.Fprintf(w, "Hello, %s!\n", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handleRequest)

	var wg sync.WaitGroup

	// 启动多个 goroutine 来处理并发请求
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			http.ListenAndServe(":8080", nil)
		}()
	}

	wg.Wait()
}

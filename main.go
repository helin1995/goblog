package main

import (
	"fmt"
	"net/http"
)

func handelerFunc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Hello, 这里是 goblog</h1>")
	} else if r.URL.Path == "/about" {
		fmt.Fprint(w, "此博客使用一记录编程笔记，如您有反馈或建议，请联系"+
			"<a href=\"mailto:summer@example.com\">summer@example.com</a>")
	} else {
		fmt.Fprint(w, "<h1>请求页面未找到：(</h1>"+
			"<p>如有疑惑，请联系我们。</p>")
	}

}

func main() {
	http.HandleFunc("/", handelerFunc)
	http.ListenAndServe(":3000", nil)
}

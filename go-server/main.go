package main // 表示这是一个可以编译成可执行程序的Go代码，

// 导入所需的标准库和外部包
import (
	"fmt"
	"log"
	"net/http"
)

// 处理 "/form" 路径的HTTP请求的函数。它接受两个参数：w类型为http.ResponseWriter和r类型为http.Request。这些参数分别表示响应写入器和服务器接收到的请求
func formHandler(w http.ResponseWriter, r *http.Request) {
	// 解析请求中的表单数据。如果解析过程中出现错误，它会向响应写入器写入错误消息并返回
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful") // 向响应写入器写入一条 "POST 请求成功" 的消息
	// 获取表单数据中 "name" 和 "address" 字段的值，并将这些值写入响应写入器
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

// 处理 "/hello" 路径的HTTP请求的函数。
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// 首先检查请求的URL路径是否为 "/hello"。如果不是，返回一个 "404 找不到页面" 的错误，并停止后续处理
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	// 检查请求的HTTP方法是否为GET。如果不是，它会返回一个 "不支持的请求方法" 的错误，并停止后续处理
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	// 如果两个检查都通过，它会向响应写入器写入 "hello"
	fmt.Fprintf(w, "hello")
}

// 程序的入口点
func main() {
	fileServer := http.FileServer(http.Dir("./static")) // 创建文件服务器，用于从 "./static" 目录提供静态文件
	http.Handle("/", fileServer)                        // 将文件服务器注册到根URL路径 ("/")
	http.HandleFunc("/form", formHandler)               // 处理"/form"的请求
	http.HandleFunc("/hello", helloHandler)             // 处理"/hello"的请求

	fmt.Printf("Starting server at port 8080\n")

	// 使用http.ListenAndServe在端口8080上启动HTTP服务器。如果启动服务器时出现任何错误，它会记录错误并退出程序
	if err := http.ListenAndServe(":8080", nil); err != nil { // nil作为参数表示使用默认的请求多路复用器DefaultServeMux
		log.Fatal(err) // log.Fatal()通常使用在发生严重错误时，不可恢复的情况下使用，例如无法启动服务器、无法连接到数据库等。它通常在主程序中使用，以便将错误信息输出到控制台并停止程序的执行。对于库或模块的函数，通常使用返回错误值的方式来处理错误
	}
}

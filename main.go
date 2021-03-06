package main

import (
    "fmt"
    "net/http"

    "github.com/PuerkitoBio/fetchbot"
)

// メイン処理
func main() {
    f := fetchbot.New(fetchbot.HandlerFunc(handler))
    queue := f.Start()
    queue.SendStringHead("http://google.com", "http://golang.org", "http://golang.org/doc")
    queue.Close()
}

// 実行時処理
func handler(ctx *fetchbot.Context, res *http.Response, err error) {
    if err != nil {
        fmt.Printf("error: %s\n", err)
        return
    }
    fmt.Printf("[%d] %s %s\n", res.StatusCode, ctx.Cmd.Method(), ctx.Cmd.URL())

    proxyUrl, err := url.Parse("http://proxyIp:proxyPort")
myClient := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}
}


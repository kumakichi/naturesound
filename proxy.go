package main

import (
	"flag"
	"fmt"
	"github.com/elazarl/goproxy"
	"net/http"
	"strings"
)

var (
	pport = flag.Int("pport", 8899, "proxy port")
	lport = flag.Int("lport", 1234, "local http port")
)

func localHttpServer() {
	host := fmt.Sprintf(":%d", *lport)
	http.Handle("/", http.FileServer(http.Dir(".")))
	fmt.Printf("Listen local at %s\n", host)
	http.ListenAndServe(host, nil)
}

func proxyServer() {
	localHost := fmt.Sprintf("localhost:%d", *lport)
	serv_host := fmt.Sprintf(":%d", *pport)
	iproxy := goproxy.NewProxyHttpServer()
	iproxy.OnRequest().DoFunc(
		func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
			fmt.Println("URI:", r.RequestURI, r.Method)
			if strings.Contains(r.RequestURI, "naturesoundsfor.me") {
				xx := strings.Replace(r.RequestURI, "www.naturesoundsfor.me", localHost, -1)
				oo := strings.Replace(xx, "naturesoundsfor.me", localHost, -1)
				req, _ := http.NewRequest("GET", oo, nil)
				return req, nil
			}
			return r, nil
		})
	iproxy.Verbose = false
	fmt.Printf("Proxy listen at %s\n", serv_host)
	http.ListenAndServe(serv_host, iproxy)
}

func main() {
	flag.Parse()
	go localHttpServer()
	proxyServer()
}

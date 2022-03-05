package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

type Response struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Version string `json:"version"`
}

func main()  {
	http.HandleFunc("/postHandle", postHandle)
	http.HandleFunc("/healthz", healthz)
	err := http.ListenAndServe(":8891", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func healthz(w http.ResponseWriter, r *http.Request)  {
	w.WriteHeader(200)
	resp := Response{200, "系统可正常调用", ""}
	json.NewEncoder(w).Encode(resp)
}

func postHandle(w http.ResponseWriter, r *http.Request) {
	//获取请求参数
	decoder := json.NewDecoder(r.Body)
	var params map[string]string
	decoder.Decode(&params)
	for k,v := range params {
		fmt.Printf("k===%s,value===%s \n", k, v)
	}

	//header处理
	header := r.Header
	for k,_ := range header {
		if k != "Content-Length" {
			w.Header().Set(k, header.Get(k))
		}
	}

	//返回结果
	returnCode := 200
	resp := Response{returnCode, "调用成功", os.Getenv("VERSION")}
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		fmt.Printf("err is %s \n", err)
	}
	fmt.Printf("hostIp is %s, returnCode is %d \n", getClientIp(r), returnCode)
}

func getClientIp(req *http.Request) string {
	remoteAddr := req.RemoteAddr
	if ip := req.Header.Get("Remote_addr"); ip != "" {
		remoteAddr = ip
	} else {
		remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
	}

	if remoteAddr == "::1" {
		remoteAddr = "127.0.0.1"
	}
	return remoteAddr
}
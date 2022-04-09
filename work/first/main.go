package main

import (
	"context"
	"encoding/json"
	"github.com/golang/glog"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Response struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Version string `json:"version"`
}

func main()  {
	// 处理SIGTERM 信号
	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)

	mux := http.NewServeMux()
	mux.HandleFunc("/postHandle", postHandle)
	mux.HandleFunc("/healthz", healthz)
	srv := &http.Server{Addr: ":8891", Handler: mux}

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	<-stopChan // wait for SIGINT or SIGTERM
	log.Println("Shutting down server...")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	srv.Shutdown(ctx)
	log.Println("Server gracefully stopped")
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
		glog.V(3).Info("k===%s,value===%s \n", k, v)
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
		glog.V(3).Info("err is %s \n", err)
	}
	glog.V(3).Info("hostIp is %s, returnCode is %d \n", getClientIp(r), returnCode)
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
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func worker(ctx context.Context) {
	for {
		println("sleep")
		time.Sleep(1e0)
	}
}

func main() {
	go func() {
		//打开debug,方便远程获取pprof数据
		log.Println(http.ListenAndServe(":9090", nil))
	}()
	done := make(chan bool)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())
	for i:=0;i<30;i++{
		go worker(ctx)
	}
	
	go func() {
		sig := <-sigs
		fmt.Println(sig)
		cancel()
		done <- true
	}()
	<-done
	log.Println("[MAIN_INFO]=>stop service.")
}

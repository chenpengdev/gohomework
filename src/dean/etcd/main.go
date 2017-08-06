package main

import (
	_ "dean/etcd/routers"

	"fmt"
	"os"
	"os/signal"
	"syscall"

	"runtime/pprof"

	"github.com/astaxie/beego"
)

func main() {
	// cpu性能分析
	f, err := os.OpenFile("cpu.prof", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("server start error:", err)
		return
	}
	pprof.StartCPUProfile(f)

	// 信号处理
	go SignalProc(f)

	beego.Run()
}

func SignalProc(f *os.File) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGTERM)

	for {
		msg := <-ch
		switch msg {
		case syscall.SIGHUP:
			ProFile(f)
		}
	}
}

func ProFile(f *os.File) {
	fm, err := os.OpenFile("mem.prof", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("server start error:", err)
		return
	}

	pprof.StopCPUProfile()
	f.Close()

	pprof.WriteHeapProfile(fm)
	fm.Close()
}

package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-echarts/statsview"

	"apiports/internal/infrastructure/container"
	"apiports/internal/infrastructure/ui/adapter"
)

func main() {

	// http://localhost:18066/debug/statsview or http://localhost:18066/debug/pprof
	mgr := statsview.New()
	go mgr.Start()

	go func() {
		// create root context
		ctx := context.Background()

		// create container of dependencies
		ctn := container.New()

		// application entry point
		// inserting could be also triggered from HTTP, gRPC or other user interfaces (let's call them adapters)
		adapter.FileImport(ctx, ctn)
	}()

	// graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit)
	exitCh := make(chan int)

	go func() {
		for {
			s := <-quit
			if s == syscall.SIGTERM || s == syscall.SIGKILL || s == syscall.SIGINT {
				fmt.Println("got signal, initialising graceful shutdown...")
				time.Sleep(time.Second * 2)
				os.Exit(0)
			}
		}
	}()

	code := <-exitCh
	os.Exit(code)
}

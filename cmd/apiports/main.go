package main

import (
	"apiports/internal/infrastructure/container"
	"apiports/internal/infrastructure/httpserver"
	"apiports/internal/infrastructure/ui/adapter"
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-echarts/statsview"
)

func main() {

	// http://localhost:18066/debug/statsview or http://localhost:18066/debug/pprof
	mgr := statsview.New()
	go mgr.Start()

	// create root context
	ctx := context.Background()

	// create container of dependencies
	ctn := container.New()

	// wait till import from file ends up
	adapter.FileImport(ctx, ctn)

	// run HTTP serv in background
	go httpserver.New(ctn)

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

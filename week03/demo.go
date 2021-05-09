package main

import (
	"context"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())
	mux := http.NewServeMux()
	mux.HandleFunc("/check", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("success"))
	})

	// 退出
	serverOut := make(chan struct{})
	mux.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
		serverOut <- struct{}{}
	})

	server := http.Server{
		Handler: mux,
		Addr:    ":8080",
	}

	// g2
	g.Go(func() error {
		log.Println("----------> g2 running.")
		select {
		case <-ctx.Done():
			log.Println("----------> errgroup exit. g2 quit.")
		case <-serverOut:
			log.Println("----------> server will out.")
		}

		timeoutCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		log.Println("----------> shutting down server.")
		log.Println("----------> g2 quit.")
		return server.Shutdown(timeoutCtx)
	})

	// g3
	g.Go(func() error {
		log.Println("----------> g3 running.")
		quit := make(chan os.Signal, 0)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		select {
		case <-ctx.Done():
			log.Println("----------> g3 quit.")
			return ctx.Err()
		case sig := <-quit:
			return errors.Errorf(" g3 get os signal: %v. quit", sig)
		}
	})

	// g1
	g.Go(func() error {
		return server.ListenAndServe()
	})

	log.Println("exit", g.Wait())
}

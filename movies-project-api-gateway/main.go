package main

import (
	"context"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	l := log.New(os.Stdout, "movies-gateway ", log.LstdFlags)

	sm := mux.NewRouter()

	proxyHandler := NewProxyHandler()

	sm.PathPrefix("/").HandlerFunc(proxyHandler.ServeHTTP)

	s := &http.Server{
		Addr:         "localhost:8080",
		Handler:      sm,
		ErrorLog:     l,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		l.Printf("API Gateway listening on %s\n", s.Addr)

		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			l.Fatalf("listen: %s\n", err)
		}
	}()

	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)

	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	s.Shutdown(tc)
}

type ProxyHandler struct {
	reverseProxy *httputil.ReverseProxy
}

func NewProxyHandler() *ProxyHandler {
	targetURL, _ := url.Parse("http://localhost:9090")
	return &ProxyHandler{
		reverseProxy: httputil.NewSingleHostReverseProxy(targetURL),
	}
}

func (ph *ProxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ph.reverseProxy.ServeHTTP(w, r)
}

package cmd

import (
	"context"
	"flag"
	"fmt"
	"go-archive/internal/app/api"
	"go-archive/internal/pkg/logs"
	"go-archive/pkg/chi"
	"go-archive/pkg/chi/middleware"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	port        int
	logLevel    string
	showConsole bool
)

func init() {
	flag.IntVar(&port, "p", 5500, "web run in port (default is 5500)")
	flag.StringVar(&logLevel, "l", "debug", "log level (default is Debug)")
	flag.BoolVar(&showConsole, "c", true, "log show in console")
	flag.Parse()
}

func Exec() {
	logs.Init(logLevel, showConsole)
	logs.Info("service is start")
	addr := fmt.Sprintf("0.0.0.0:%d", port)
	logs.Info("server listen at:", addr)
	logs.Debug("log level is:", logLevel)
	middleware.InitLogRouter(logs.Router)
	r := chi.NewRouter()
	r.Use(middleware.RealIP)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	api.SetupRoutes(r)
	Startup(r, addr)
}

func Startup(r http.Handler, addr string) {
	server := &http.Server{Addr: addr, Handler: r}
	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			logs.Fatal(err)
		}
	}()
	SetupGracefulStop(server)
}

func SetupGracefulStop(server *http.Server) {
	serverCtx, serverStopCtx := context.WithCancel(context.Background())
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig
		logs.Warn("service is shuting down...")
		shutdownCtx, _ := context.WithTimeout(serverCtx, 30*time.Second)
		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				logs.Error("graceful shutdown timed out.. forcing exit.")
			}
		}()
		err := server.Shutdown(shutdownCtx)
		if err != nil {
			logs.Error(err)
		}
		logs.Info("service is shutdown")
		serverStopCtx()
	}()

	<-serverCtx.Done()
}

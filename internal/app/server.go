package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"

	"upperfile.com/internal/config"
	"upperfile.com/internal/db"
)

func Start() {
  r := mux.NewRouter()

  config.Load()
  db.Load()

	LoadRoutes(r)

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", config.Env.HOST, config.Env.PORT),

    WriteTimeout: time.Second * 15,
    ReadTimeout: time.Second * 15,
    IdleTimeout: time.Second * 60,
		Handler: r,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen: %s\n", err)
		}
	}()

	fmt.Println(` ___  ____    __    ____  ____  ____  ____  `)
	fmt.Println(`/ __)(_  _)  /__\  (  _ \(_  _)( ___)(  _ \ `)
	fmt.Println(`\__ \  )(   /(__)\  )   /  )(   )__)  )(_) )`)
	fmt.Println(`(___/ (__) (__)(__)(_)\_) (__) (____)(____/ `)
  fmt.Println("")
  fmt.Println("HOST: ", config.Env.HOST)
  fmt.Println("PORT: ", config.Env.PORT)

	// Signal channel to capture syscalls
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, os.Interrupt)

	<-done

	log.Print("Server Stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}

	log.Print("Server Exited Properly")

  os.Exit(0)
}

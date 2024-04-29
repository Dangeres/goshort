// Package main entry point to program
package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Dangeres/goshort/internal/constants"
	"github.com/Dangeres/goshort/internal/domain/handlers"
	"github.com/Dangeres/goshort/internal/middlewares"
	"github.com/Dangeres/goshort/internal/repository/redis"
)

func main() {
	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)

	defer stop()

	mux := http.NewServeMux()

	redis := redis.New()

	hndls := handlers.New(redis)

	mux.HandleFunc(
		fmt.Sprintf("GET /get/{%s}", constants.PathURL),

		middlewares.CalcTime(hndls.HUnShort),
	)

	mux.HandleFunc(
		"POST /create",

		middlewares.CalcTime(hndls.HShort),
	)

	fullIP := fmt.Sprintf("%s:%s", constants.ServerAddress, constants.ServerPort)

	server := http.Server{
		Addr:         fullIP,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	go func() {
		log.Printf("Starting server at %s\n", fullIP)

		if err := server.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Println(err)
		}
	}()

	<-ctx.Done()

	log.Println("Gracefull shutdown start")

	err := server.Shutdown(ctx)

	if err != nil {
		log.Println(err)
	}

	err = redis.Close()

	if err != nil {
		log.Println(err)
	}

	log.Println("Gracefull shutdown end")

	log.Println("EXIT main")
}

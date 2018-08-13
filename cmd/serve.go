package cmd

import (
	"os/signal"
	"syscall"
	"github.com/s4kibs4mi/zeroserver/api"
	"net/http"
	"fmt"
	"strconv"
	"context"
	"time"
	"os"
	"github.com/s4kibs4mi/zeroserver/log"
)

func ServeHttp() {
	port := 8282
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGKILL, syscall.SIGINT, syscall.SIGQUIT)

	r := api.Router()

	htsrvr := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", port),
		Handler: r,
	}

	go func() {
		log.Logger().Println("HTTP: Listening on port " + strconv.Itoa(port))
		log.Logger().Fatal(htsrvr.ListenAndServe())
	}()

	<-stop

	log.Logger().Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	htsrvr.Shutdown(ctx)

	log.Logger().Println("Server shutteddown gracefully")
}

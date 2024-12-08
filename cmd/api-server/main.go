package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorvk/todoapp/internal/initializers"
	_ "github.com/gorvk/todoapp/internal/routes"
)

func main() {
	initializers.LoadEnv()

	initializers.ConnectDB()

	initializers.ConnectAuth()

	server := configureServer()

	go listenServerFailure(server)

	startServer(server)
}

// starting the server
func startServer(server *http.Server) {
	apiHost := os.Getenv("API_HOST")
	fmt.Println("Server started at " + apiHost + server.Addr)
	err := server.ListenAndServe()

	if err != nil {
		fmt.Println("Server Shutdown!")
		panic(err)
	}
}

// shutdown server gracefully incase of failure
func listenServerFailure(server *http.Server) {
	osSignalChan := make(chan os.Signal, 1)
	signal.Notify(osSignalChan, os.Interrupt, syscall.SIGTERM)
	osSignal := <-osSignalChan

	fmt.Println("Gracefully shutting down server due to :", osSignal)
	context, cancelContext := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancelContext()
	server.Shutdown(context)
}

package utils

import (
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v3"
)

// StartServerWithGracefulShutdown function for starting server with a graceful shutdown.
func StartServerWithGracefulShutdown(a *fiber.App) {
	// Create channel for idle connections.
	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt) // Catch OS signals.
		<-sigint

		// Received an interrupt signal, shutdown.
		if err := a.Shutdown(); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("服务器无法关闭! 原因: %v", err)
		}

		close(idleConnsClosed)
	}()

	// Build Fiber connection URL.
	fiberConnURL, _ := ConnectionURLBuilder("fiber")

	// Run server.
	if err := a.Listen(fiberConnURL); err != nil {
		log.Printf("服务器无法运行! 原因: %v", err)
	}

	<-idleConnsClosed
}

// StartServer func for starting a simple server.
func StartServer(a *fiber.App) {
	// Build Fiber connection URL.
	fiberConnURL, _ := ConnectionURLBuilder("fiber")

	// Run server.
	if err := a.Listen(fiberConnURL); err != nil {
		log.Printf("服务器没有运行! 原因: %v", err)
	}
}

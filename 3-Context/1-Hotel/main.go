package main

import (
	"context"
	"fmt"
	"time"
)

// https://pkg.go.dev/context
func main() {
	// Iniciamos um contexto
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second) // Adicionamos um timeout de 3 segundos
	defer cancel()
	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Hotel booking cancelled. Timeout reached.")
		return
	case <-time.After(5 * time.Second):
		fmt.Println("Hotel booked successfully.")
	}
}

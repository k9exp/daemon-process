package main

import (
	"context"
	"io"
	"log"
	"time"
)

func run(ctx context.Context, config *Config, output io.Writer) error {
	log.SetOutput(output)

	for {
		ticker := time.NewTicker(config.interval)
		log.Printf("INTERVAL: %v\n", config.interval)
		select {
		case <-ticker.C:
			log.Printf("%s: %d", config.name, config.counter.Load())
			config.counter.Add(1)
		case <-ctx.Done():
			log.Println("Cleaning Run...")
			return nil
		}
	}
}

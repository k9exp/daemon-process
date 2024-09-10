package main

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"
)

type Config struct {
	counter  atomic.Int64
	interval time.Duration
	name     string
}

func (c *Config) Update() error {
	viper.SetConfigName("uv")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/etc/uv/")

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	d, err := time.ParseDuration(viper.GetString("interval"))
	if err != nil {
		return err
	}

	c.interval = d
	c.name = viper.GetString("name")
	c.counter = atomic.Int64{}

	return nil
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	fmt.Printf("PID: %d\n", os.Getpid())

	config := new(Config)

	err := config.Update()
	if err != nil {
		log.Fatal(err)
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	go func() {
		for {
			select {
			case s := <-signalChan:
				switch s {
				case syscall.SIGHUP:
					log.Println("Reloading Configuration...")
					if err := config.Update(); err != nil {
						log.Fatal(err)
					}
					log.Println("Reloaded!")
				case syscall.SIGTERM, syscall.SIGINT:
					log.Println("Stopping application gracefully...")
					cancel() // everyone that consumed this ctx
					// let them clean first, they must have handled <-ctx.Done,
					// that's why we didn't call os.Exit(code)
				}
			case <-ctx.Done():
				log.Println("Shutting application down...")
				// we didn't call cancel() here, because. the cancel() is already called
				// by someone else,
				// that's why <-ctx.Done() is executed

				// ctx.Done() is the product of cancel()
				os.Exit(0)
			}
		}
	}()

	if err := run(ctx, config, os.Stdout); err != nil {
		log.Fatal(err)
	}
}

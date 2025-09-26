package main

import (
	"fmt"

	"github.com/gooneraki/blog-aggregator-go/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("Error reading config: %v", err)
	}
	fmt.Println(cfg)

	config.SetUser("gooneraki")

	cfg, err = config.Read()
	if err != nil {
		fmt.Printf("Error reading config: %v", err)
	}
	fmt.Println(cfg)

}

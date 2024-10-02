package main

import (
	"fmt"

	сonfig "url-shortener/internal/config"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := сonfig.MustLoad()
	fmt.Println(cfg)

}

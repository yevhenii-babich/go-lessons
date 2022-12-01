package main

import (
	"flag"
	"log"

	"echo-with-workers/route"

	_ "github.com/joho/godotenv/autoload"
	"github.com/kelseyhightower/envconfig"
)

type myConfig struct {
	Nothing string `envconfig:"PARAM_NAME_NOTHING"`
}

var ip = flag.String("addr", "localhost", "server address")

func main() {
	var config myConfig
	flag.Parse()
	if err := envconfig.Process("", &config); err != nil {
		log.Fatalf("can't load config: %v", err)
	}
	log.Printf("cofig: %+v", config)
	router := route.Init()
	if err := router.Start(*ip + ":8081"); err != nil {
		router.Logger.Fatalf("can't start: %v", err)
	}
}

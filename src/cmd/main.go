package main

import (
	env "github.com/Netflix/go-env"
	"github.com/dblencowe/go-cocktails/src/pkg/cocktail"
	"github.com/dblencowe/go-cocktails/src/pkg/http"
	"log"
)

type Config struct {
	RecipePath  string `env:"RECIPE_PATH,required=true"`
	BindAddress string `env:"BIND_ADDR,default=:8080"`
}

func main() {
	var cfg Config
	if _, err := env.UnmarshalFromEnviron(&cfg); err != nil {
		log.Fatal(err)
	}

	if err := cocktail.LoadFromFile(cfg.RecipePath); err != nil {
		log.Fatal(err)
	}

	log.Fatalln(http.Serve(cfg.BindAddress))
}

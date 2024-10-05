package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/yesmishgan/go-pokeball/internal/app/dummy"
	"github.com/yesmishgan/go-pokeball/pkg/app"
)

// init is invoked before main()
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	dummyService := dummy.NewDummyService()

	a, _ := app.New()

	a.Run(
		dummyService,
	)
}

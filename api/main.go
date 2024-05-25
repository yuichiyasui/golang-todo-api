package main

import (
	"api/gen"
	"api/handler"
	"api/infrastructure"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	middleware "github.com/oapi-codegen/gin-middleware"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	port := flag.String("port", "8080", "Port for test HTTP server")
	flag.Parse()

	swagger, err := gen.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}

	db, err := infrastructure.NewDBClient()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Clear out the servers array in the swagger spec, that skips validating
	// that server names match. We don't know how this thing will be run.
	swagger.Servers = nil

	server, err := handler.NewServer(db)
	if err != nil {
		log.Fatal(err)
	}
	serverStrictHandler := gen.NewStrictHandler(server, nil)

	// This is how you set up a basic gin router
	r := gin.Default()

	origin := os.Getenv("APP_URL")
	if origin == "" {
		log.Fatal("Error loading .env file")
	}

	corsConfig := cors.New(cors.Config{
		AllowOrigins:     []string{origin},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
	r.Use(corsConfig)

	// Use our validation middleware to check all requests against the
	// OpenAPI schema.
	r.Use(middleware.OapiRequestValidator(swagger))

	gen.RegisterHandlers(r, serverStrictHandler)

	s := &http.Server{
		Handler: r,
		Addr:    net.JoinHostPort("0.0.0.0", *port),
	}

	log.Fatal(s.ListenAndServe())
}

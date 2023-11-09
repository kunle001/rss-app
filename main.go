package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main(){
	fmt.Println("hello world")
	godotenv.Load(".env")
	portString :=os.Getenv("PORT") 
	if portString==""{
		log.Fatal("PORT is not found in the environmet")
	}

	router:= chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods:[]string{"POST", "GET", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge:  300,
	}))

	// casting a functionality to a route
	v1Router := chi.NewRouter()

	v1Router.Get("/healthz", handleReadiness)
	v1Router.Get("/err", handleError)

	// mounting the router 
	router.Mount("/v1", v1Router)

	srv:= &http.Server{
		Handler: router,
		Addr: ":" + portString,
	}

	log.Printf("Sever starting on port %v", portString)
	err :=srv.ListenAndServe()

	if err != nil{
		log.Fatal(err)
	}
}
package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/kunle001/rss-app/internal/database"

	_ "github.com/lib/pq"
)

type apiConfig struct{
	DB *database.Queries
}

func main(){
	fmt.Println("hello world")
	godotenv.Load(".env")
	portString :=os.Getenv("PORT") 
	if portString==""{
		log.Fatal("PORT is not found in the environmet")
	}

	dbUrl:= os.Getenv("DB_URL");
	if dbUrl==""{
		log.Fatal("DB URL is not found")
	}

	conn, err := sql.Open("postgres", dbUrl)
	if err != nil{
		log.Fatal(("Can't connect to DB"))
	};

	

	apiCfg := apiConfig{
		DB: database.New(conn),
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
	v1Router.Post("/users", apiCfg.handlerCreateuser)
	v1Router.Get("/get-user",apiCfg.middlewareAuth(apiCfg.handlerGetUser))
	v1Router.Post("/create-feed",apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))
	v1Router.Get("/get-feeds",apiCfg.handlerGetFeed)
	v1Router.Post("/follow-feed",apiCfg.middlewareAuth(apiCfg.handlerCreateFeedFollow))
	v1Router.Get("/my-feed-followers",apiCfg.middlewareAuth(apiCfg.handlerGetMyFeedFollowers))
	v1Router.Delete("/unfollow-feed/{feedId}",apiCfg.middlewareAuth(apiCfg.handlerUnfollowFeed))

	// mounting the router 
	router.Mount("/v1", v1Router)

	srv:= &http.Server{
		Handler: router,
		Addr: ":" + portString,
	}

	log.Printf("Sever starting on port %v", portString)
	err = srv.ListenAndServe()

	if err != nil{
		log.Fatal(err)
	}
}
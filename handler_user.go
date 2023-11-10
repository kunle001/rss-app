package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/kunle001/rss-app/internal/database"
)

func ( apiCfg *apiConfig) handlerCreateuser(w http.ResponseWriter, r *http.Request){
	type parameters struct{
		Name string `json:"name"`
	}
	
	decoder:= json.NewDecoder(r.Body)

	params := parameters{}

	
	
	err := decoder.Decode(&params)

	if err!=nil{
		respondWithError(w, 500, fmt.Sprintf("Error parsing data %v", err))
	};

	user, err:= apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID: uuid.New(), 
		Name: params.Name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	});

	if err!=nil{
		respondWithError(w, 400, fmt.Sprintf("Error creting User %v", err))
	}

 
	respondWithJson(w, 201, user)
}

func ( apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User){
		respondWithJson(w, 200,databaseUserToUser(user))
}
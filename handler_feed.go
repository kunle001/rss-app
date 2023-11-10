package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/kunle001/rss-app/internal/database"
)

func ( apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User){
	type parameters struct{
		Name string `json:"name"`
		URL string `json:"url"`
	}
	
	decoder:= json.NewDecoder(r.Body)

	params := parameters{}

	err := decoder.Decode(&params)

	if err!=nil{
		respondWithError(w, 500, fmt.Sprintf("Error parsing data %v", err))
	};

	feed, err:= apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID: uuid.New(), 
		Name: params.Name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID: user.ID,
		Url: params.URL,
	});

	if err!=nil{
		respondWithError(w, 400, fmt.Sprintf("Error creting User %v", err))
	}

 
	respondWithJson(w, 201, databaseFeedToFeed(feed))
}

func ( apiCfg *apiConfig) handlerGetFeed(w http.ResponseWriter, r *http.Request){
		feeds, err:= apiCfg.DB.GetFeeds(r.Context());

		if err!=nil{
			respondWithError(w, 400, fmt.Sprintf("something went wrong getting feeds %v", err))
		}

		respondWithJson(w, 200,databaseFeedsToFeeds(feeds))
}
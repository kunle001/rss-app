package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/kunle001/rss-app/internal/database"
)

func ( apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User){
	type parameters struct{
		Feed_id uuid.UUID `json:"feed_id"`
	}
	
	decoder:= json.NewDecoder(r.Body)

	params := parameters{}

	err := decoder.Decode(&params)

	if err!=nil{
		respondWithError(w, 500, fmt.Sprintf("Error parsing data %v", err))
	};

		_, err= apiCfg.DB.GetFeedById(r.Context(), params.Feed_id);

		if err != nil{
			respondWithError(w, 404, "There is no feed with this id")
		}

	feedfollow, err:= apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID: uuid.New(), 
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID: user.ID,
		FeedID: params.Feed_id,
	});

	if err!=nil{
		respondWithError(w, 400, fmt.Sprintf("Error creting field %v", err))
	}

 
	respondWithJson(w, 201, feedfollow)
}

func ( apiCfg *apiConfig) handlerGetMyFeedFollowers(w http.ResponseWriter, r *http.Request, user database.User){

		followers, err:= apiCfg.DB.GetMyFeedFollowers(r.Context(), user.ID);

		if err!=nil{
			respondWithError(w, 400, fmt.Sprintf("something went wrong getting feeds %v", err))
		}

		respondWithJson(w, 200,followers)
}

func ( apiCfg *apiConfig) handlerUnfollowFeed(w http.ResponseWriter, r *http.Request, user database.User){
	fieldStr := chi.URLParam(r,"feedId")

	fieldId, err:= uuid.Parse(fieldStr)

	if err!=nil{
		respondWithError(w, 400, fmt.Sprintf("something went wrong getting feeds %v", err))
		return
	}


	_, err= apiCfg.DB.UnFollowFeed(r.Context(), database.UnFollowFeedParams{
		UserID: user.ID,
		FeedID: fieldId,
	});

	if err!=nil{
		respondWithError(w, 400, fmt.Sprintf("something went wrong getting feeds %v", err))
		return 
	}

	respondWithJson(w, 200,"unfollowed sucessfullly")
}


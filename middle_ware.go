package main

import (
	"fmt"
	"net/http"

	"github.com/kunle001/rss-app/internal/auth"
	"github.com/kunle001/rss-app/internal/database"
)


type authHandler func(http.ResponseWriter, *http.Request, database.User)


func (apiCfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request){
		apiKey, err := auth.GetApiKey(r.Header)

		if err!=nil{
			respondWithError(w, 403, fmt.Sprintf("user is not logged in: %v", err))
			return
		}

		user, err:=apiCfg.DB.GetUserByAPIKey(r.Context(),apiKey);

		if err!=nil{
			respondWithError(w, 401, fmt.Sprintf("invalid apiKey %v",err))
			return	
		};

		handler(w,r,user)
		}

}
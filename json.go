package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJson(w http.ResponseWriter, code int, payload interface{}){
	dat, err := json.Marshal(payload)

	if err !=nil{
		w.WriteHeader(500)
		log.Printf("Failed to marshall json response %v", payload )
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}

func respondWithError(w http.ResponseWriter, code int, msg string){
	if code> 499{
		log.Println("an internal server Error:", msg)
	};

	type errResponse struct {
		Error string `json:"error"`//--adding a json reflect tag, to describe how we want the message to be 
	}

	respondWithJson(w, code, errResponse{
		Error: msg,
	})
}
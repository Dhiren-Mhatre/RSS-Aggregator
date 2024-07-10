package main

import (
	"encoding/json"
	"net/http"
	"log"
)

func respondWithError(w http.ResponseWriter, code int,msg string){
	if code>499{
		log.Println("responding with 5xx err:",msg)
	}
	type errResponse struct {
		Error string 	 `json:"error"`
	}
	respondWithJSON(w ,code,errResponse{Error: msg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	dat, err := json.Marshal(payload)
	if err != nil {
			log.Printf("failed to marshal payload: %v", payload)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}
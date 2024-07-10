package main

import(
  "fmt"
  "net/http"
)


func handlerReadiness(w http.ResponseWriter,r *http.Request){
  respondWithJson(w,200,struct{}{})
}
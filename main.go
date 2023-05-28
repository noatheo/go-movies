package main 

import (
	"fmt"
	"log"
	_"os"
	_"github.com/noatheo/movies/store"
	"github.com/noatheo/movies/api"
	"github.com/noatheo/movies/auth"
	_"context"
	"net/http"
	"github.com/gorilla/mux"
)


func main() {
	r := mux.NewRouter()
	var authMiddle auth.AuthMiddleware

	//r.HandleFunc("/movies", Api.GetMovies).Methods(http.MethodGet) 
	r.HandleFunc("/movies/{mid}", Api.GetMovie).Methods(http.MethodGet)
	r.HandleFunc("/movies", Api.GetMovies).Methods(http.MethodGet)
	r.Handle("/movies/{mid}", authMiddle.IsAuthorized(http.HandlerFunc(Api.DeleteMovie))).Methods(http.MethodDelete)
	r.Handle("/movies", authMiddle.IsAuthorized(http.HandlerFunc(Api.CreateMovie))).Methods(http.MethodPost)
	r.Handle("/movies/{mid}", authMiddle.IsAuthorized(http.HandlerFunc(Api.UpdateMovie))).Methods(http.MethodPut)
	r.HandleFunc("/signup", Api.SignUp).Methods(http.MethodPost)
	// r.HandleFunc("/users", Api.GetUsers).Methods(http.MethodGet)
	r.HandleFunc("/login", Api.Login).Methods(http.MethodPost)
	r.Handle("/movies/upsert", authMiddle.IsAuthorized(http.HandlerFunc(Api.UpsertMovies))).Methods(http.MethodPost)
	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000",r))
	
}
package Api

import (

	"encoding/json"
	_"math/rand"
	"net/http"
	_"strconv"
	_"log"
	"github.com/gorilla/mux"
	"github.com/noatheo/movies/store"
	"fmt"
	
)

type LoginToken struct {
	Token string   `json:"token"`

}

func GetMovies(w http.ResponseWriter, r *http.Request )  {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader( http.StatusOK)
	
	test, err := store.ReadTable()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError )
		return
	}
	R, err2 := json.Marshal(test)
	
	if err2 != nil  {
		http.Error(w, err.Error(), http.StatusInternalServerError )
		// http.Error(w, err2.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(R)
	
	
	

}



func GetMovie(w http.ResponseWriter, r *http.Request){
	w.WriteHeader( http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	Result, err := store.ReadMovie(params["mid"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError )
		return
	}
	R, err2 := json.Marshal(Result)
	
	if err2 != nil  {
		http.Error(w, err.Error(), http.StatusInternalServerError )
		// http.Error(w, err2.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(R)
	
}

func DeleteMovie(w http.ResponseWriter, r *http.Request){
	w.WriteHeader( http.StatusOK)
	w.Header().Set("Content.Type", "application/json")

	params := mux.Vars(r)
	
	Result, err := store.DeleteMovie(params["mid"])
    if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError )
		return
	}

	R, err2 := json.Marshal(Result)
	
	if err2 != nil  {
		http.Error(w, err.Error(), http.StatusInternalServerError )
		// http.Error(w, err2.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(R)
}


func CreateMovie(w http.ResponseWriter, r *http.Request){
	w.WriteHeader( http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	var movie store.MovieTable
	_ = json.NewDecoder(r.Body).Decode(&movie)
    
	Result, err := store.CreateMovie(&movie)
    if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError )
		return
	}
	R, err2 := json.Marshal(Result)
	
	if err2 != nil  {
		http.Error(w, err.Error(), http.StatusInternalServerError )
		// http.Error(w, err2.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(R)

}

func UpdateMovie(w http.ResponseWriter, r *http.Request){
	w.WriteHeader( http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	var movie store.MovieTable
	_ = json.NewDecoder(r.Body).Decode(&movie)
    params := mux.Vars(r)
	Result, err := store.UpdateMovie(params["mid"] ,&movie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError )
		return
	}
	R, err2 := json.Marshal(Result)
	
	if err2 != nil  {
		http.Error(w, err.Error(), http.StatusInternalServerError )
		// http.Error(w, err2.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(R)


}

func SignUp(w http.ResponseWriter, r *http.Request){
	w.WriteHeader( http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	var user store.UserTable
	_ = json.NewDecoder(r.Body).Decode(&user)
	Result, err := store.SignUpUser(&user)
    if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError )
		return
	}
	R, err2 := json.Marshal(Result)
	
	if err2 != nil  {
		http.Error(w, err.Error(), http.StatusInternalServerError )
		// http.Error(w, err2.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(R)
}

 func Login(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader( http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	var user store.UserTable
	_ = json.NewDecoder(r.Body).Decode(&user)
	Result, err := store.LoginUser(&user)
	// Token, err2 := json.Marshal(Result)
	// fmt.Println(Token)
	var result LoginToken
	result.Token = Result
	
    if err != nil  {
		http.Error(w, err.Error(), http.StatusInternalServerError )
		// http.Error(w, err2.Error(), http.StatusInternalServerError)
		return
	}
	R, err2 := json.Marshal(result)
	
	if err2 != nil  {
		http.Error(w, err.Error(), http.StatusInternalServerError )
		// http.Error(w, err2.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(R)


	
     
 }


func UpsertMovies(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader( http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	var movie []*store.MovieTable
	_ = json.NewDecoder(r.Body).Decode(&movie)
	//fmt.Println(movie)
    fmt.Println("no")
	Results, err := store.Upsert(movie)
	if err != nil {
		
		http.Error(w, err.Error(), http.StatusInternalServerError )
		// http.Error(w, err2.Error(), http.StatusInternalServerError)
		
		return

	}

	R, err2 := json.Marshal(Results)
	if err2 != nil  {
		http.Error(w, err.Error(), http.StatusInternalServerError )
		// http.Error(w, err2.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(R)


}


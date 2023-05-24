package Api

import (

	"encoding/json"
	_"math/rand"
	"net/http"
	_"strconv"
	_"log"
	"github.com/gorilla/mux"
	"github.com/noatheo/movies/store"
	
)



func GetMovies(w http.ResponseWriter, r *http.Request )  {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader( http.StatusOK)
	
	test, err := store.ReadTable()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError )
		return
	}
	json.NewEncoder(w).Encode(test)
	
	
	

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
	json.NewEncoder(w).Encode(Result)
	
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

	json.NewEncoder(w).Encode(Result)
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
	json.NewEncoder(w).Encode(Result)

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
	json.NewEncoder(w).Encode(Result)


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
	json.NewEncoder(w).Encode(Result)
}

 func Login(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader( http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	var user store.UserTable
	_ = json.NewDecoder(r.Body).Decode(&user)
	Result, err := store.LoginUser(&user)
	// Token, err2 := json.Marshal(Result)
	// fmt.Println(Token)
	
    if err != nil  {
		http.Error(w, err.Error(), http.StatusInternalServerError )
		// http.Error(w, err2.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(Result)
     
 }

// func GetUsers(w http.ResponseWriter, r *http.Request ) {
// 	w.Header().Set("Content-Type", "application/json")
// 	test, err := store.Read()
// 	fmt.Println(err)
// 	json.NewEncoder(w).Encode(test)
	
// }


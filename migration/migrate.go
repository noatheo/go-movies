package main

import (
    "fmt"
    "encoding/json"
    _"log"
    "net/http"
    "os"
	_"github.com/lib/pq"
	"database/sql"
    "sync"
    "io/ioutil"
    "github.com/joho/godotenv"
)

var (
	DB           *sql.DB
	ConnectOnce sync.Once
)





func Connect() *sql.DB {
    err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Error :: %s", err.Error())
	}

    var DSN string = "postgres://" + os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@" + os.Getenv("DB_HOST") + "/" + os.Getenv("DB_NAME") + "?sslmode=" + os.Getenv("SSL_MODE")
	if DB == nil {
		ConnectOnce.Do(func() {
		var err error
		DB, err = sql.Open("postgres", DSN)
		if err != nil {
			fmt.Printf("%s", err)
			return
		}

		connectionError := DB.Ping()

		if connectionError != nil {
			fmt.Printf("%s", connectionError)
			return
		}	
	})
		}
	
	
		return DB
	}




func CloseConnection() {
	DB.Close()
}



type MovieTable struct {
	MID            float64 
    Title          string 
    Overview       string
    Release_date   string 
    Rating         float64 
}
func main(){
    Connect()
    ExecSql()
    
    ReadFromApi()
    CloseConnection()


}
func ExecSql() {
    query, err := ioutil.ReadFile("./schema.sql")
    if err != nil {
        panic(err)
    }
    if _, err := DB.Exec(string(query)); err != nil {
        panic(err)
    }
    

}

func ReadFromApi() {
    url := "https://api.themoviedb.org/3/discover/movie?include_adult=false&include_video=false&language=en-US&page=1&sort_by=popularity.desc"


	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiI1YjI0MzRlZWQzYjcwNjg1YjkxYzMxOWE5OTEwODk2ZSIsInN1YiI6IjY0NmNkMmYxNzA2YjlmMDEzODM0ZDRiNiIsInNjb3BlcyI6WyJhcGlfcmVhZCJdLCJ2ZXJzaW9uIjoxfQ.pis27KgK4qYNdrO34YTHCRXOMI6PqXwJKfVFIrqDzT0")
    

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
    var result map[string][]interface{}
    //var result map[string]any
    json.Unmarshal(body, &result)
    //results := result["results"].(map[string]interface{})
    // fmt.Println(len(result))
    // keys := make([]string, len(result))
    // i := 0
    // for k := range result {
    //     keys[i] = k
    //     i++
    // }
    // fmt.Println(keys)
    var movies = result["results"]
    i := 0
    
    var moviedb MovieTable
    for  i  < len(movies) {
        movie := movies[i].(map[string]interface{})

        moviedb.MID          = movie["id"].(float64)
        moviedb.Title        = movie["title"].(string)
        moviedb.Release_date = movie["release_date"].(string)
         moviedb.Rating       = movie["popularity"].(float64)
         moviedb.Overview  = movie["overview"].(string)
        _, erro := InsertData(&moviedb)
        fmt.Println(erro)
        
        i++
    }
    
    test := result["results"][1].(map[string]interface{})
    fmt.Println(len(test))
    key := make([]string, len(test))
    j := 0
    for k := range test {
        key[j] = k
        j++
    }
    fmt.Println(key)
    fmt.Println(test["id"]) 
    
   
}

func InsertData(movie *MovieTable) ( sql.Result , error ) {
	query := "INSERT INTO movies(mid , title , overview , release_date ,rating) VALUES ($1 ,$2 ,$3 ,$4 ,$5)"
	//mid, _ := strconv.ParseInt(movie.MID , 10 ,64 )
    Result, err := DB.Exec(query, movie.MID, movie.Title, movie.Overview, movie.Release_date, movie.Rating )
    if err != nil {
        return nil, fmt.Errorf("%v",  err)
    }

    return Result, nil
}








    






	
    










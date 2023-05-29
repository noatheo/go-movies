package store

import (
	_ "context"
	"database/sql"
	"fmt"
	_ "log"
	"sync"
	"time"

	
	_ "github.com/lib/pq"

	//_"github.com/joho/godotenv"
	"os"
	"strconv"
    
	"github.com/noatheo/movies/auth"
)

var (
	DB           *sql.DB
	ConnectOnce sync.Once
)

type UserTable struct {
    UID string  `json:"uid"`
    Username string  `json:"username"`
    Password string   `json:"password"`
    Email string     `json:"email"`
    Token sql.NullString     `json:"token"`
    Created_on  string `json:"created_on"`
}

type MovieTable struct {
	MID string        `json:"mid"`
    Title string      `json:"title"`
	Overview string    `json:"overview"`
	Release_date string `json:"release_date"`
    Rating string   `json:"rating"`
}



var DSN string = "postgres://" + os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@" + os.Getenv("DB_HOST") + "/" + os.Getenv("DB_NAME") + "?sslmode=" + os.Getenv("SSL_MODE")


func Connect(DSNs string) *sql.DB {
	if DB == nil {
		ConnectOnce.Do(func() {
		var err error
		DB, err = sql.Open("postgres", DSNs)
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







// Query Functions For Movie Table


func ReadTable() ([]MovieTable , error){
	DB := Connect(DSN)
	var table []MovieTable 
    query := "SELECT * FROM movies"
    rows, err := DB.Query(query)
    if err != nil {
        return nil, fmt.Errorf("%v",  err)
    }


    defer rows.Close()
    // Loop through rows, using Scan to assign column data to struct fields.

    for rows.Next() {
        var column MovieTable
        if err := rows.Scan(&column.MID, &column.Title, &column.Overview, &column.Release_date , &column.Rating ); err != nil {
            return nil, fmt.Errorf("%v",  err)
        }
        table = append(table, column)
    }

    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("%v",  err)
    }
    return table, nil
}

func DeleteMovie(mid string) (sql.Result , error){
	DB := Connect(DSN) 
    query := "DELETE FROM movies WHERE mid=$1"
    Result, err := DB.Exec(query, mid)
    if err != nil {
        return nil, fmt.Errorf("%v",  err)
    }

    return Result, nil
}

func ReadMovie(mid string ) (MovieTable, error) {
	DB := Connect(DSN)
	var movie MovieTable 
    query := "SELECT * FROM movies WHERE mid=$1"
    row  := DB.QueryRow(query, mid)
    
    // Loop through rows, using Scan to assign column data to struct fields.
	if err := row.Scan(&movie.MID, &movie.Title, &movie.Overview, &movie.Release_date , &movie.Rating ); err != nil {
		return movie, fmt.Errorf("%v",  err)
	}

    return movie, nil

}

func CreateMovie(movie *MovieTable) ( sql.Result , error ) {
	DB := Connect(DSN)
	query := "INSERT INTO movies(mid , title , overview , release_date, rating) VALUES ($1 ,$2 ,$3 ,$4 ,$5)"
	mid, _ := strconv.ParseInt(movie.MID , 10 ,64 )
    Result, err := DB.Exec(query, &mid, movie.Title, movie.Overview, movie.Release_date ,movie.Rating)
    if err != nil {
        return nil, fmt.Errorf("%v",  err)
    }

    return Result, nil
}

func UpdateMovie(midold string ,movie *MovieTable) ( sql.Result , error ) {
	DB := Connect(DSN)
	query := "UPDATE movies SET  title=$1 , overview=$2 , release_date=$3 ,rating=$4 WHERE mid = $5"
	midnew, _ := strconv.ParseInt(midold , 10 ,64 )
    Result, err := DB.Exec(query,  movie.Title, movie.Overview, movie.Release_date ,movie.Rating , &midnew  )
    if err != nil {
        return nil, fmt.Errorf("%v",  err)
    }

    return Result, nil
}






// Query Functions For Users Table

func SignUpUser(user *UserTable ) (sql.Result , error) {
	DB := Connect(DSN)
	query := "INSERT INTO users(uid , username , user_password , email ,created_on) VALUES ($1 ,$2 ,crypt($3, gen_salt('md5')) ,$4 ,$5 )"
	uid, _ := strconv.ParseInt(user.UID , 10 ,64 )
	created_on := time.Now()
    Result, err := DB.Exec(query, &uid, user.Username , user.Password , user.Email , &created_on  )
    if err != nil {
        return nil, fmt.Errorf("%v",  err)
    }

    return Result, nil

}

func LoginUser(user *UserTable) (string, error){
    DB := Connect(DSN)
    queryPassword := "SELECT (user_password = crypt($1 , (SELECT user_password FROM users WHERE email = $2))) AS passowrd_match FROM users WHERE email= $3"
    row  := DB.QueryRow(queryPassword, user.Password , user.Email , user.Email)
   
    
    
	if err := row.Scan(&user.Token ); err != nil {
		return user.Token.String , fmt.Errorf("%v",  err)
	}
    if user.Token.String != "true" {
        return "Wrong Password, please try again" , fmt.Errorf("unauthorized access")
    }
    var Token string
    var err error
    
    
    Token, err  = auth.GenJwt(user.Email)
    if err != nil {
        return "", err
    }    
    

    return Token , nil


}



func Upsert(movies []*MovieTable) ([]sql.Result , error)  {
    DB := Connect(DSN)

    query := "INSERT INTO movies(mid , title , overview , release_date, rating) VALUES ($1 ,$2 ,$3 ,$4 ,$5) ON CONFLICT (mid) DO UPDATE SET title = EXCLUDED.title , overview = EXCLUDED.overview , release_date = EXCLUDED.release_date , rating = EXCLUDED.rating "
    var results []sql.Result  
    for i := 0 ; i < len(movies); {
       // mid,_ := strconv.ParseInt(movie[i].MID , 10 ,64 )
        Result, err := DB.Exec(query, movies[i].MID, movies[i].Title, movies[i].Overview, movies[i].Release_date ,movies[i].Rating)
        if err != nil {
            return nil, fmt.Errorf("%v",  err)
        }
        results = append(results , Result)
        i++
    }

    return results , nil

}

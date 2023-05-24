package auth

import (
	"github.com/golang-jwt/jwt"
	"fmt"
	"time"
	"os"
	"strings"
	"net/http"
)

var SigningKey = []byte(os.Getenv("KEY"))
type AuthMiddleware struct{}

type Generate struct {
	Token string `json:""`
}


func GenJwt(email string ) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["email"] = email
	
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
    to := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	Token, err := to.SignedString(SigningKey)

	if err != nil {
		return "", err
	}

	return Token, nil	
}

func ValJwt(Token string ) (*jwt.Token, error) {

		token, err := jwt.Parse(Token , func(token *jwt.Token) (interface{}, error) {
			return []byte(SigningKey), nil
		})
		
		
		if err != nil {
			return nil, err
		}
	
		return token, nil
	}



func (auth *AuthMiddleware) IsAuthorized(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		isAuth := false
		au := r.Header.Get("Authorization")
		if au == " " {
			http.Error(w, "you must pass the token", http.StatusUnauthorized)
			return
		}

		tokenParts := strings.Split(au, " ")
		if len(tokenParts) < 2 {
			http.Error(w, "you must pass the token", http.StatusUnauthorized)
			return
		}
		token, err := ValJwt(tokenParts[1])
		if err != nil {
			fmt.Println(err)
			http.Error(w, "you must pass the token", http.StatusUnauthorized)
			return
		}
		fmt.Println(token)
		if email, ok := token.Claims.(jwt.MapClaims)["email"].(string); ok {
			if strings.HasSuffix(email, "@hotmail.com") || strings.HasSuffix(email, "@gmail.com") ||   strings.HasSuffix(email, "outlook.com")  {
				isAuth = true
			}
		}
		if !isAuth {
			fmt.Println(" user not authorized")
			http.Error(w, "you must pass the token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}


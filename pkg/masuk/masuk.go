package masuk

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"

	"eight/pkg/masuk/auth"
)

type Authenticator interface {
	Authenticate(r *http.Request) (auth.Info, error)
	Me(r *http.Request) (auth.Info, error)
}

type authenticator struct {
	secret string
}

func (a authenticator) Me(r *http.Request) (auth.Info, error) {
	return auth.NewDefaultUser(r.Context().Value("username").(string), 1), nil
}

func (a authenticator) Authenticate(r *http.Request) (auth.Info, error) {
	tokenString := getTokenString(r)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(a.secret), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		email := claims["email"]
		exp := claims["exp"]
		emailString := email.(string)
		expString := exp.(float64)
		fmt.Println(email)
		user := auth.NewDefaultUser(emailString, int64(expString))
		return user, nil
	} else {
		fmt.Println(err)
		return nil, err
	}
}

func getTokenString(r *http.Request) string {
	bearer := r.Header.Get("Authorization")
	return bearer[7:]
}

func New(secret string) Authenticator {
	return &authenticator{
		secret: secret,
	}
}

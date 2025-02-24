package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sinappsebackend/app"
	"sinappsebackend/services/auth"
	"sinappsebackend/services/users"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

func handleAuth(ctx *gin.Context) {
	url := app.OAuth.AuthCodeURL("state", oauth2.AccessTypeOffline)
	ctx.Redirect(http.StatusFound, url)
}

func handleAuthCallback(ctx *gin.Context) {
	code := ctx.Query("code")
	token, err := app.OAuth.Exchange(context.Background(), code)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	client := app.OAuth.Client(context.Background(), token)
	res, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}
	defer res.Body.Close()

	var userData map[string]any
	err = json.NewDecoder(res.Body).Decode(&userData)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	email, _ := userData["email"].(string)
	var user users.User

	if u, err := users.GetByEmail(email); err == nil {
		fmt.Printf("USUARIO EXISTENTE: %v\n", u)
		user = *u
	} else {
		username, _ := userData["name"].(string)
		user = users.User{Email: email, Username: username}
		id, err := users.CreateUser(user)
		if err != nil {
			fmt.Printf("error %s\n", err.Error())
			ctx.Status(http.StatusBadRequest)
			return
		}
		user.Id = id

		fmt.Printf("USUARIO CRIADO!!!: %d\n", id)
		fmt.Println(userData)
	}

	jwtToken := auth.GenerateToken(user)
	tokenStr, err := jwtToken.SignedString([]byte(app.JWTSecret))
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	fmt.Println("Token: " + tokenStr)
	ctx.Redirect(http.StatusFound, "/pages/login?token="+tokenStr)
}

/*func handleAuthCallback(ctx *gin.Context) {
	provider := ctx.Param("provider")
	conf := app.AuthConfigs[provider]
	if conf == nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	code := ctx.Query("code")
	token, err := conf.Exchange(context.Background(), code)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	client := conf.Client(context.Background(), token)
	res, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}
	defer res.Body.Close()

	var userData map[string]any
	err = json.NewDecoder(res.Body).Decode(&userData)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
	}

	fmt.Println(userData)
	ctx.Redirect(http.StatusFound, "/")
}*/
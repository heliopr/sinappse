package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.Static("/static", "./website/static")
	server.Static("/pages", "./website/pages")
	server.Static("/components", "./website/components")

	//server.StaticFile("/favicon.ico", "./website/static/SinappseLogo.png")
	server.StaticFile("/", "./website/pages/index.html")
	server.StaticFile("/login", "./website/pages/login/index.html")
	server.StaticFile("/decks", "./website/pages/decks/index.html")
	server.GET("/editar/:id", handleEditDeck)
	server.GET("/deck/:id", handleTrainDeck)

	server.GET("/auth/login/", handleAuth)
	server.GET("/auth/login/callback", handleAuthCallback)

	api := server.Group("/api")

	api.GET("/users/info", handleGetUserInfo)
	api.GET("/users/:id/decks", handleGetDecks)

	api.GET("/decks/:id", handleGetDeck)
	api.POST("/decks", handleCreateDeck)
	api.POST("/decks/:id/cards", handlePostCards)

	api.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "API version 1.0")
	})
}

func handleEditDeck(ctx *gin.Context) {
	ctx.File("./website/pages/editar/index.html")
}

func handleTrainDeck(ctx *gin.Context) {
	ctx.File("./website/pages/deck/index.html")
}
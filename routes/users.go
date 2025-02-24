package routes

import (
	"net/http"
	"sinappsebackend/services/auth"
	"sinappsebackend/services/decks"
	"sinappsebackend/services/users"
	"strconv"

	"github.com/gin-gonic/gin"
)

func handleGetUserInfo(ctx *gin.Context) {
	tokenStr := ctx.Query("token")
	if !auth.IsValidToken(tokenStr) {
		ctx.JSON(http.StatusOK, gin.H{
			"success": false,
		})
		return
	}

	id, err := auth.GetIdFromToken(tokenStr)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"success": false,
		})
		return
	}

	user, err := users.GetById(id)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"success": false,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"id": user.Id,
		"username": user.Username,
		"email": user.Email,
	})
}

func handleGetDecks(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "não foi possível converter id para inteiro",
		})
		return
	}

	res, err := decks.GetFromUserId(uint32(id))
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "não foi possível encontrar decks",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"decks": res,
	})
}
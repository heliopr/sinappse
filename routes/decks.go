package routes

import (
	"fmt"
	"net/http"
	"sinappsebackend/services/auth"
	"sinappsebackend/services/decks"
	"strconv"

	"github.com/gin-gonic/gin"
)

func handleGetDeck(ctx *gin.Context) {
	valid, userId, err := auth.AuthenticateRequest(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "erro ao autenticar",
		})
		return
	}

	if !valid {
		ctx.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "token de autenticação inválido",
		})
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "não foi possível converter id para inteiro",
		})
		return
	}

	res, err := decks.GetDeck(uint32(id))
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "não foi possível buscar deck",
		})
		fmt.Println(err.Error())
		return
	}

	if res.UserId != userId {
		ctx.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "você não é dono desse deck",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"deck": res,
	})
}

func handleCreateDeck(ctx *gin.Context) {
	valid, id, err := auth.AuthenticateRequest(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "erro ao autenticar",
		})
		return
	}

	if !valid {
		ctx.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "token de autenticação inválido",
		})
		return
	}

	type payload struct {
		Name string
	}
	var req payload
	err = ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "erro ao ler payload",
		})
		return
	}

	if nameLen := len(req.Name); nameLen < 3 || nameLen > 50 {
		ctx.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "nome inválido",
		})
		return
	}

	if exists, err := decks.DeckExistsFromUserByName(req.Name, id); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "erro",
		})
		return
	} else if (exists) {
		ctx.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "deck com esse nome já existe",
		})
		return
	}

	err = decks.CreateDeck(req.Name, id)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "erro ao criar deck",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func handlePostCards(ctx *gin.Context) {
	valid, userId, err := auth.AuthenticateRequest(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "erro ao autenticar",
		})
		return
	}

	if !valid {
		ctx.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "token de autenticação inválido",
		})
		return
	}

	deckId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "não foi possível converter deck id para inteiro",
		})
		return
	}

	type payload struct {
		Cards []map[string]any
	}
	var req payload
	err = ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "erro ao ler payload",
		})
		return
	}

	if isOwned, err := decks.IsDeckOwnedByUser(uint32(deckId), userId); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "erro ao consultar bd",
		})
		return
	} else if !isOwned {
		ctx.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "você não é dono desse deck",
		})
		return
	}

	if err = decks.UpdateCards(uint32(deckId), &req.Cards); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "erro",
		})
		fmt.Println(err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}
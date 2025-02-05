package main

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
)

type Player struct {
	Motivation int `json:"motivation"`
	Burnout    int `json:"burnout"`
}

var player = Player{Motivation: 100, Burnout: 0}
var events = []string{"Пятничный баг", "Выгорание", "Крутой релиз"}

func main() {
	r := gin.Default()

	r.POST("/start-game", func(c *gin.Context) {
		player = Player{Motivation: 100, Burnout: 0}
		c.JSON(http.StatusOK, gin.H{
			"message": "Игра началась",
			"player":  player,
		})
	})

	r.GET("/next-event", func(c *gin.Context) {
		event := events[rand.Intn(len(events))]
		c.JSON(http.StatusOK, gin.H{"event": event})
	})

	r.POST("/action", func(c *gin.Context) {
		action := struct {
			Choice string `json:"choice"`
		}{}
		if err := c.ShouldBindJSON(&action); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if action.Choice == "work" {
			player.Motivation -= 10
		} else if action.Choice == "rest" {
			player.Motivation += 5
		}

		c.JSON(http.StatusOK, gin.H{"player": player})
	})

	r.Run(":8080")
}

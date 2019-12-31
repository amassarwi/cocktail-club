package main

import (
	"cocktail-club/common"
	"cocktail-club/http_handlers"
	"cocktail-club/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	common.StoreInit()

	r.GET("/ping", http_handlers.Ping)
	r.GET("/cocktail/ingredient", middleware.QueryChecker(), http_handlers.CocktailByIngredient)
	r.GET("/cocktail/name/:name", http_handlers.CocktailByName)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

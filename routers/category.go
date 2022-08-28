package routers

import (
	"ericarthurc.com/handlers"

	"github.com/gin-gonic/gin"
)

func CategoryRouter(app *gin.Engine) {
	seriesRoutes := app.Group("/category")
	{
		seriesRoutes.GET("/:category", handlers.GetCategoryByParam)
	}
}

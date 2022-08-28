package routers

import (
	"ericarthurc.com/handlers"

	"github.com/gin-gonic/gin"
)

func SeriesRouter(app *gin.Engine) {
	seriesRoutes := app.Group("/series")
	{
		seriesRoutes.GET("/", handlers.GetSeriesList)
		seriesRoutes.GET("/:series", handlers.GetSeriesByParam)
	}
}

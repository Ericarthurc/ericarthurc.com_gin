package routers

import (
	"ericarthurc.com/handlers"

	"github.com/gin-gonic/gin"
)

func BlogRouter(app *gin.Engine) {
	indexBlogRoutes := app.Group("/")
	{
		indexBlogRoutes.GET("/", handlers.GetBlogList)
	}

	blogRoutes := app.Group("/blog")
	{
		blogRoutes.GET("/:blog", handlers.GetBlogByParam)
	}
}

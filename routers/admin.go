package routers

import "github.com/gin-gonic/gin"

func testAdminHandler(c *gin.Context) {
	c.JSON(200, map[string]string{"route": "admin"})
}

func AdminRouter(app *gin.Engine) {
	adminRoutes := app.Group("/admin")
	{
		adminRoutes.GET("/", testAdminHandler)
	}

	adminApiRoutes := adminRoutes.Group("/api/v1")
	{
		adminApiRoutes.POST("/login", testAdminHandler)
		adminApiRoutes.POST("/logout", testAdminHandler)
	}
	{
		adminApiRoutes.GET("/post/:id", testAdminHandler)
		adminApiRoutes.GET("/posts", testAdminHandler)
		adminApiRoutes.POST("/post", testAdminHandler)
		adminApiRoutes.PUT("/post/:id", testAdminHandler)
		adminApiRoutes.DELETE("/post/:id", testAdminHandler)

	}
}

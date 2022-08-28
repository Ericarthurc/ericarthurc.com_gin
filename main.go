package main

import (
	"fmt"
	"net/http"
	"os"

	"ericarthurc.com/routers"
	"ericarthurc.com/utilities"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environmental variables
	godotenv.Load()

	gin.SetMode(os.Getenv("GIN_MODE"))

	// Create gin engine
	app := gin.Default()

	// Set static file http paths
	app.Static("/public", "./public")
	app.StaticFile("/sw.js", "./public/js/sw.js")
	app.StaticFile("/robots.txt", "./public/robots.txt")

	// Set custom pongo2 html renderer
	app.HTMLRender = utilities.NewPongo(utilities.RenderOptions{
		TemplateDir:      "templates",
		TemplateExtenson: "j2",
		ContentType:      "text/html; charset=utf-8",
	})

	// Load in routers
	routers.BlogRouter(app)
	routers.SeriesRouter(app)
	routers.CategoryRouter(app)
	routers.AdminRouter(app)

	// Start http server
	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("HTTP_PORT")), app)
}

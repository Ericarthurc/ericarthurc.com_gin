package handlers

import (
	"ericarthurc.com/utilities/parsers"

	"github.com/flosch/pongo2/v6"
	"github.com/gin-gonic/gin"
)

/*
path / | method GET | privacy Public
*/
func GetBlogList(c *gin.Context) {

	posts := parsers.GetBlogPostsSlice()

	c.HTML(200, "index", pongo2.Context{"blogList": posts})
}

/*
path /blog/:blog | method GET | param blog | privacy Public
*/
func GetBlogByParam(c *gin.Context) {
	blog := c.Param("blog")

	markdown, meta := parsers.GetBlogPostByName(blog)

	c.HTML(200, "blogPost", pongo2.Context{"markdown": markdown, "meta": meta})
}

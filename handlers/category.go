package handlers

import (
	"ericarthurc.com/utilities/parsers"

	"github.com/flosch/pongo2/v6"
	"github.com/gin-gonic/gin"
)

/*
path /series/:series | method GET | privacy Public
*/
func GetCategoryByParam(c *gin.Context) {
	category := c.Param("category")

	categoryList := parsers.GetSliceByCategoryParam(category)

	c.HTML(200, "category", pongo2.Context{"category": category, "categoryList": categoryList})

}

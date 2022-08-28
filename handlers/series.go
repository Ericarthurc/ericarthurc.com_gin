package handlers

import (
	"fmt"

	"ericarthurc.com/utilities/parsers"

	"github.com/flosch/pongo2/v6"
	"github.com/gin-gonic/gin"
)

/*
path /series | method GET | privacy Public
*/
func GetSeriesList(c *gin.Context) {

	series := parsers.GetSeriesSlice()

	fmt.Println(series)

	c.HTML(200, "seriesIndex", pongo2.Context{"seriesList": series})

}

/*
path /series/:series | method GET | privacy Public
*/
func GetSeriesByParam(c *gin.Context) {
	series := c.Param("series")

	seriesList := parsers.GetSliceBySeriesParam(series)

	c.HTML(200, "series", pongo2.Context{"series": series, "seriesList": seriesList})

}

package api

import (
	"github.com/gin-gonic/gin"
)

func GetImage(c *gin.Context) {
	name := c.Param("name")
	imgUrl := `./image/` + name
	c.File(imgUrl)
}

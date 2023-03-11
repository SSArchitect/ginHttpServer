package get

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetIndexPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", "main page")
}

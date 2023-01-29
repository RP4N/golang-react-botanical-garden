package welcome

import (
	"fmt"
	"net/http"
	"reimanexample/service/internal/routing"

	"github.com/gin-gonic/gin"
)

const Path = "/welcome"
const Method = routing.GET

func HandleRequest(c *gin.Context) {
	name := c.Query("name")
	c.JSON(http.StatusOK, gin.H{"msg": fmt.Sprintf("Welcome %v!", name)})
}

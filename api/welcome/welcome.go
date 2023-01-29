package welcome

import (
	"net/http"
	"reimanexample/service/internal/routing"

	"github.com/gin-gonic/gin"
)

const Path = "/welcome"
const Method = routing.GET

func HandleRequest(c *gin.Context) {

	c.String(http.StatusOK, "Welcome!")
}

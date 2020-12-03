package customization

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Router _
func Router(w http.ResponseWriter, req *http.Request) (err error) {
	defer func() {
		if v := recover(); v != nil {
			err = fmt.Errorf("panicked with %v+ value", v)
		}
	}()
	engine.ServeHTTP(w, req)
	return
}

var engine *gin.Engine

func init() {
	engine = gin.New()
	rg := engine.RouterGroup
	rg.POST("token", token)
}

func token(c *gin.Context) {
	posted := struct {
		GrantType string `form:"grant_type"`
		Username  string `form:"username"`
		Password  string `form:"password"`
	}{}
	if err := c.Bind(&posted); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
	}
	fmt.Println(posted)
	c.JSON(http.StatusOK, "ok")
}

package routers
import (
       "net/http"
       "github.com/gin-gonic/gin"
	"example.com/v1/src/controllers"
)
//StartGin function
func InitRoute(router *gin.Engine)*gin.Engine {
       api := router.Group("/article")
       {
              api.GET("/", controllers.Index)
              api.GET("/hello", controllers.Hello)
       }
       router.NoRoute(func(c *gin.Context) {
              c.AbortWithStatus(http.StatusNotFound)
       })

       return router
}

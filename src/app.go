package src
import (
	"example.com/v1/src/routers"
	"github.com/gin-gonic/gin"
	"example.com/v1/src/drivers"
	"fmt"
	"example.com/v1/src/config"
	"example.com/v1/src/models"
)

func InitApp(){
	configDb := config.GetDbConfig();
	configWebServ := config.GetWebServConfig();
	conn , err := drivers.Connect(configDb);
	if err != nil{
		// stop process because database error
		panic(err)
	}

	// init created table from schema 
	models.Migrate(conn);

	webSev:= gin.Default();
	routers.InitRoute(webSev);
	webSev.Run(fmt.Sprintf(":%s",configWebServ["PORT"]));
}
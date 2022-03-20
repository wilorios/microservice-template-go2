//Package routes has the Gin web server Engine  instance and
//the routes of the ms solution
package routes

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/wilorios/microservice-template-go2/internal/configurations"
	"github.com/wilorios/microservice-template-go2/internal/controller"
	"github.com/wilorios/microservice-template-go2/internal/service"
)

var (
	personService    service.PersonServ         = *service.New()
	personController controller.PersonControlle = *controller.New(personService)
)

//SetupRouter start the gin Web Server with the routes.
func SetupRouter(conf configurations.Application) {
	log.Println("msg", "starting http server", "http:", conf.Port)

	server := gin.Default()

	server.GET("/person", func(ctx *gin.Context) {
		ctx.JSON(200, personController.FindAll())
	})

	server.POST("/person", func(ctx *gin.Context) {
		ctx.JSON(200, personController.Save(ctx))
	})

	server.Run(":" + conf.Port)
}

package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kendax/tic_tac_toe_go_internal/controllers"
)

//Handle routing for the program
func SetupRoutes() *gin.Engine {

	r := gin.Default()
	r.LoadHTMLGlob("templates/*/*.html")
	r.Static("/assets", "./assets")

	r.GET("/", controllers.Display)
	r.POST("/postinput", controllers.GameSave)
	r.GET("/userindex", controllers.ResultsValidation)
	r.GET("/display", controllers.Display)
	r.GET("/restart", controllers.Restart)

	return r
}

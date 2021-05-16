package main


import (
	"github.com/m0stly1/golang-restfulapi/router"
	"github.com/m0stly1/golang-restfulapi/handlers"
	)


var (
	httpRouter router.Router = router.NewMuxRouter()
)


func main() {
	httpRouter.GET("/messages/", handlers.GetMessages)
	httpRouter.GET("/message/{id:[0-9]+}", handlers.GetMessage)
	httpRouter.POST("/message/", handlers.AddMessage)
	httpRouter.PUT("/message/{id:[0-9]+}", handlers.UpdateMessage)
	httpRouter.DELETE("/message/{id:[0-9]+}", handlers.DeleteMessage)
	httpRouter.SERVE(":8000")
}



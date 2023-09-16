package server

import (
	"github.com/gin-gonic/gin"
)

func Start() {

	r := gin.Default()

	r.POST("/teetime", MakeReservationHandler)
	r.Run()
}

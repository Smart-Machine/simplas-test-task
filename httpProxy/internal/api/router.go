package api

import (
	"github.com/Smart-Machine/simplas-test-task/service/pkg/proto"
	"github.com/gin-gonic/gin"
)

var ServiceClient proto.ServiceClient

func SetupRouter(serviceClient proto.ServiceClient) *gin.Engine {
	ServiceClient = serviceClient

	router := gin.Default()
	ad := router.Group("/advertisement")
	ad.POST("/", Create)
	ad.PUT("/:id", Update)
	ad.DELETE("/:id", Delete)
	ad.GET("/", GetList)
	ad.GET("/:id", GetOne)

	return router
}

package routes

import "github.com/gin-gonic/gin"

func UserRoutes(Routes *gin.Engine) {
	Routes.POST("/users/signup")
	Routes.POST("/users/login")
	Routes.POST("/admin/addproduct")
	Routes.GET("/users/productview")
	Routes.GET("/users/search")
}

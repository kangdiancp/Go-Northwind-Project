package server

import (
	"codeid.northwind/controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter(routers *gin.Engine, controllerMgr *controllers.ControllerManager) *gin.Engine {

	categoryRoute := routers.Group("/category")
	{
		//router endpoint
		categoryRoute.GET("/", controllerMgr.CategoryController.GetListCategory)
		categoryRoute.GET("/:id", controllerMgr.CategoryController.GetCategory)
		categoryRoute.POST("/", controllerMgr.CategoryController.CreateCategory)

		categoryRoute.POST("/withProduct", controllerMgr.CategoryController.CreateCategoryWithProduct)

		categoryRoute.PUT("/:id", controllerMgr.CategoryController.UpdateCategory)
		categoryRoute.DELETE("/:id", controllerMgr.CategoryController.DeleteCategory)
	}

	userRoute := routers.Group("/user")
	{
		userRoute.POST("/signup", controllerMgr.UsersController.Signup)
		userRoute.POST("/signin", controllerMgr.UsersController.Signin)
		userRoute.POST("/logout", controllerMgr.UsersController.Logout)
	}

	return routers

}

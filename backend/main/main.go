package main

import (
	"Project01/main/controller"
	"Project01/main/dao"
	"Project01/main/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	userDao        dao.UserDao               = dao.InitUserDao()
	userService    service.UserService       = service.NewUser(userDao)
	userController controller.UserController = controller.NewUser(userService)

	itemDao        dao.ItemDao               = dao.InitItemDao()
	itemService    service.ItemService       = service.NewItem(itemDao)
	itemController controller.ItemController = controller.NewItem(itemService)
)

func main() {
	defer userDao.CloseDB()
	server := gin.New()
	server.Use(gin.Recovery(), gin.Logger())

	apiGroup := server.Group("/user")
	{
		apiGroup.GET("/", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"list":    userController.FindAllUser(),
				"message": "Success!"})
		})

		apiGroup.POST("/", func(context *gin.Context) {
			err := userController.InsertUser(context)
			if err != nil {
				context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				context.JSON(http.StatusOK, gin.H{"message": "Success!"})
			}
		})

		apiGroup.PUT("/:id", func(context *gin.Context) {
			err := userController.UpdateUser(context)
			if err != nil {
				context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				context.JSON(http.StatusOK, gin.H{"message": "Success!"})
			}
		})

		apiGroup.DELETE("/:id", func(context *gin.Context) {
			err := userController.DeleteUser(context)
			if err != nil {
				context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				context.JSON(http.StatusOK, gin.H{"message": "Success!"})
			}
		})
	}
	server.Run()
}

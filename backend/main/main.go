package main

import (
	"backend/main/controller"
	"backend/main/dao"
	"backend/main/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	userDao        = dao.InitUserDao()
	userService    = service.NewUserService(userDao)
	userController = controller.NewUserController(userService)

	itemDao = dao.InitItemDao()

	cartDao        = dao.InitCartDao()
	cartService    = service.NewCartService(itemDao, cartDao)
	cartController = controller.NewCartController(cartService)
)

func main() {
	defer userDao.CloseDB()
	server := gin.New()

	//store := cookie.NewStore([]byte("amazon"))
	//userSession := sessions.Sessions("userSession", store)
	//server.Use(gin.Recovery(), gin.Logger(), userSession)

	server.Use(gin.Recovery(), gin.Logger())

	userApiGroup := server.Group("/")
	{
		userApiGroup.POST("/register", func(context *gin.Context) {
			err := userController.Register(context)
			if err != nil {
				context.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
			} else {
				context.JSON(http.StatusOK, gin.H{
					"status": "success",
				})
			}
		})

		userApiGroup.POST("/login", func(context *gin.Context) {
			err := userController.Login(context)
			if err != nil {
				context.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
			} else {
				context.JSON(http.StatusOK, gin.H{
					"status": "success",
				})
			}
		})

		userApiGroup.POST("/logout", func(context *gin.Context) {
			err := userController.Logout(context)
			if err != nil {
				context.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
			} else {
				context.JSON(http.StatusOK, gin.H{
					"status": "success",
				})
			}
		})
	}

	cartApiGroup := server.Group("/")
	{
		cartApiGroup.POST("/addtoCart", func(context *gin.Context) {
			err := cartController.AddToCart(context)
			if err != nil {
				context.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
			} else {
				context.JSON(http.StatusOK, gin.H{
					"status": "success",
				})
			}
		})

		cartApiGroup.GET("/getCartItems", func(context *gin.Context) {
			cartVo, err := cartController.GetCartList(context)
			if err != nil {
				context.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
			} else {
				context.JSON(http.StatusOK, gin.H{
					"status": "success",
					"cart":   cartVo,
				})
			}
		})
	}

	err := server.Run()
	if err != nil {
		return
	}
}

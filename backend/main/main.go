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

	itemDao        = dao.InitItemDao()
	itemService    = service.NewItemService(itemDao)
	itemController = controller.NewItemController(itemService)

	cartDao        = dao.InitCartDao()
	cartService    = service.NewCartService(itemDao, cartDao)
	cartController = controller.NewCartController(cartService)

	orderDao        = dao.InitOrderDao()
	orderItemDao    = dao.InitOrderItemDao()
	orderService    = service.NewOrderService(orderDao, orderItemDao, cartDao, itemDao)
	orderController = controller.NewOrderController(orderService)
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

	itemApiGroup := server.Group("/")
	{
		itemApiGroup.GET("/getItems", func(context *gin.Context) {
			itemList, err := itemController.GetItemList(context)
			if err != nil {
				context.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
			} else {
				context.JSON(http.StatusOK, gin.H{
					"status": "success",
					"list":   itemList,
				})
			}
		})

		itemApiGroup.GET("/search", func(context *gin.Context) {
			itemList, err := itemController.SearchItem(context)
			if err != nil {
				context.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
			} else {
				context.JSON(http.StatusOK, gin.H{
					"status": "success",
					"list":   itemList,
				})
			}
		})

		itemApiGroup.GET("/getItemByID", func(context *gin.Context) {
			item, err := itemController.SearchItemById(context)
			if err != nil {
				context.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
			} else {
				context.JSON(http.StatusOK, gin.H{
					"status": "success",
					"item":   item,
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

	orderApiGroup := server.Group("/")
	{
		orderApiGroup.POST("/checkout", func(context *gin.Context) {
			err := orderController.Checkout(context)
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

		orderApiGroup.GET("/getOrderHistory", func(context *gin.Context) {
			orderHistoryVoList, err := orderController.GetHistory(context)
			if err != nil {
				context.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
			} else {
				context.JSON(http.StatusOK, gin.H{
					"status":    "success",
					"histories": orderHistoryVoList,
				})
			}
		})
	}

	err := server.Run()
	if err != nil {
		return
	}
}

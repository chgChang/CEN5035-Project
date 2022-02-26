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

func setUpServer() *gin.Engine {
	//defer userDao.CloseDB()
	server := gin.New()

	//store := cookie.NewStore([]byte("amazon"))
	//userSession := sessions.Sessions("userSession", store)
	//server.Use(gin.Recovery(), gin.Logger(), userSession)

	server.Use(gin.Recovery(), gin.Logger())

	userApiGroup := server.Group("/")
	{
		userApiGroup.POST("/api/register", func(context *gin.Context) {
			err := userController.Register(context)
			if err != nil {
				context.JSON(http.StatusOK, gin.H{
					"status": "error",
					"msg":    err.Error(),
				})
			} else {
				context.JSON(http.StatusOK, gin.H{
					"status": "success",
					"msg":    "register success",
				})
			}
		})

		userApiGroup.POST("/api/login", func(context *gin.Context) {
			err := userController.Login(context)
			if err != nil {
				context.JSON(http.StatusOK, gin.H{
					"status": "error",
					"msg":    err.Error(),
				})
			} else {
				context.JSON(http.StatusOK, gin.H{
					"status": "success",
					"msg":    "login success",
				})
			}
		})

		userApiGroup.POST("/api/logout", func(context *gin.Context) {
			err := userController.Logout(context)
			if err != nil {
				context.JSON(http.StatusOK, gin.H{
					"status": "success",
					"error":  err.Error(),
				})
			} else {
				context.JSON(http.StatusOK, gin.H{
					"status": "success",
					"msg":    "logout success",
				})
			}
		})

		userApiGroup.GET("/api/currentUser", func(c *gin.Context) {
			cookie, err := c.Request.Cookie("currentUserName")
			cookie2, err := c.Request.Cookie("currentUserEmail")
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"status": "error",
					"msg":    err.Error(),
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"status":    "succss",
					"userName":  cookie.Value,
					"userEmail": cookie2.Value,
				})
			}
		})

	}

	itemApiGroup := server.Group("/")
	{
		itemApiGroup.GET("/api/getItems", func(context *gin.Context) {
			itemList, err := itemController.GetItemList(context)
			if err != nil {
				context.JSON(http.StatusOK, gin.H{
					"status": "error",
					"msg":    err.Error(),
				})
			} else {
				context.JSON(http.StatusOK, gin.H{
					"status": "success",
					"msg":    "get item success",
					"list":   itemList,
				})
			}
		})

		itemApiGroup.GET("/api/search", func(context *gin.Context) {
			itemList, err := itemController.SearchItem(context)
			if err != nil {
				context.JSON(http.StatusOK, gin.H{
					"status": "error",
					"msg":    err.Error(),
				})
			} else {
				context.JSON(http.StatusOK, gin.H{
					"status": "success",
					"msg":    "search success",
					"list":   itemList,
				})
			}
		})

		itemApiGroup.GET("/api/getItemByID", func(context *gin.Context) {
			item, err := itemController.SearchItemById(context)
			if err != nil {
				context.JSON(http.StatusOK, gin.H{
					"status": "error",
					"msg":    err.Error(),
				})
			} else {
				context.JSON(http.StatusOK, gin.H{
					"status": "success",
					"msg":    "get item by id success",
					"item":   item,
				})
			}
		})
	}

	cartApiGroup := server.Group("/")
	{
		cartApiGroup.POST("/api/addtoCart", func(context *gin.Context) {
			err := cartController.AddToCart(context)
			if err != nil {
				context.JSON(http.StatusOK, gin.H{
					"status": "error",
					"msg":    err.Error(),
				})
			} else {
				context.JSON(http.StatusOK, gin.H{
					"status": "success",
					"msg":    "add to cart success",
				})
			}
		})

		cartApiGroup.GET("/api/getCartItems", func(context *gin.Context) {
			cartVo, err := cartController.GetCartList(context)
			if err != nil {
				context.JSON(http.StatusOK, gin.H{
					"status": "error",
					"error":  err.Error(),
				})
			} else {
				context.JSON(http.StatusOK, gin.H{
					"status": "success",
					"msg":    "get cart items success",
					"cart":   cartVo,
				})
			}
		})

		cartApiGroup.POST("/api/removeCart", func(context *gin.Context) {
			err := cartController.RemoveCart(context)
			if err != nil {
				context.JSON(http.StatusOK, gin.H{
					"status": "error",
					"msg":    err.Error(),
				})
			} else {
				context.JSON(http.StatusOK, gin.H{
					"status": "success",
					"msg":    "remove cart success",
				})
			}
		})

		cartApiGroup.POST("/api/deleteCartByItemId", func(context *gin.Context) {
			err := cartController.DeleteCartByItemId(context)
			if err != nil {
				context.JSON(http.StatusOK, gin.H{
					"status": "error",
					"msg":    err.Error(),
				})
			} else {
				context.JSON(http.StatusOK, gin.H{
					"status": "success",
					"msg":    "delete cart by item id success",
				})
			}
		})

		cartApiGroup.POST("/api/updateCart", func(context *gin.Context) {
			err := cartController.UpdateCart(context)
			if err != nil {
				context.JSON(http.StatusOK, gin.H{
					"status": "error",
					"msg":    err.Error(),
				})
			} else {
				context.JSON(http.StatusOK, gin.H{
					"status": "success",
					"msg":    "update cart success",
				})
			}
		})
	}

	orderApiGroup := server.Group("/")
	{
		orderApiGroup.POST("/api/checkout", func(context *gin.Context) {
			err := orderController.Checkout(context)
			if err != nil {
				context.JSON(http.StatusOK, gin.H{
					"status": "error",
					"msg":    err.Error(),
				})
			} else {
				context.JSON(http.StatusOK, gin.H{
					"status": "success",
					"msg":    "checkout success",
				})
			}
		})

		orderApiGroup.GET("/api/getOrderHistory", func(context *gin.Context) {
			orderHistoryVoList, err := orderController.GetHistory(context)
			if err != nil {
				context.JSON(http.StatusOK, gin.H{
					"status": "error",
					"msg":    err.Error(),
				})
			} else {
				context.JSON(http.StatusOK, gin.H{
					"status":    "success",
					"msg":       "get order history success",
					"histories": orderHistoryVoList,
				})
			}
		})
	}

	return server
}

func main() {
	defer userDao.CloseDB()
	server := setUpServer()

	err := server.Run()
	if err != nil {
		return
	}
}

package main

import (
	"backend/main/controller"
	"backend/main/dao"
	"backend/main/service"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	userDao        = dao.InitUserDao()
	userService    = service.NewUserService(userDao, cartDao, orderDao, orderItemDao)
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
	server.Use(static.Serve("/", static.LocalFile("./resources/build", true)))

	userApiGroup := server.Group("/api/")
	{
		userApiGroup.POST("/register", func(context *gin.Context) {
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

		userApiGroup.POST("/login", func(context *gin.Context) {
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

		userApiGroup.POST("/logout", func(context *gin.Context) {
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

		userApiGroup.GET("/currentUser", func(c *gin.Context) {
			user, err := userController.GetUserInfo(c)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"status": "error",
					"msg":    err.Error(),
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"status": "success",
					"data":   user,
				})
			}
		})

		userApiGroup.GET("/deleteUser", func(context *gin.Context) {
			err := userController.DeleteUser(context)
			if err != nil {
				context.JSON(http.StatusOK, gin.H{
					"status": "error",
					"msg":    err.Error(),
				})
			} else {
				context.JSON(http.StatusOK, gin.H{
					"status": "success",
					"msg":    "delete user success",
				})
			}
		})

	}

	itemApiGroup := server.Group("/api/")
	{
		itemApiGroup.POST("/addItems", func(context *gin.Context) {
			err := itemController.AddItem(context)
			if err != nil {
				context.JSON(http.StatusOK, gin.H{
					"status": "error",
					"msg":    err.Error(),
				})
			} else {
				context.JSON(http.StatusOK, gin.H{
					"status": "success",
					"msg":    "add item success",
				})
			}
		})

		itemApiGroup.POST("/updateItems", func(context *gin.Context) {
			err := itemController.UpdateItem(context)
			if err != nil {
				context.JSON(http.StatusOK, gin.H{
					"status": "error",
					"msg":    err.Error(),
				})
			} else {
				context.JSON(http.StatusOK, gin.H{
					"status": "success",
					"msg":    "update item success",
				})
			}
		})

		itemApiGroup.POST("/deleteItems", func(context *gin.Context) {
			err := itemController.DeleteItem(context)
			if err != nil {
				context.JSON(http.StatusOK, gin.H{
					"status": "error",
					"msg":    err.Error(),
				})
			} else {
				context.JSON(http.StatusOK, gin.H{
					"status": "success",
					"msg":    "update item success",
				})
			}
		})

		itemApiGroup.GET("/getItems", func(context *gin.Context) {
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

		itemApiGroup.GET("/search", func(context *gin.Context) {
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

		itemApiGroup.GET("/getItemByID", func(context *gin.Context) {
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

	cartApiGroup := server.Group("/api/")
	{
		cartApiGroup.POST("/addtoCart", func(context *gin.Context) {
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

		cartApiGroup.GET("/getCartItems", func(context *gin.Context) {
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

		cartApiGroup.POST("/removeCart", func(context *gin.Context) {
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

		cartApiGroup.POST("/deleteCartByItemId", func(context *gin.Context) {
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

		cartApiGroup.POST("/updateCart", func(context *gin.Context) {
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

		cartApiGroup.GET("/deleteCartByEmail", func(context *gin.Context) {
			err := cartController.DeleteCartByEmail(context)
			if err != nil {
				context.JSON(http.StatusOK, gin.H{
					"status": "error",
					"msg":    err.Error(),
				})
			} else {
				context.JSON(http.StatusOK, gin.H{
					"status": "success",
					"msg":    "delete cart by email success",
				})
			}
		})
	}

	orderApiGroup := server.Group("/api/")
	{
		orderApiGroup.POST("/checkout", func(context *gin.Context) {
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

		orderApiGroup.GET("/getOrderHistory", func(context *gin.Context) {
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

		orderApiGroup.GET("/deleteOrderByEmail", func(context *gin.Context) {
			err := orderController.DeleteOrder(context)
			if err != nil {
				context.JSON(http.StatusOK, gin.H{
					"status": "error",
					"msg":    err.Error(),
				})
			} else {
				context.JSON(http.StatusOK, gin.H{
					"status": "success",
					"msg":    "delete order by email success",
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

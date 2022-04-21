package main

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	_ "gorm.io/gorm/schema"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	server = setUpServer()
	//w      = httptest.NewRecorder()
)

func TestItem(t *testing.T) {
	// Get All The Items
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/getItems", nil)
	server.ServeHTTP(w, req)
	//itemList, _ := itemService.GetItemList()
	response := "{\"list\":[{\"id\":1,\"name\":\"Apple iPhone 13\",\"price\":829,\"pic_url\":\"https://m.media-amazon.com/images/I/71GLMJ7TQiL.jpg\",\"description\":\"6.1-inch Super Retina XDR display\"},{\"id\":2,\"name\":\"Apple MacBook Air - 13.3 inches - 8 GB RAM\",\"price\":999,\"pic_url\":\"https://m.media-amazon.com/images/I/71TPda7cwUL._AC_SL1500_.jpg\",\"description\":\"All-Day Battery Life – Go longer than ever with up to 18 hours of battery life.\"}],\"msg\":\"get item success\",\"status\":\"success\"}"
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, w.Body.String(), response)

	// Search Item By Id
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/api/getItemByID?id=1", nil)
	server.ServeHTTP(w, req)
	response = "{\"item\":{\"id\":1,\"name\":\"Apple iPhone 13\",\"price\":829,\"pic_url\":\"https://m.media-amazon.com/images/I/71GLMJ7TQiL.jpg\",\"description\":\"6.1-inch Super Retina XDR display\"},\"msg\":\"get item by id success\",\"status\":\"success\"}"
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, w.Body.String(), response)

	// Search Item By Name
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/api/search?searchKey=Apple", nil)
	server.ServeHTTP(w, req)
	response = "{\"list\":[{\"id\":1,\"name\":\"Apple iPhone 13\",\"price\":829,\"pic_url\":\"https://m.media-amazon.com/images/I/71GLMJ7TQiL.jpg\",\"description\":\"6.1-inch Super Retina XDR display\"},{\"id\":2,\"name\":\"Apple MacBook Air - 13.3 inches - 8 GB RAM\",\"price\":999,\"pic_url\":\"https://m.media-amazon.com/images/I/71TPda7cwUL._AC_SL1500_.jpg\",\"description\":\"All-Day Battery Life – Go longer than ever with up to 18 hours of battery life.\"}],\"msg\":\"search success\",\"status\":\"success\"}"
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, w.Body.String(), response)

}

func TestUserLogin(t *testing.T) {
	server := setUpServer()
	//server.Use(gin.Recovery(), gin.Logger(), userSession)
	//email, _, server := login()

	// Failed Register: Email Already Exists
	w := httptest.NewRecorder()
	param := make(map[string]string)
	param["username"] = "czhang"
	param["email"] = "czhang@qy"
	param["password"] = "123"
	jsonByte, _ := json.Marshal(param)
	req, _ := http.NewRequest("POST", "/api/register", bytes.NewReader(jsonByte))
	server.ServeHTTP(w, req)
	response := "{\"msg\":\"email already exists\",\"status\":\"error\"}"
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, w.Body.String(), response)

	// Register
	//w4 := httptest.NewRecorder()
	//param4 := make(map[string]string)
	//param4["username"] = "cztest"
	//param4["email"] = "czhang1@5"
	//param4["password"] = "123"
	//jsonByte4, _ := json.Marshal(param4)
	//req4, _ := http.NewRequest("POST", "/register", bytes.NewReader(jsonByte4))
	//server.ServeHTTP(w4, req4)
	//response4 := "{\"msg\":\"register success\",\"status\":\"success\"}"
	//assert.Equal(t, 200, w4.Code)
	//assert.Equal(t, w4.Body.String(), response4)

	// Failed Login: Not registered
	w2 := httptest.NewRecorder()
	param2 := make(map[string]string)
	param2["email"] = "czdcnm@qy"
	param2["password"] = "123412"
	jsonByte2, _ := json.Marshal(param2)
	req2, _ := http.NewRequest("POST", "/api/login", bytes.NewReader(jsonByte2))
	server.ServeHTTP(w2, req2)
	response2 := "{\"msg\":\"email or password is wrong\",\"status\":\"error\"}"
	assert.Equal(t, 200, w2.Code)
	assert.Equal(t, w2.Body.String(), response2)

	// Success Login
	w3 := httptest.NewRecorder()
	param3 := make(map[string]string)
	param3["email"] = "czhang@qy"
	param3["password"] = "123"
	jsonByte3, _ := json.Marshal(param3)
	req3, _ := http.NewRequest("POST", "/api/login", bytes.NewReader(jsonByte3))
	server.ServeHTTP(w3, req3)
	response3 := "{\"msg\":\"login success\",\"status\":\"success\"}"
	assert.Equal(t, 200, w3.Code)
	assert.Equal(t, w3.Body.String(), response3)
	//return param3["email"], param3["password"]

	// Failed Logout: Not log in.
	w4 := httptest.NewRecorder()
	param4 := make(map[string]string)
	param4["email"] = "czd@qy"
	jsonByte4, _ := json.Marshal(param4)
	req4, _ := http.NewRequest("POST", "/api/logout", bytes.NewReader(jsonByte4))
	server.ServeHTTP(w4, req4)
	response4 := "{\"error\":\"user not logged in\",\"status\":\"success\"}"
	assert.Equal(t, 200, w4.Code)
	assert.Equal(t, w4.Body.String(), response4)

	// Success Logout
	//w5 := httptest.NewRecorder()
	//param5 := make(map[string]string)
	//param5["email"] = "czhang@qy"
	//jsonByte5, _ := json.Marshal(param5)
	//req5, _ := http.NewRequest("POST", "/api/logout", bytes.NewReader(jsonByte5))
	//cookie := &http.Cookie{Name: "currentUser", Value: param5["email"]}
	//req5.AddCookie(cookie)
	//server.ServeHTTP(w5, req5)
	//response5 := "{\"msg\":\"logout success\",\"status\":\"success\"}"
	//assert.Equal(t, 200, w5.Code)
	//assert.Equal(t, w5.Body.String(), response5)

	w6 := httptest.NewRecorder()
	param6 := make(map[string]int)
	param6["itemId"] = 1
	param6["quantity"] = 1
	jsonByte6, _ := json.Marshal(param6)
	req6, _ := http.NewRequest("POST", "/api/addtoCart", bytes.NewReader(jsonByte6))
	cookie6 := &http.Cookie{Name: "currentUser", Value: param["email"]}
	req6.AddCookie(cookie6)
	server.ServeHTTP(w6, req6)
	response6 := "{\"msg\":\"add to cart success\",\"status\":\"success\"}"
	assert.Equal(t, 200, w6.Code)
	assert.Equal(t, w6.Body.String(), response6)

}

func TestAddToCart(t *testing.T) {
	server := setUpServer()

	// Login First.
	w3 := httptest.NewRecorder()
	param3 := make(map[string]string)
	param3["email"] = "czhang@qy"
	param3["password"] = "123"
	jsonByte3, _ := json.Marshal(param3)
	req3, _ := http.NewRequest("POST", "/api/login", bytes.NewReader(jsonByte3))
	server.ServeHTTP(w3, req3)
	response3 := "{\"msg\":\"login success\",\"status\":\"success\"}"
	assert.Equal(t, 200, w3.Code)
	assert.Equal(t, w3.Body.String(), response3)

	// Clear Cart
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/removeCart", nil)
	cookie := &http.Cookie{Name: "currentUser", Value: param3["email"]}
	req.AddCookie(cookie)
	server.ServeHTTP(w, req)

	// Check Cart List
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/api/getCartItems", nil)
	cookie = &http.Cookie{Name: "currentUser", Value: param3["email"]}
	req.AddCookie(cookie)
	server.ServeHTTP(w, req)
	response := "{\"cart\":{\"itemList\":null,\"totalPrice\":0},\"msg\":\"get cart items success\",\"status\":\"success\"}"
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, w.Body.String(), response)

	// Failed Add: Item not exists
	w = httptest.NewRecorder()
	param := make(map[string]int)
	param["itemId"] = 3
	param["quantity"] = 1
	jsonByte, _ := json.Marshal(param)
	req, _ = http.NewRequest("POST", "/api/addtoCart", bytes.NewReader(jsonByte))
	cookie = &http.Cookie{Name: "currentUser", Value: param3["email"]}
	req.AddCookie(cookie)
	server.ServeHTTP(w, req)
	response = "{\"msg\":\"item doesn't exist\",\"status\":\"error\"}"
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, w.Body.String(), response)

	// Failed Add: Quantity not positive integer.
	w = httptest.NewRecorder()
	param = make(map[string]int)
	param["itemId"] = 1
	param["quantity"] = -1
	jsonByte, _ = json.Marshal(param)
	req, _ = http.NewRequest("POST", "/api/addtoCart", bytes.NewReader(jsonByte))
	cookie = &http.Cookie{Name: "currentUser", Value: param3["email"]}
	req.AddCookie(cookie)
	server.ServeHTTP(w, req)
	response = "{\"msg\":\"please input the correct quantity\",\"status\":\"error\"}"
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, w.Body.String(), response)

	// Success Add
	w = httptest.NewRecorder()
	param = make(map[string]int)
	param["itemId"] = 1
	param["quantity"] = 1
	jsonByte, _ = json.Marshal(param)
	req, _ = http.NewRequest("POST", "/api/addtoCart", bytes.NewReader(jsonByte))
	cookie = &http.Cookie{Name: "currentUser", Value: param3["email"]}
	req.AddCookie(cookie)
	server.ServeHTTP(w, req)
	response = "{\"msg\":\"add to cart success\",\"status\":\"success\"}"
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, w.Body.String(), response)

	// Re-check Cart List
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/api/getCartItems", nil)
	cookie = &http.Cookie{Name: "currentUser", Value: param3["email"]}
	req.AddCookie(cookie)
	server.ServeHTTP(w, req)
	response = "{\"cart\":{\"itemList\":[{\"itemId\":1,\"itemName\":\"Apple iPhone 13\",\"price\":829,\"picUrl\":\"https://m.media-amazon.com/images/I/71GLMJ7TQiL.jpg\",\"description\":\"6.1-inch Super Retina XDR display\",\"quantity\":1}],\"totalPrice\":829},\"msg\":\"get cart items success\",\"status\":\"success\"}"
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, w.Body.String(), response)

	// Failed Update: Item Not In Cart
	w = httptest.NewRecorder()
	param = make(map[string]int)
	param["itemId"] = 2
	param["quantity"] = 1
	jsonByte, _ = json.Marshal(param)
	req, _ = http.NewRequest("POST", "/api/updateCart", bytes.NewReader(jsonByte))
	cookie = &http.Cookie{Name: "currentUser", Value: param3["email"]}
	req.AddCookie(cookie)
	server.ServeHTTP(w, req)
	response = "{\"msg\":\"this item is not in the cart\",\"status\":\"error\"}"
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, w.Body.String(), response)

	// Successful Updata: 1 -> 2
	w = httptest.NewRecorder()
	param = make(map[string]int)
	param["itemId"] = 1
	param["quantity"] = 2
	jsonByte, _ = json.Marshal(param)
	req, _ = http.NewRequest("POST", "/api/updateCart", bytes.NewReader(jsonByte))
	cookie = &http.Cookie{Name: "currentUser", Value: param3["email"]}
	req.AddCookie(cookie)
	server.ServeHTTP(w, req)
	response = "{\"msg\":\"update cart success\",\"status\":\"success\"}"
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, w.Body.String(), response)

	// Re-check Cart List
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/api/getCartItems", nil)
	cookie = &http.Cookie{Name: "currentUser", Value: param3["email"]}
	req.AddCookie(cookie)
	server.ServeHTTP(w, req)
	response = "{\"cart\":{\"itemList\":[{\"itemId\":1,\"itemName\":\"Apple iPhone 13\",\"price\":829,\"picUrl\":\"https://m.media-amazon.com/images/I/71GLMJ7TQiL.jpg\",\"description\":\"6.1-inch Super Retina XDR display\",\"quantity\":2}],\"totalPrice\":1658},\"msg\":\"get cart items success\",\"status\":\"success\"}"
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, w.Body.String(), response)

	// Failed Delete:
	w = httptest.NewRecorder()
	param = make(map[string]int)
	param["itemId"] = 2
	jsonByte, _ = json.Marshal(param)
	req, _ = http.NewRequest("POST", "/api/deleteCartByItemId", bytes.NewReader(jsonByte))
	cookie = &http.Cookie{Name: "currentUser", Value: param3["email"]}
	req.AddCookie(cookie)
	server.ServeHTTP(w, req)
	response = "{\"msg\":\"this item is not in the cart\",\"status\":\"error\"}"
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, w.Body.String(), response)

	// Success Delete:
	w = httptest.NewRecorder()
	param = make(map[string]int)
	param["itemId"] = 1
	jsonByte, _ = json.Marshal(param)
	req, _ = http.NewRequest("POST", "/api/deleteCartByItemId", bytes.NewReader(jsonByte))
	cookie = &http.Cookie{Name: "currentUser", Value: param3["email"]}
	req.AddCookie(cookie)
	server.ServeHTTP(w, req)
	response = "{\"msg\":\"delete cart by item id success\",\"status\":\"success\"}"
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, w.Body.String(), response)

	// Re-check Cart List
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/api/getCartItems", nil)
	cookie = &http.Cookie{Name: "currentUser", Value: param3["email"]}
	req.AddCookie(cookie)
	server.ServeHTTP(w, req)
	response = "{\"cart\":{\"itemList\":null,\"totalPrice\":0},\"msg\":\"get cart items success\",\"status\":\"success\"}"
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, w.Body.String(), response)
}

func TestCheckout(t *testing.T) {
	server := setUpServer()

	// Login First.
	w3 := httptest.NewRecorder()
	param3 := make(map[string]string)
	param3["email"] = "czhang@qy"
	param3["password"] = "123"
	jsonByte3, _ := json.Marshal(param3)
	req3, _ := http.NewRequest("POST", "/api/login", bytes.NewReader(jsonByte3))
	server.ServeHTTP(w3, req3)
	response3 := "{\"msg\":\"login success\",\"status\":\"success\"}"
	assert.Equal(t, 200, w3.Code)
	assert.Equal(t, w3.Body.String(), response3)

	// Check Order History
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/getOrderHistory", nil)
	cookie := &http.Cookie{Name: "currentUser", Value: param3["email"]}
	req.AddCookie(cookie)
	server.ServeHTTP(w, req)
	response := "{\"histories\":[{\"orderId\":\"d9d81ed1-e1e6-4a18-a579-cb78dbbc3438\",\"totalPrice\":829,\"orderDate\":\"2022-03-04T10:45:39.426-05:00\",\"itemList\":[{\"itemId\":1,\"itemName\":\"Apple iPhone 13\",\"price\":829,\"picUrl\":\"https://m.media-amazon.com/images/I/71GLMJ7TQiL.jpg\",\"description\":\"6.1-inch Super Retina XDR display\",\"quantity\":1}]}],\"msg\":\"get order history success\",\"status\":\"success\"}"
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, w.Body.String(), response)

	// Clear Cart
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/api/removeCart", nil)
	cookie = &http.Cookie{Name: "currentUser", Value: param3["email"]}
	req.AddCookie(cookie)
	server.ServeHTTP(w, req)

	// Success Add
	w = httptest.NewRecorder()
	param := make(map[string]int)
	param["itemId"] = 1
	param["quantity"] = 1
	jsonByte, _ := json.Marshal(param)
	req, _ = http.NewRequest("POST", "/api/addtoCart", bytes.NewReader(jsonByte))
	cookie = &http.Cookie{Name: "currentUser", Value: param3["email"]}
	req.AddCookie(cookie)
	server.ServeHTTP(w, req)
	response = "{\"msg\":\"add to cart success\",\"status\":\"success\"}"
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, w.Body.String(), response)

	// Check Cart List
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/api/getCartItems", nil)
	cookie = &http.Cookie{Name: "currentUser", Value: param3["email"]}
	req.AddCookie(cookie)
	server.ServeHTTP(w, req)
	response = "{\"cart\":{\"itemList\":[{\"itemId\":1,\"itemName\":\"Apple iPhone 13\",\"price\":829,\"picUrl\":\"https://m.media-amazon.com/images/I/71GLMJ7TQiL.jpg\",\"description\":\"6.1-inch Super Retina XDR display\",\"quantity\":1}],\"totalPrice\":829},\"msg\":\"get cart items success\",\"status\":\"success\"}"
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, w.Body.String(), response)

	// Checkout
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/api/checkout", nil)
	cookie = &http.Cookie{Name: "currentUser", Value: param3["email"]}
	req.AddCookie(cookie)
	server.ServeHTTP(w, req)
	//response = "{\"msg\":\"checkout success\",\"status\":\"success\"}"
	//assert.Equal(t, 200, w.Code)
	//assert.Equal(t, w.Body.String(), response)

	// Check Order History
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/api/getOrderHistory", nil)
	cookie = &http.Cookie{Name: "currentUser", Value: param3["email"]}
	req.AddCookie(cookie)
	server.ServeHTTP(w, req)
	//response = "{\"msg\":\"order history is empty\",\"status\":\"error\"}"
	//assert.Equal(t, 200, w.Code)
	//assert.Equal(t, w.Body.String(), response)
}

//func TestSearchItemById(t *testing.T) {
//	server := setUpServer()
//	w := httptest.NewRecorder()
//	var id int = 1
//	req, _ := http.NewRequest("GET", "/getItemByID?id="+strconv.Itoa(id), nil)
//	server.ServeHTTP(w, req)
//	response := ""
//
//	assert.Equal(t, 200, w.Code)
//	assert.Equal(t, w.Body.String(), response)
//}
//
//func TestCart(t *testing.T) {
//	server := setUpServer()
//	w := httptest.NewRecorder()
//	req, _ := http.NewRequest("GET", "/getCartItems", nil)
//	server.ServeHTTP(w, req)
//	//itemList, _ := itemService.GetItemList()
//	response := "{\"cart\":{\"itemList\":null,\"totalPrice\":0},\"msg\":\"get cart items success\",\"status\":\"success\"}"
//	assert.Equal(t, 200, w.Code)
//	assert.Equal(t, w.Body.String(), response)
//
//	// add to cart
//	param := make(map[string]int)
//	param["itemId"] = 1
//	param["quantity"] = 1
//	jsonByte, _ := json.Marshal(param)
//	req, _ = http.NewRequest("POST", "/addtoCart?1=1", bytes.NewReader(jsonByte))
//	server.ServeHTTP(w, req)
//	//itemList, _ := itemService.GetItemList()
//	response = "{\"msg\":\"add to cart success\",\"status\":\"success\"}"
//	assert.Equal(t, 200, w.Code)
//	assert.Equal(t, w.Body.String(), response)
//}

func StringEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	if (a == nil) != (b == nil) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func ParseToStr(mp map[string]string) string {
	values := ""
	for key, val := range mp {
		values += "&" + key + "=" + val
	}
	temp := values[1:]
	values = "?" + temp
	return values
}

func PostForm(uri string, param map[string]string, router *gin.Engine) *httptest.ResponseRecorder {
	req := httptest.NewRequest("POST", uri+ParseToStr(param), nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)
	return w
}

func PostJson(uri string, param map[string]interface{}, router *gin.Engine) *httptest.ResponseRecorder {
	jsonByte, _ := json.Marshal(param)
	req := httptest.NewRequest("POST", uri, bytes.NewReader(jsonByte))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	_ "gorm.io/gorm/schema"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestGetItem(t *testing.T) {
	server := setUpServer()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/getItems", nil)
	server.ServeHTTP(w, req)
	//itemList, _ := itemService.GetItemList()
	response := "{\"list\":[{\"id\":1,\"name\":\"Apple iPhone 13\",\"price\":829,\"pic_url\":\"https://m.media-amazon.com/images/I/71GLMJ7TQiL.jpg\",\"description\":\"6.1-inch Super Retina XDR display\"},{\"id\":2,\"name\":\"Apple MacBook Air - 13.3 inches - 8 GB RAM\",\"price\":999,\"pic_url\":\"https://m.media-amazon.com/images/I/71TPda7cwUL._AC_SL1500_.jpg\",\"description\":\"All-Day Battery Life – Go longer than ever with up to 18 hours of battery life.\"}],\"msg\":\"get item success\",\"status\":\"success\"}"

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, w.Body.String(), response)
}

func TestUserLogin(t *testing.T) {
	server := setUpServer()

	// Failed Register: Email Already Exists
	w := httptest.NewRecorder()
	param := make(map[string]string)
	param["username"] = "czhang"
	param["email"] = "czhang@qy"
	param["password"] = "123"
	jsonByte, _ := json.Marshal(param)
	req, _ := http.NewRequest("POST", "/register", bytes.NewReader(jsonByte))
	server.ServeHTTP(w, req)
	response := "{\"msg\":\"email already exists\",\"status\":\"error\"}"
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, w.Body.String(), response)
	fmt.Print(1)

	// Register
	w4 := httptest.NewRecorder()
	param4 := make(map[string]string)
	param4["username"] = "cz"
	param4["email"] = "czhang@5"
	param4["password"] = "123"
	jsonByte4, _ := json.Marshal(param4)
	req4, _ := http.NewRequest("POST", "/register", bytes.NewReader(jsonByte4))
	server.ServeHTTP(w4, req4)
	response4 := "{\"msg\":\"register success\",\"status\":\"success\"}"
	assert.Equal(t, 200, w4.Code)
	assert.Equal(t, w4.Body.String(), response4)

	// Failed Login: Not registered
	w2 := httptest.NewRecorder()
	param2 := make(map[string]string)
	param2["email"] = "czdcnm@qy"
	param2["password"] = "123412"
	jsonByte2, _ := json.Marshal(param2)
	req2, _ := http.NewRequest("POST", "/login", bytes.NewReader(jsonByte2))
	server.ServeHTTP(w2, req2)
	response2 := "{\"msg\":\"email or password is wrong\",\"status\":\"error\"}"
	assert.Equal(t, 200, w2.Code)
	assert.Equal(t, w2.Body.String(), response2)
	fmt.Print(2)

	// Failed Login: Email or Password is wrong.
	w5 := httptest.NewRecorder()
	param5 := make(map[string]string)
	param5["email"] = "czhang@qy"
	param5["password"] = "1234x"
	jsonByte5, _ := json.Marshal(param5)
	req5, _ := http.NewRequest("POST", "/login", bytes.NewReader(jsonByte5))
	server.ServeHTTP(w5, req5)
	response5 := "{\"msg\":\"email or password is wrong\",\"status\":\"error\"}"
	assert.Equal(t, 200, w5.Code)
	assert.Equal(t, w5.Body.String(), response5)

	// Success Login
	w3 := httptest.NewRecorder()
	param3 := make(map[string]string)
	param3["email"] = "czhang@qy"
	param3["password"] = "123"
	jsonByte3, _ := json.Marshal(param3)
	req3, _ := http.NewRequest("POST", "/login", bytes.NewReader(jsonByte3))
	server.ServeHTTP(w3, req3)
	response3 := "{\"msg\":\"login success\",\"status\":\"success\"}"
	assert.Equal(t, 200, w3.Code)
	assert.Equal(t, w3.Body.String(), response3)

	// Failed Logout: Not log in.
	//w = httptest.NewRecorder()
	//param = make(map[string]string)
	//param["email"] = "czd@qy"
	//jsonByte, _ = json.Marshal(param)
	//req, _ = http.NewRequest("POST", "/logout", bytes.NewReader(jsonByte))
	//server.ServeHTTP(w, req)
	//response = "{\"error\":\"user not logged in\",\"status\":\"success\"}"
	//assert.Equal(t, 200, w.Code)
	//assert.Equal(t, w.Body.String(), response)
	//
	//// Success Logout
	//w = httptest.NewRecorder()
	//param = make(map[string]string)
	//param["email"] = "czhang@qy"
	//jsonByte, _ = json.Marshal(param)
	//req, _ = http.NewRequest("POST", "/logout", bytes.NewReader(jsonByte))
	//server.ServeHTTP(w, req)
	//response = "{\"msg\":\"logout success\",\"status\":\"success\"}"
	//assert.Equal(t, 200, w.Code)
	//assert.Equal(t, w.Body.String(), response)

}

func Printf(s string, i int) {

}

func TestSearchItemById(t *testing.T) {
	server := setUpServer()
	w := httptest.NewRecorder()
	var id int = 1
	req, _ := http.NewRequest("GET", "/getItemByID?id="+strconv.Itoa(id), nil)
	server.ServeHTTP(w, req)
	response := ""

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, w.Body.String(), response)
}

func TestCart(t *testing.T) {
	server := setUpServer()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/getCartItems", nil)
	server.ServeHTTP(w, req)
	//itemList, _ := itemService.GetItemList()
	response := "{\"cart\":{\"itemList\":null,\"totalPrice\":0},\"msg\":\"get cart items success\",\"status\":\"success\"}"
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, w.Body.String(), response)

	// add to cart
	param := make(map[string]int)
	param["itemId"] = 1
	param["quantity"] = 1
	jsonByte, _ := json.Marshal(param)
	req, _ = http.NewRequest("POST", "/addtoCart?1=1", bytes.NewReader(jsonByte))
	server.ServeHTTP(w, req)
	//itemList, _ := itemService.GetItemList()
	response = "{\"msg\":\"add to cart success\",\"status\":\"success\"}"
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, w.Body.String(), response)
}

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

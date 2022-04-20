package main

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	server = setUpServer()
)

func TestItem(t *testing.T) {
	// Get All The Items
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/getItems", nil)
	server.ServeHTTP(w, req)
	//itemList, _ := itemService.GetItemList()
	response := "{\"list\":[{\"id\":1,\"name\":\"Apple iPhone 13\",\"price\":829,\"pic_url\":\"https://m.media-amazon.com/images/I/71GLMJ7TQiL.jpg\",\"description\":\"6.1-inch Super Retina XDR display\"},{\"id\":2,\"name\":\"Apple MacBook Air - 13.3 inches - 8 GB RAM\",\"price\":999,\"pic_url\":\"https://m.media-amazon.com/images/I/71TPda7cwUL._AC_SL1500_.jpg\",\"description\":\"All-Day Battery Life – Go longer than ever with up to 18 hours of battery life.\"},{\"id\":3,\"name\":\"BISSELL, 3079 Featherweight Cordless XRT 14.4V Stick Vacuum\",\"price\":159.64,\"pic_url\":\"https://m.media-amazon.com/images/I/71BiBaMzGhL._AC_SL1500_.jpg\",\"description\":\"Cordless, Lightweight Vacuum. 14.4V lithium-ion battery provides up to 25 minutes of run time and charges in just 4.5 hours.\"},{\"id\":4,\"name\":\"Elden Ring - Standard - Xbox [Digital Code]\",\"price\":59.99,\"pic_url\":\"https://m.media-amazon.com/images/I/61xmJC5KdeL._SL1476_.jpg\",\"description\":\"THE NEW FANTASY ACTION RPG - Rise, Tarnished, and be guided by grace to brandish the power of the Elden Ring and become an Elden Lord in the Lands Between.\"},{\"id\":5,\"name\":\"XTERRA Fitness TR150 Folding Treadmill\",\"price\":499.99,\"pic_url\":\"https://m.media-amazon.com/images/I/81kygN6HwuL._AC_SL1500_.jpg\",\"description\":\"Large 16\\\" x 50\\\" walking/running surface accommodates users of many sizes and stride lengths\"},{\"id\":6,\"name\":\"TAG Heuer Monaco Men's Watch\",\"price\":75000,\"pic_url\":\"https://m.media-amazon.com/images/I/61tHxjEE6RL._AC_UL1024_.jpg\",\"description\":\"Minimalist watches design with a unique texture of dial. This classic simple gent's wrist watch will be greatly suitable for any occasion.\"},{\"id\":7,\"name\":\"Outroad 26 Inch Folding Mountain Bike\",\"price\":178.97,\"pic_url\":\"https://m.media-amazon.com/images/I/61wyLwfCUqL._AC_SL1200_.jpg\",\"description\":\"21 Speed Full Suspension High-Carbon Steel MTB Foldable Bicycle\"},{\"id\":8,\"name\":\"TCL 32-inch 1080p Roku Smart LED TV - 32S327\",\"price\":219.99,\"pic_url\":\"https://m.media-amazon.com/images/I/71wYJc19PiL._AC_SL1500_.jpg\",\"description\":\"Easy Voice Control: Works with Amazon Alexa or Google Assistant to help you find movie titles\"},{\"id\":9,\"name\":\"Amazon.com Gift Card in Various Gift Boxes\",\"price\":100,\"pic_url\":\"https://m.media-amazon.com/images/I/818Xr8h2C4L._SL1500_.jpg\",\"description\":\"Gift Card is nested inside a specialty gift box\"},{\"id\":10,\"name\":\"Vonanda Faux Leather 3 Seater Sofa\",\"price\":479.99,\"pic_url\":\"https://m.media-amazon.com/images/I/71IzG7Z2bXL._AC_SL1500_.jpg\",\"description\":\"【Exquisite Style】 Featuring a mid-century feel, this sofa simply exudes class and elegance, designed to keep its look casual and clean even in high-traffic family rooms. This couch completes your home's decor with exquisite style.\"},{\"id\":11,\"name\":\"Sony Digitial Camera 24.3MP SLR Camera with 3.0-Inch LCD \",\"price\":893.31,\"pic_url\":\"https://m.media-amazon.com/images/I/41TJ26qVUSL._AC_.jpg\",\"description\":\"24 MP APS-C CMOS sensor and Focus Sensitivity Range :EV 0 to EV 20 (at ISO 100 equivalent with F2.8 lens attached)\"},{\"id\":12,\"name\":\"adidas Boys' Classic Puffer Jacket\",\"price\":49.99,\"pic_url\":\"https://m.media-amazon.com/images/I/71tkTLGUjVL._AC_UL1500_.jpg\",\"description\":\"Regular fit strikes a comfortable balance between loose and snug\"},{\"id\":13,\"name\":\"adidas Unisex-Adult MLS Soccer Ball\",\"price\":44.99,\"pic_url\":\"https://m.media-amazon.com/images/I/71pt-LBM01L._AC_SL1500_.jpg\",\"description\":\"A replica training soccer ball based on the Nativo 21 match ball\"},{\"id\":14,\"name\":\"adidas Men's 2020-21 Spain Home Jersey\",\"price\":54.9,\"pic_url\":\"https://m.media-amazon.com/images/I/71tFFRswkDS._AC_UL1500_.jpg\",\"description\":\"THE SPANISH HOME JERSEY, MADE BY ADIDAS FOR THE FANS. Born from the culture of the country and its lively streets. This jersey unites generations of fans.\"},{\"id\":15,\"name\":\"Playstation 5 Disc Version PS5 Console\",\"price\":1099,\"pic_url\":\"https://m.media-amazon.com/images/I/31JaiPXYI8L._AC_.jpg\",\"description\":\"4K-TV Gaming, 16.GB GDDR6 RAM, 8K Output, WiFi 6. Ultra-High Speed 825GB SSD - U Deal\"},{\"id\":16,\"name\":\"Sharpie S-Gel\",\"price\":10.09,\"pic_url\":\"https://m.media-amazon.com/images/I/814OjhsZbqS._AC_SL1500_.jpg\",\"description\":\"Medium Point (0.7mm), Black Ink Gel Pen, 12 Count\"},{\"id\":17,\"name\":\"Comix Disposable Face-Masks with 3-Layer Adult Masks\",\"price\":21.99,\"pic_url\":\"https://m.media-amazon.com/images/I/61bBQyLJY3L._AC_UL1000_.jpg\",\"description\":\"Keep Safe: These Face masks are class 1 medical devices, and can be used for healthcare use, but should not be used in surgical or high risk situations, or where there is a high risk of contact with liquids or infectious materials.\"},{\"id\":18,\"name\":\"Germ-X Hand Sanitizer, Original, 32 Fl. Oz (Pack of 4)\",\"price\":39.96,\"pic_url\":\"https://m.media-amazon.com/images/I/81IPVP6mccL._AC_SL1500_.jpg\",\"description\":\"Effective at eliminating 99.99% of many common harmful germs and bacteria in as little as 15 seconds\"},{\"id\":19,\"name\":\"hp Printer Paper\",\"price\":35.77,\"pic_url\":\"https://m.media-amazon.com/images/I/71+cIzAI5XL._AC_SL1500_.jpg\",\"description\":\"Perfect everyday office paper: Superior quality, reliability, and dependability for high-volume printing at home, at school and in the office. Perfect for everyday black and white printing.\"},{\"id\":20,\"name\":\"Blue Buffalo Life Protection Formula Natural Adult Dry Dog Food\",\"price\":62.99,\"pic_url\":\"https://m.media-amazon.com/images/I/71gUaaYXmjL._AC_SL1500_.jpg\",\"description\":\"REAL MEAT FIRST: Blue Buffalo foods always feature real meat as the first ingredient; High-quality protein from real chicken helps your dog build and maintain healthy muscles; Plus they contain wholesome whole grains, garden veggies and fruit\"}],\"msg\":\"get item success\",\"status\":\"success\"}"
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, w.Body.String(), response)

	// Search Item By id: Success
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/api/getItemByID?id=1", nil)
	server.ServeHTTP(w, req)
	response = "{\"item\":{\"id\":1,\"name\":\"Apple iPhone 13\",\"price\":829,\"pic_url\":\"https://m.media-amazon.com/images/I/71GLMJ7TQiL.jpg\",\"description\":\"6.1-inch Super Retina XDR display\"},\"msg\":\"get item by id success\",\"status\":\"success\"}"
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, w.Body.String(), response)

	// Search Item By Name: Success
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/api/search?searchKey=Apple", nil)
	server.ServeHTTP(w, req)
	response = "{\"list\":[{\"id\":1,\"name\":\"Apple iPhone 13\",\"price\":829,\"pic_url\":\"https://m.media-amazon.com/images/I/71GLMJ7TQiL.jpg\",\"description\":\"6.1-inch Super Retina XDR display\"},{\"id\":2,\"name\":\"Apple MacBook Air - 13.3 inches - 8 GB RAM\",\"price\":999,\"pic_url\":\"https://m.media-amazon.com/images/I/71TPda7cwUL._AC_SL1500_.jpg\",\"description\":\"All-Day Battery Life – Go longer than ever with up to 18 hours of battery life.\"}],\"msg\":\"search success\",\"status\":\"success\"}"
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, w.Body.String(), response)

	// Search Item By id: Fail (Item of id=999 doesn't exist)
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/api/getItemByID?id=999", nil)
	server.ServeHTTP(w, req)
	response = "{\"msg\":\"id doesn't exist\",\"status\":\"error\"}"
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, w.Body.String(), response)

	// Search Item By Name: Fail (Item of name "BMW" doesn't exist)
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/api/search?searchKey=BMW", nil)
	server.ServeHTTP(w, req)
	response = "{\"msg\":\"the search result is empty\",\"status\":\"error\"}"
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
	req, _ := http.NewRequest("POST", "/api/register", bytes.NewReader(jsonByte))
	server.ServeHTTP(w, req)
	response := "{\"msg\":\"email already exists\",\"status\":\"error\"}"
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, w.Body.String(), response)

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

	// Failed login: Wrong info
	w2 = httptest.NewRecorder()
	param2 = make(map[string]string)
	param2["email"] = "cz6678@qy"
	param2["password"] = "123412"
	jsonByte2, _ = json.Marshal(param2)
	req2, _ = http.NewRequest("POST", "/api/login", bytes.NewReader(jsonByte2))
	server.ServeHTTP(w2, req2)
	response2 = "{\"msg\":\"email or password is wrong\",\"status\":\"error\"}"
	assert.Equal(t, 200, w2.Code)
	assert.Equal(t, w2.Body.String(), response2)

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

	// Success Logout
	w5 := httptest.NewRecorder()
	param5 := make(map[string]string)
	param5["email"] = "czhang@qy"
	jsonByte5, _ := json.Marshal(param5)
	req5, _ := http.NewRequest("POST", "/api/logout", bytes.NewReader(jsonByte5))
	cookie := &http.Cookie{Name: "currentUser", Value: param5["email"]}
	req5.AddCookie(cookie)
	server.ServeHTTP(w5, req5)
	response5 := "{\"msg\":\"logout success\",\"status\":\"success\"}"
	assert.Equal(t, 200, w5.Code)
	assert.Equal(t, w5.Body.String(), response5)
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
	param["itemId"] = 33
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

	// Clear Cart
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/removeCart", nil)
	cookie := &http.Cookie{Name: "currentUser", Value: param3["email"]}
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
	response := "{\"msg\":\"add to cart success\",\"status\":\"success\"}"
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

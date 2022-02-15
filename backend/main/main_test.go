package main

import (
	"github.com/go-playground/assert/v2"
	"net/http"
	"net/http/httptest"
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

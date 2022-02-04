# Sprint 1

Wei Wu (UFID: 8028-5179)
Chang Zhou (UFID: 70697466)
Jianan He (UFID: 68530029)
Chi Zhang (UFID: 91832967)
Github Link: https://github.com/chgChang/CEN5035-Project

---

## Application Description

Wikipedia says:

```
E-commerce is the activity of buying or selling online. Electronic commerce draws on technologies such as mobile commerce, electronic funds transfer, supply chain management, Internet marketing, online transaction processing, electronic data interchange (EDI), inventory management systems, and automated data collection systems. 
```

The application is an e-commercial web application like Amazon.

## Backend

- Demo Video Link:

Approach to start the backend:

1. ```shell
   cd backend/main
   ```

2. ```shell
   go run main.go
   ```

Below is the technology and environment of the backend.

```
Go v1.17.5
Gin v1.7.7
Gorm v1.22.5
MySQL 8.0
```

We choose `Gin` as our web application framework. Below is the description of `Gin`.

```
Gin is a web framework written in Go (Golang). It features a martini-like API with performance that is up to 40 times faster thanks to httprouter. If you need performance and good productivity, you will love Gin.
```

In sprint 1, we have developed 4 function modules, including user module, item module, cart module, order module.

The link of Api Document: https://documenter.getpostman.com/view/12317519/UVeGpQY8

### User Module

In user module, there are 3 apis which match functions of register, login and logout.

1. User Register Api

   The request body needs to include username, email and password. If someone has used the email to register an account, the server will respond an error including the message `"email already exists"`. Otherwise, the response will be "success".

2. User Login Api

   The request body needs to include email and password. If the email matches the password, the server will respond a success message. Otherwise, the response will be `"email or password is wrong"`. After logging in, the email will be saved into a cookie.

3. User Logout Api

   The request body needs to include email. The response will be "success" success only if the email matches the value of the cookie.

### Item Module

In user module, there are 3 apis which match functions including getting the full item list, searching specific products and search product by item id.

4. Get Item List Api

   The request body or parameters is not needed, because both logged-in users and visitors can browse the products. The server will respond a list of all the products.

5. Search Item Api

   The request needs to include the keyword of searching. The server will respond the list of products that has the keyword  in their names.

6. Use ID to Get Item Api

   The request needs to include the ID of the item. The server will respond the detail imformation of the result product. If the ID doesn't exist, the response will be an error.

### Cart Module

In cart module, there are 5 apis which match functions including adding products to cart, getting the list of products, updating the quantity of products, deleting specific item and remove all products from the cart. The server will get the email of the current user from the cookie.

7. Add Product into Cart Api

   Users need to send the request including item ID and quantity that they want to add. The legal quantity must be greater than 0 and the item ID must exist.

8. Get the Cart List Api

   The request can be null. The response will be the list of products in the cart.

9. Update Cart Api

   The request needs to include the item ID and the quantity to be updated. The item ID must match the item in the cart and the quantity must not be less than 0. If the quantity equals to 0, the item will be deleted from the cart.

10. Delete Item from Cart Api

    The request only needs to include the item ID to be deleted. The item ID must match the item in the cart.

11. Remove All Items from Cart Api

    The request parameter and body can be null. If the cart is not empty before, the items in cart will be all removed.

### Order Module

In order module, there are 2 apis which match the functions including checkout and getting the order history.

12. Checkout Api

    The request body needs to include the shipping information of the user and the cart must not be empty. The server will generate a unique order ID randomly for each order and save the information of the products in cart into database. After that, all items will be removed from the cart.

13. Get Order History Api

    The request parameter and body can be null. The response will be the list of the order history including the shipping information, the item information and the total price.


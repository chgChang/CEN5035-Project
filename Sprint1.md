# Sprint 1

Wei Wu (UFID: 8028-5179)

Chang Zhou (UFID: 70697466)

Jianan He (UFID: 68530029)

Chi Zhang (UFID: 91832967)

Github Link: https://github.com/chgChang/CEN5035-Project

---

## Application Description

From [Wikipedia](https://en.wikipedia.org/wiki/E-commerce):

E-commerce is the activity of buying or selling online. Electronic commerce draws on technologies such as mobile commerce, electronic funds transfer, supply chain management, Internet marketing, online transaction processing, electronic data interchange (EDI), inventory management systems, and automated data collection systems. 

This project is an Amazon-like E-commercial web application.

## Backend

In sprint 1, we have developed 4 function modules including user module, item module, cart module and order module.

- [The link of the Api Document](https://documenter.getpostman.com/view/12317519/UVeGpQY8)

- [The link of the Demo Video](https://youtu.be/l16GmmpDJmc)

Approach to start the backend:

1. ```shell
   cd backend/main
   ```

2. ```shell
   go run main.go
   ```

The technologies and environments of the backend.
```
Go v1.17.5
Gin v1.7.7
Gorm v1.22.5
MySQL 8.0
```

We choose [Gin](https://github.com/gin-gonic/gin) as our web application framework. Below is the description of it:

Gin is a web framework written in Go (Golang). It features a martini-like API with performance that is up to 40 times faster thanks to httprouter. 



### User Module

The user module has 3 APIs for functions of **register**, **login** and **logout**.

1. User Register API

   The request body includes username, email and password. The email must be **unquie**. 
   - If the register successes, the server returns a message `"success"`
   - If someone uses a duplicated email to register, the server responds an error including the message `"email already exists"`. 


2. User Login API

   The request body includes email and password. 
   - If the user's credential (email and password) matches, the server responds a message `"success"`. 
   - Otherwise, the server responses `"email or password is wrong"`. 
   After logging in, the email is saved into a cookie.

3. User Logout API

   The request body includes email. The response is "success" success only if the email matches the value of the cookie.
   - If the email matches that of the cookie, the servers responses `"success"`.
   - Otherwise, the server response `"user not logged in"`

### Item Module

The item module has 3 APIs for functions of **getting the full item list**, **fuzzy searching by name** and **exact searching id**.

4. Get Item List API

   The request body or parameters is not needed. 
   - The server responds by returning a list of all the products.

5. Search Item API

   The request includes the keyword to be searched. 
   - The server responds by returning a list of products having the keyword in their names.
   - If the given pattern doesn't match any item names, the server returns an error message.

6. Use ID to Get Item API

   The request includes the item ID to be searched. 
   - The server returns the detailed information of the result product. 
   - If the ID doesn't exist, the response will be an error.

### Cart Module

The cart module has 5 APIs which match functions of **adding products to cart**, **getting the list of products**, **updating the quantity of products**, **deleting a specific item** and **remove all the products from the cart**. The server will get the email of the current user from the cookie.

7. Add Product into Cart API

   The request includes item ID and the quantity to be added. The quantity must be a positive integer and the item ID must exist.
   - The server returns the message `"success"`
   - If the item ID doesn't exist, the server returns `"item doesn't exist"`
   - If the item quantity is not a positive integer, the server returns `"please input the correct quantity"`

8. Get the Cart List API

   The request body or parameters is not needed.
   - The server response by returning the list of products in the cart.

9. Update Cart API

   The request includes the item ID and the quantity to be updated. The item ID must match an item in the cart and the quantity must be an integer no less than 0. If the quantity equals to 0, the item will be deleted from the cart.
   - The server returns the message `"success"`
   - If the item ID doesn't exist in the cart, the server returns `"this item is not in the cart"`
   - If the item quantity is not a natural number, the server returns `"please input the correct quantity"`

10. Delete Item from Cart API

    The request includes the item ID to be deleted. The item ID must match an item in the cart.
    - The server returns the message `"success"`
    - If the item ID not in the cart, the server returns `"this item is not in the cart"`

11. Remove All Items from Cart API

    The request body or parameters is not needed. If the current cart is not empty, all the items in cart will be removed.
    - The server returns the message `"success"`
    - If the cart is empty, the server returns `"cart is empty, cannot remove"`

### Order Module

The order module has 2 APIs for functions of **checkout** and **getting the order history**.

12. Checkout API

    The request body includes the shipping information (shipping address, phone number and name) of the receiver and the cart must not be empty. The server randomly generates an unique order ID for each order and save the information of the products in the cart into the database. Next, all the items will be removed from the cart.
    - The server returns `"success"`
    - If the cart is empty, the server returns `"cart is empty"`

13. Get Order History API

    The request body or parameters is not needed. The server returns the list of the preveious order including the shipping information, the item information and the total price.
    - The server returns a list of previous orders.
    - If the current user has no previous order, the server returns `"order history is empty"`

## Frontend

1. Use react, umi, ant design components library, ant design pro for front-end page design and development of Amazon like shopping website
2. Use json file to simulate backend data transfer
3. Implemented pages: login interface, registration interface, shopping home page, shopping cart interface, checkout interface, account setting interface
4. Function:
   1. Users can log in or register a new account on the login interface
   2. After logging in to the account, users can select the page they want to enter through the side navigation bar. The current navigation bar has three pages: shopping homepage, shopping cart and account management
   3. On the shopping homepage, users can browse the product list and add to the shopping cart
   4. In the shopping cart interface, the user can change the quantity of the product and delete the product
   5. In the account management interface, users can set their own account 
   6. On the right side of the header of the page, you can log out of the current user login
   7. Click proceed to checkout on the shopping cart interface to enter the checkout interface. The checkout interface is divided into four parts, mailing address/payment method/order confirmation/order success interface

# Frontend Document

| User Document | Develope Document | Management Document |
| ------------- |:-----------------:|--------------------:|
| User Manual   | Develope Plan     | Test Plan(STP) |
|               |                   | Test Report(STR)|

# User Manual
## Using Ant Design Pro

This project is developed Amazon like shopping website. It's initialized with [Ant Design Pro](https://pro.ant.design). Follow is the quick guide for how to use.

### Environment Prepare

Install `node_modules`:

```bash
npm install
```

or

```bash
yarn
```

### Provided Scripts

Ant Design Pro provides some useful script to help you quick start and build with web project, code style check and test.

Scripts provided in `package.json`. It's safe to modify or add additional script:

#### Start project

```bash
npm start
```

#### Build project

```bash
npm run build
```

#### Check code style

```bash
npm run lint
```

You can also use script to auto fix some lint error:

```bash
npm run lint:fix
```

#### Test code

```bash
npm test
```

### More

You can view full document on our [official website](https://pro.ant.design). And welcome any feedback in our [github](https://github.com/ant-design/ant-design-pro).

## Visit Website

User should login or register account first. After this, user is able to visit our welcome page. We privide pages for this web application. 

### Welcome Page

We provide a dashboard for commodities overview and a search funtion that could be used to find items the user wants by key word.

### Cart List

User could check his/her items in the cart, change the number of the items, delete the items and checkout here.

### Checkout Page

User could only go to this page by clicking the button "Proceed to checkout" in the cart list page. User should input information needed: shipping address, payment method in order to checkout.

Shipping address include information: country, full name, phone number, address, city, state and zip code. All this information are required except city. Phone number must be 11 numbers and zip code must be 5 numbers.

Payment method include information: card number, name, expiration date, CVV.
After inputting these information, user goes to the 3rd part for checking the order. Once the order is placed, he/she is not able to change the order.


# Setup Guide: AnupYadav121/freshers-bootcamp/ExerciseSolutions/Day4-5_Codebase


# 1. Download

Download repo to your local machine:
> git clone https://github.com/AnupYadav121/freshers-bootcamp.git


# 2. Go Packages
GORM
> github.com/jinzhu/gorm

MySQL Driver
> github.com/go-sql-driver/mysql

Gin Web Framework
> github.com/gin-gonic/gin

mockery
> github.com/vektra/mockery

mockgen
> github.com/golang/mock/mockgen

# 3. Run Unit Tests

	go test ./... -v -coverpkg=./...

# 4. Models

### Customer Model

````azure
{   
    "id" : 1,
    "name":"dhaked",
    "password":"4321"
}
````

### Retailer Model

````azure
{   
    "id" : 1,
    "name":"dhaked",
    "password":"4321"
}
````

### Product Model

````azure
{
    "id":4
    "product_name": "bottle",
    "price": 50,
    "quantity": 40,
    "retailer_id": 1
}
````
### Order Model

````azure
{   
    "id":2
    "customer_id": 1,
    "product_id": 2,
    "quantity": 1
}
````

# 5. Go Routes

### Authentication
Use Basic Auth as authentication header while making API calls
Ex credentials `"username": "password"`

# Product Routes (relates to Retailer)

POST: "product/:retailerID"
> creates and saves the products of a retailer

GET: "product/:retailerID/:id"
> returns out the product of given id from all orders of given retailer

GET: "products/:retailerID"
> returns the all product of given retailer ID

PUT: "product/:retailerID/:id"
> updates the given id product for the given id retailer

DELETE: "product/:retailerID/:id"
> delete the given id product for the given id retailer

POST: "retailer"
> does authentication of given retailer

POST: "retailer/:retailerID"
> removes authentication of given id retailer

# Order Routes (relates to Customer)

POST : "order/:customerID"
> creates order for given customer ID

POST : "orders/:customerID"
> creates multiple order for given customer ID

GET : "order/:customerID/:id"
> return the updated order with status for given customer id and given order id

POST : "order/:customerID/:id"
> saves the order status of given customer id and order id

GET : "orders/:customerID"
> returns the all order of given customerID

GET : "transactionOrders/:retailerID"
> return the all orders of given retailerID made product ids

POST : "user"
> saves and authenticates a customer details

DELETE : "user/:customerID"
> removes the authentication of given customerID
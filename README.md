# Test Aiforesee

## List API

** Get All **

Method : "GET"
Path : "/fuel"

** Get By ID **
Method : "GET"
Path : "/fuel/:id"

** Create **
Method : "POST"
Path : "/fuel"
Body : {
    "fuel_liter": 2,
    "fuel_type": "Premium 2",
    "fuel_price": 20000
}

** Update **
Method : "PUT"
Path : "/fuel"
Body : {
    "id": 41,
    "fuel_liter": 1,
    "fuel_type": "Premium 1xx",
    "fuel_price": 11000
}

** Delete **
Method : "DELETE"
Path : "/fuel"
Body : {
    "id": 41
}


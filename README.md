# Tax Calculator

Language: Golang

Database: mysql

## Database Design

Table: tax

![](https://raw.githubusercontent.com/egig/tax_calculator/master/diagram.png)


## Api Routes

### POST /tax_objects

Description: Create Tax Object

Example Request

```$xslt
POST /tax_objects HTTP/1.1
Host: localhost:8080
Content-Type: application/json

{
	"name":"Lucky Strike",
	"tax_code": 2,
	"price": 1000
}

```

Example Reponse
```$xslt
{
    "id": 1,
    "name": "Lucky Strike",
    "tax_code": 2,
    "price": 1000
}
```

### GET /bill

Description: Get tax list, tax subtotal, price subtoal and grand total

Example Request:

```
GET /bill HTTP/1.1
Host: localhost:8080
```

Example Response:

```$xslt
{
    "price_sub_total": 1000,
    "tax_sub_total": 30,
    "grand_total": 1030,
    "tax_list": [
        {
            "id": 1,
            "name": "Lucky Strike",
            "tax_code": 2,
            "price": 1000,
            "type_name": "Tobacco",
            "refundable": false,
            "tax": 30,
            "amount": 1030
        }
    ]
}
```

## Run

Follow this step to run this application.

1. Make sure you have Internet connection and docker installed

2. Clone this repo
    
    ```$xslt
    git clone https://github.com/egig/tax_calculator.git
    ```
    
3. Change Working directory

    ```$xslt
    cd tax_calculator
    ```
    
4. Run docker compose
    ```$xslt
    docker-compose up
    ```
    
5. Visit the route in browser or api tester (e.g Postman)

    ```$xslt
     http://localhost:8080/bill
    ```
    
    


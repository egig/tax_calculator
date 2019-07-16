# Tax Calculator

Language: Golang

Database: mysql

## Database Design

Table: tax

![](https://raw.githubusercontent.com/egig/tax_calculator/master/diagram.png)


## Api Routes

**POST /tax_objects**

**GET /bill**

Complete API doc: [https://taxcalculator.docs.apiary.io](https://taxcalculator.docs.apiary.io)

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

## Test
```$xslt
go test ./...
```

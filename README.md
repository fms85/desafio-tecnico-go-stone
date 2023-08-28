# desafio-tecnico-go-stone

This API allows you to perform fund transfers between internal accounts of a digital bank.

## Dependencies

To run the application, you need to have **Docker** installed on your system. The environment configuration can be found in the `docker/docker-compose.yml` file.

## Dependencies

To run the application, you need to have Docker installed on your system. The environment configuration can be found in the `docker/docker-compose.yml` file.

## Getting Started

To start the application, run the following command:

    docker-compose -f docker/docker-compose.yml up
    
This will build and launch the application in a Docker container.

## Routes

Here are the available routes and their descriptions:

### Create a New Account
Create a new account with provided details.

`POST /accounts`

    curl --location 'http://localhost:8080/accounts' \
    --header 'Content-Type: application/json' \
    --data '{
        "name": "Mineiro",
        "cpf": "31974408035",
        "secret": "123456"
    }' -i
    
### Get Account Balance
Retrieve the balance of a specific account.

`GET /accounts/id/balance`

    curl --location 'http://localhost:8080/accounts/1/balance' -i

### Get List of Accounts
Retrieve a list of all accounts.

`GET /accounts`

    curl --location 'http://localhost:8080/accounts' -i

### Get Token for Account
Authenticate and retrieve a token for the account.

`POST /login`

    curl --location 'http://localhost:8080/login' \
    --header 'Content-Type: application/json' \
    --data '{
        "cpf": "31974408035",
        "secret": "123456"
    }' -i

### Create a Transfer
Create a new transfer between accounts.

`POST /accounts`

    curl --location 'http://localhost:8080/transfers' \
    --header 'Content-Type: application/json' \
    --header 'Authorization: Bearer TOKEN' \
    --data '{
        "account_destination_id": 2,
        "amount": 10
    }' -i

### Get List of Transfers
Retrieve a list of transfers for an authenticated account.

`GET /transfers`

    curl --location 'http://localhost:8080/transfers' \
    --header 'Authorization: Bearer TOKEN' -i

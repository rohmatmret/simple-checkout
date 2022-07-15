# Store Simple Checkout 
## the system should have the following promotions:
 - Each sale of a MacBook Pro comes with a free Raspberry Pi B
 - Buy 3 Google Homes for the price of 2
 - Buying more than 3 Alexa Speakers will have a 10% discount on all Alexa speakers


## How to run
```sh
 $ go run .
 ```

## How to run with tests
```sh
 $ go test ./...
 ```

 ## How to Build Linux
```sh
$ GOOS=linux GOARCH=amd64  go build -ldflags='-s' -o=.simple-store
```

## Build Docker Image
```sh
    $ docker build -t store .
```


## Example  API CALL

```sh
curl --location --request POST 'localhost:8080/store' \
--header 'Content-Type: application/json' \
--data-raw '{
    "Items":[
        { "name":"Google Home"},
        { "name":"Google Home"},
        { "name":"Google Home"}
    ]
}'
```
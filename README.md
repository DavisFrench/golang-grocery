# golang grocery
Small golang api for a grocery store's inventory

## Run code
Clone the repo and cd into the root.
Run the api with:
`go run cmd/app/main.go`

Once the api is running you can try a few curl comands in a separate terminal

Available curl commands:
```
// returns an ok status json if api is running
curl localhost:8888/grocery/ping

// return all inventory
curl localhost:8888/grocery/produce

// get info of a produce by its produceCode
// x is an alphanumeric character that is case insensitive
curl localhost:8888/grocery/produce/xxxx-xxxx-xxxx-xxxx

// delete produce from the inventory by its produceCode
// x is an alphanumeric character that is case insensitive
curl -X DELETE localhost:8888/grocery/produce/xxxx-xxxx-xxxx-xxxx

// add produce via a json
// add_produce.json is an example of the format to use when creating a schema
curl -X POST -H "Content-Type: application/json" -d @add_produce.json localhost:8888/grocery/produce
```

## Run tests
From the root run `go test ./...`

## Future improvements
* More tests
* JSON schema for testing input
* Implement a real database

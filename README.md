# curator

backend server for gallery project

## Requirement

* golang installed
* add $GOPATH/bin to $PATH

```
export PATH=$PATH:$GOPATH/bin
```

## Usage

* install go-bindata

```
make setup
```

* run server

```
make run
```

* build static binary

```
make
```

## Controller

### Ping

GET

```
curl localhost:9527/ping
```

### GraphQL

#### Photo

POST

```
curl -X POST -H 'Content-Type: application/json' -d '{"query": "{ photos { id,contributor,urls } }"}' localhost:9527/graphql
```

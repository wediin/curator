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

### Upload

#### Post

Simple
```
curl -F file=@123.jpeg localhost:9527/upload
```

Full
```
curl 'http://localhost:9527/upload' \
-H 'Origin: http://localhost:9527' \
-H 'Accept-Encoding: gzip, deflate, br' \
-H 'Accept-Language: zh-TW,zh;q=0.9,en-US;q=0.8,en;q=0.7' \
-H 'User-Agent: Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Mobile Safari/537.36' \
-H 'Content-Type: multipart/form-data' \
-H 'Accept: */*' \
-H 'Referer: http://localhost:9527' \
-H 'Connection: keep-alive' \
-F file=@123.jpeg \
--compressed
```

### GraphQL

#### Photo

POST

```
curl -X POST -H 'Content-Type: application/json' -d '{"query": "{ photos { id,contributor,urls } }"}' localhost:9527/graphql
```

# curator

backend server for gallery project

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

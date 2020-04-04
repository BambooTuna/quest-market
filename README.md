# quest-market


## 動作確認

### 認証
```bash
$ export API_ENDPOINT=localhost:8080/v1
$ export API_ENDPOINT=https://market-114.appspot.com/v1

$ curl -X POST -H "Content-Type: application/json" -d '{"mail":"bambootuna@gmail.com","pass":"pass"}' ${API_ENDPOINT}/signup -i
$ curl -X POST -H "Content-Type: application/json" -d '{"mail":"bambootuna@gmail.com","pass":"pass"}' ${API_ENDPOINT}/signin -i
// HeadreName: Set-AuthorizationにSessionTokenがセットされている

$ export SESSION_TOKEN=[~~~]
$ curl -X GET -H "Authorization: $SESSION_TOKEN" ${API_ENDPOINT}/health -i
```

### 資金操作
```bash
$ curl -X GET -H "Authorization: $SESSION_TOKEN" ${API_ENDPOINT}/money -i
$ curl -X POST -H "Authorization: $SESSION_TOKEN" -H "Content-Type: application/json" -d '{"to":"2719b422-d8a0-37c7-b4f0-a1800421beda","amount":10}' ${API_ENDPOINT}/money -i
```

### 出品
```bash
$ curl -X GET ${API_ENDPOINT}/items -i
$ curl -X GET -H "Authorization: $SESSION_TOKEN" ${API_ENDPOINT}/items -i

$ curl -X GET -H "Authorization: $SESSION_TOKEN" ${API_ENDPOINT}/items/my -i

$ curl -X POST -H "Authorization: $SESSION_TOKEN" -H "Content-Type: application/json" -d '{"title":"title","detail":"detail","price":1000}' ${API_ENDPOINT}/item -i
$ export ITEM_ID=[~~~]


$ curl -X GET ${API_ENDPOINT}/item/${ITEM_ID} -i
$ curl -X GET -H "Authorization: $SESSION_TOKEN" ${API_ENDPOINT}/item/${ITEM_ID} -i

$ curl -X PUT -H "Authorization: $SESSION_TOKEN" "${API_ENDPOINT}/item/${ITEM_ID}/purchase" -i
$ curl -X PUT -H "Authorization: $SESSION_TOKEN" "${API_ENDPOINT}/item/${ITEM_ID}/payment" -i
$ curl -X PUT -H "Authorization: $SESSION_TOKEN" "${API_ENDPOINT}/item/${ITEM_ID}/receipt" -i
```


## ローカル環境構築
```bash
$ VUE_APP_SERVER_ENDPOINT=http://localhost:8080/v1 npm run serve -- --port 8080 --host 0.0.0.0
$ VUE_APP_SERVER_ENDPOINT=http://localhost:8080/v1 npm run build

$ go mod init
$ go get -v -t -d ./...
$ go run main.go
```

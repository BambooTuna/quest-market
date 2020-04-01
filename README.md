# quest-market


## 動作確認
```bash
$ export API_ENDPOINT=localhost:8080/v1
$ export API_ENDPOINT=https://market-114.appspot.com/v1

$ curl -X POST -H "Content-Type: application/json" -d '{"mail":"bambootuna@gmail.com","pass":"pass"}' ${API_ENDPOINT}/signup -i
$ curl -X POST -H "Content-Type: application/json" -d '{"mail":"bambootuna@gmail.com","pass":"pass"}' ${API_ENDPOINT}/signin -i
// HeadreName: Set-AuthorizationにSessionTokenがセットされている

$ export SESSION_TOKEN=[~~~]
$ curl -X GET -H "Authorization: $SESSION_TOKEN" ${API_ENDPOINT}/health -i
```

```bash
$ curl -X GET -H "Authorization: $SESSION_TOKEN" ${API_ENDPOINT}/money -i
$ curl -X POST -H "Authorization: $SESSION_TOKEN" -H "Content-Type: application/json" -d '{"to":"2719b422-d8a0-37c7-b4f0-a1800421beda","amount":10}' ${API_ENDPOINT}/money -i

$ curl -X GET ${API_ENDPOINT}/products -i



$ export PRODUCT_ID=[~~~]
$ curl -X GET -H "Authorization: $SESSION_TOKEN" ${API_ENDPOINT}/purchase -i

$ curl -X PUT -H "Authorization: $SESSION_TOKEN" "${API_ENDPOINT}/purchase/${PRODUCT_ID}?type=waiting_for_payment" -i
$ curl -X PUT -H "Authorization: $SESSION_TOKEN" "${API_ENDPOINT}/purchase/${PRODUCT_ID}?type=waiting_to_receive" -i
```


## ローカル環境構築
```bash
$ VUE_APP_SERVER_ENDPOINT=http://localhost:8080/v1 npm run build

$ go mod init
$ go get -v -t -d ./...
$ go run main.go
```

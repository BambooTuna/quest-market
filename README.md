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


## ローカル環境構築
```bash
$ VUE_APP_SERVER_ENDPOINT=http://localhost:8080/v1 npm run build

$ go mod init
$ go get -v -t -d ./...
$ go run main.go
```

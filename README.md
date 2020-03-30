# quest-market


## 動作確認
```bash
$ curl -X POST -H "Content-Type: application/json" -d '{"mail":"bambootuna@gmail.com","pass":"pass"}' localhost:8080/signup -i
$ curl -X POST -H "Content-Type: application/json" -d '{"mail":"bambootuna@gmail.com","pass":"pass"}' localhost:8080/signin -i
// HeadreName: Set-AuthorizationにSessionTokenがセットされている


$ export SESSION_TOKEN=[~~~]
$ curl -X GET -H "Authorization: $SESSION_TOKEN" localhost:8080/health -i
```


## ローカル環境構築
```bash
$ VUE_APP_SERVER_ENDPOINT=http://localhost:8080 npm run build

$ go mod init
$ go get -v -t -d ./...
$ go run main.go
```

# quest-market


## 動作確認
```bash
$ curl -X POST -H "Content-Type: application/json" -d '{"mail":"bambootuna@gmail.com","pass":"pass"}' localhost:8080/signup -i
$ curl -X POST -H "Content-Type: application/json" -d '{"mail":"bambootuna@gmail.com","pass":"pass"}' localhost:8080/signin -i


$ export SESSION_TOKEN=[~~~]
$ curl -X GET -H "Authorization: $SESSION_TOKEN" localhost:8080/health -i
```

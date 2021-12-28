# ichinen_bingo

ホットリロード起動
```
$ air
```

リリースコマンド
```
$ heroku login
$ heroku container:login
$ heroku container:push web
$ heroku container:release web
```


テスト
```
$ go get github.com/joho/godotenv
$ godotenv -f .env go test -v ./...
```

ログの出力
```
$ heroku logs --tail
```
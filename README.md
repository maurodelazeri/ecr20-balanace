simple go program to read erc-20 balance of a given account.

to run:
```
$ go get
```

```
$ go build
```

```
$ go run *.go

   ____    __
  / __/___/ /  ___
 / _// __/ _ \/ _ \
/___/\__/_//_/\___/ v4.6.3
High performance, minimalist Go web framework
https://echo.labstack.com
____________________________________O/_______
                                    O\
⇨ http server started on [::]:1323
```

```json
$ curl -s "http://localhost:1323/?contract=0x3BA4c387f786bFEE076A58914F5Bd38d668B42c3&addr=0x5c4a4e39542bedf97e827f4601e66706cfc5b662" | jq
{
  "UserAddr": "0x5C4A4e39542BeDF97E827f4601e66706CFc5b662",
  "Balance": "3.2",
  "Contract": "0x3BA4c387f786bFEE076A58914F5Bd38d668B42c3",
  "Error": ""
}
```

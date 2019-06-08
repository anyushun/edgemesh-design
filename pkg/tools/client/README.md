# usage
```
cd $GOPATH/src/edgemesh

go build -o client tools/client/start.go

./client --ca=conf/certs/ca.crt --cert=conf/certs/client.crt --key=conf/certs/client.key --method=GET --url=https://localhost:8443/hello --body='{"name": "anyushun"}'
```

# go | mmock

## Source Code

- <https://github.com/jmartin82/mmock>

## Steps to run MMOCK

```bash
go build cmd/mmock/main.go
```

```bash
mv main mmock
```

```bash
nohup ./mmock -server-port=7200 -server-tls-port=7201 -console-port=7202 -server-statistics=false > mmock_$(date +%Y%m%d%H%M%S).log &
```
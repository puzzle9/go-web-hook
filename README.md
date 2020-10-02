# 一个简单的 web hook 处理

## 编译
### linux amd64
```sh
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o go-web-hook src/main.go
```

### linux arm64
```sh
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o go-web-hook src/main.go
```

## 运行

### sh
```sh
./go-web-hook -config /home/puzzle/web/go-web-hook/config.json &
```

### systemd
`/etc/systemd/system/gowebhook.service`

```
[Unit]
Description=go-web-hook service
After=network.target

[Service]
Type=simple
ExecStart=/root/go-web-hook/go-web-hook -config /root/go-web-hook/config.json

[Install]
WantedBy=multi-user.target
```

## 配置文件
```json
{
  "host": "0.0.0.0",
  "port": "8080",
  "script": "/home/puzzle/web/go-web-hook/script.sh"
}
```

## web hook url
- http://0.0.0.0:8080/webhook

# todo
- [x] 先挖一个坑
- [ ] 支持 win 服务器
- [ ] 秘钥验证
- [ ] web 管理

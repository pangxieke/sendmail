# 使用Go发送邮件
Go结合Redis队列发送邮件。其他服务写入发送数据到redis队列。
Go服务轮询队列，发送邮件。

## 使用
```
go run cmd/main.go
make test

```

## Redis格式

添加redis队列，`key`为`mail_notify`
格式如下
```
lpush mail_notify "{\"receivers\": \"pangxieke@126.com, pangxieke@126.com\", \"subject\": \"this is subject\", \"body\": \"this is body\", \"subtype\": \"html\"}"

lrange mail_notify 0 -1
```

Go struct 如下
```$xslt
type MailNotify struct {
	Receivers string `json:"receivers"`
	Subject   string `json:"subject"`
	Body      string `json:"body"`
	Subtype   string `json:"subtype"`
}
```

## 注意

docker-compose 配置smtp.exmail.qq.com:163.177.90.125，设置/etc/hosts, 否则可能出现连接超时
package model

import (
	"fmt"
	"strings"

	"github.com/go-gomail/gomail"
	"github.com/pangxieke/sendmail/config"
	"github.com/pangxieke/sendmail/log"
)

type MailNotify struct {
	Receivers string `json:"receivers"`
	Subject   string `json:"subject"`
	Body      string `json:"body"`
	Subtype   string `json:"subtype"`
}

func SendMail(data MailNotify) {
	m := gomail.NewMessage()
	//发件人
	m.SetHeader("From", config.Mail.Sender)
	//收件人
	fmt.Println(data.Receivers)
	mailTo := strings.Split(data.Receivers, ",")
	fmt.Println(mailTo)
	//mailTo := []string{data.Receivers}
	m.SetHeader("To", mailTo...)
	//抄送人
	//m.SetAddressHeader("Cc", "**", "test")
	//邮件标题
	m.SetHeader("Subject", data.Subject)
	//邮件内容
	m.SetBody("text/html", data.Body)
	//邮件附件
	// 	m.Attach("E:\\IMGP0814.JPG")

	d := gomail.NewDialer(config.Mail.Host, int(config.Mail.Port), config.Mail.Sender, config.Mail.PWD)
	//邮件发送服务器信息,使用授权码而非密码
	if err := d.DialAndSend(m); err != nil {
		log.Info("send err, data:", data, "err:", err)
	} else {
		log.Info("send success, data:", data)
	}

}

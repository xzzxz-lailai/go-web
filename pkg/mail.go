package pkg

import (
	"github.com/go-gomail/gomail"
	"go-web/config"
)

// SendVerifyCode 发送邮件（给指定邮箱发送一封“验证码邮件”）
func SendVerifyCode(toEmail, code string) error {
	// 创建一封邮件对象
	m := gomail.NewMessage()
	m.SetAddressHeader( // 自动帮你处理中文编码（UTF-8）, 名字 + 邮箱的 RFC 标准格式 , 各家 SMTP 的兼容问题
		"From",
		config.Cfg.SMTP.From,      // 邮箱地址  发件人邮箱地址
		config.Cfg.SMTP.From_Name, // 显示名称
	)
	m.SetHeader("To", toEmail)                        // 设置收件人
	m.SetHeader("Subject", "邮箱验证码")                   // 邮件标题
	m.SetBody("text/plain", "你的验证码是："+code+"，5分钟内有效") // 邮件正文

	// 创建 SMTP 连接器
	d := gomail.NewDialer(
		config.Cfg.SMTP.Host, // smtp.qq.com
		config.Cfg.SMTP.Port,
		config.Cfg.SMTP.User, // 发件人邮箱地址
		config.Cfg.SMTP.Pass,
	)
	// 开启 SSL
	d.SSL = true
	// 发送邮件
	return d.DialAndSend(m)

}

package email_service

import (
	"blogW_server/global"
	"crypto/tls"
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
)

func SendRegisterCode(to string, code string) error {
	em := global.Config.Email
	subject := fmt.Sprintf("【%s】账号注册", em.SendNickname)
	text := fmt.Sprintf("你正在进行账号注册操作，这是你的验证码 %s，十分钟内有效", code)
	return SendEmail(to, subject, text)
} // 注册账号邮件

func SendResetPwdCode(to string, code string) error {
	em := global.Config.Email
	subject := fmt.Sprintf("【%s】重置密码", em.SendNickname)
	text := fmt.Sprintf("你正在进行重置密码操作，这是你的验证码 %s，十分钟内有效", code)
	return SendEmail(to, subject, text)
} // 重置密码邮件

func SendBindEmailCode(to string, code string) error {
	em := global.Config.Email
	subject := fmt.Sprintf("【%s】重置密码", em.SendNickname)
	text := fmt.Sprintf("你正在进行邮箱绑定操作，这是你的验证码 %s，十分钟内有效", code)
	return SendEmail(to, subject, text)
} // 绑定邮箱

func SendEmail(to, subject, text string) (err error) {
	em := global.Config.Email
	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", em.SendNickname, em.SendEmail)
	e.To = []string{to}
	e.Subject = subject
	e.Text = []byte(text)

	// 使用 SSL/TLS 加密发送（465端口）
	err1 := e.SendWithTLS(
		fmt.Sprintf("%s:%s", em.Domain, em.Port), // smtp.163.com:465
		smtp.PlainAuth("", em.SendEmail, em.AuthCode, em.Domain),
		&tls.Config{
			ServerName:         em.Domain, // smtp.163.com
			InsecureSkipVerify: false,     // 生产环境应为false（测试时可临时设为true）
		},
	)

	if err1 != nil {
		fmt.Println("邮件发送失败:", err1)
		return err1
	}
	fmt.Println("发送成功")
	return nil
}

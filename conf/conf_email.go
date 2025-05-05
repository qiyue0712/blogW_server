package conf

type Email struct {
	Domain       string `yaml:"domain" json:"domain"`
	Port         string `yaml:"port" json:"port"`
	SendEmail    string `yaml:"sendEmail" json:"sendEmail"`
	AuthCode     string `yaml:"authCode" json:"authCode"` // 授权码
	SendNickname string `yaml:"sendNickname" json:"sendNickname"`
	SSL          bool   `yaml:"ssl" json:"SSL"`
	TLS          bool   `yaml:"tls" json:"TLS"`
}

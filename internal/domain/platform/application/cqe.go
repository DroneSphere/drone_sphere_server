package application

// InfoResult 平台信息查询结果
type InfoResult struct {
	Platform  string `json:"platform"`
	Workspace string `json:"workspace"`
	Desc      string `json:"desc"`
}

// MQTTParam MQTT连接参数
type MQTTParam struct {
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// HTTPParam HTTP连接参数
type HTTPParam struct {
	Host  string `json:"host"`
	Token string `json:"token"`
}

package user_app

import (
	platform_app "drone_sphere_server/internal/domain/platform/app"
	"drone_sphere_server/internal/domain/user"
)

type LoginCommand struct {
	Username string `validate:"required" json:"username"`
	Password string `validate:"required" json:"password"`
	SN       string `json:"sn"`
}

type LoginSuccessEvent struct {
	UserID string
	User   *user.User
	SN     string
}

type LoginResult struct {
	User   *user.User              `json:"user"`
	Token  string                  `json:"token"`
	Info   platform_app.InfoResult `json:"info"`
	Params struct {
		MQTT platform_app.MQTTParam `json:"mqtt"`
		HTTP platform_app.HTTPParam `json:"http"`
	} `json:"params"`
}

type RegisterCommand struct {
	Username string `validate:"required" json:"username"`
	Password string `validate:"required" json:"password"`
}

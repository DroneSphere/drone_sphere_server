package application

import (
	platformapp "drone_sphere_server/internal/domain/platform/application"
	"drone_sphere_server/internal/domain/user/entity"
)

type LoginCommand struct {
	Username string `validate:"required" json:"username"`
	Password string `validate:"required" json:"password"`
	SN       string `json:"sn"`
}

type LoginSuccessEvent struct {
	UserID string
	User   *entity.User
	SN     string
}

type LoginResult struct {
	User   *entity.User           `json:"user"`
	Token  string                 `json:"token"`
	Info   platformapp.InfoResult `json:"info"`
	Params struct {
		MQTT platformapp.MQTTParam `json:"mqtt"`
		HTTP platformapp.HTTPParam `json:"http"`
	} `json:"params"`
}

type RegisterCommand struct {
	Username string `validate:"required" json:"username"`
	Password string `validate:"required" json:"password"`
}

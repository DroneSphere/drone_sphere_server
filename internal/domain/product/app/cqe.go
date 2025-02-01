package product_app

import (
	"drone_sphere_server/internal/domain/common"
	entity "drone_sphere_server/internal/domain/product/entity"
)

// ConnectRCCommand 连接遥控器命令, 遥控器验证证书后调用
type ConnectRCCommand struct {
	// SN 遥控器序列号
	SN string `json:"sn"`
}

// UpdateTopoCommand 更新拓扑命令, 在遥控器子设备拓扑更新后触发 MQTT 消息
type UpdateTopoCommand struct {
	ProductTopo
	SubDevices []struct {
		ProductTopo
		Index int `json:"index"`
	} `json:"sub_devices"`
}

// ProductTopo 产品拓扑
type ProductTopo struct {
	Domain       entity.ProductDomain `json:"domain"`
	Type         entity.Type          `json:"type"`
	SubType      entity.SubType       `json:"sub_type"`
	DeviceSecret string               `json:"device_secret"`
	Nonce        string               `json:"nonce"`
	ThingVersion string               `json:"thing_version"`
}

type UpdateTopoEvent struct {
	common.CommonModel
	Data UpdateTopoCommand `json:"data"`
}

type UpdateTopoReplyEvent struct {
	common.CommonModel
	Data struct {
		Result int `json:"result"`
	}
}

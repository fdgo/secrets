package user

type Login struct {
	Token     string `json:"token"`
	UseradvId int64  `json:"userId"`
	DeviceId  int64  `json:"deviceId"`
}

package model

type UserOrderCommandBody struct {
	Token   string `json:"token"`
	Thingid string `json:"thingid"`
	Command string `json:"command"`
}

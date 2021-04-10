package service

import (
	"encoding/json"
	"fmt"

	"github.com/labstack/echo"
	"github.com/thanhlam/user-control-service/model"
)

func UserOrderCommand(c echo.Context) error {
	userOrderCommandBody := new(model.UserOrderCommandBody)
	err := c.Bind(userOrderCommandBody)
	if err != nil {
		fmt.Println(err)
		return c.JSON(400, map[string]interface{}{"code": "6", "message": "Body is Invalid", "data": nil})
	}
	token := userOrderCommandBody.Token
	thingID := userOrderCommandBody.Thingid
	command := userOrderCommandBody.Command
	//fmt.Println(token)
	//fmt.Println(thingID)
	//fmt.Println(command)
	//check user role and command
	result := CheckUserRole(token, thingID, command)
	if result != true {
		return c.JSON(400, map[string]interface{}{"code": "3", "message": "Not Allow", "data": nil})
	}
	//check user role and command
	//push message to kafka
	ProducerMessage(command, "controlTopic")
	//push message to kafka
	return c.JSON(400, map[string]interface{}{"code": "0", "message": "Success", "data": nil})
}

//get user role
//get thing role by thingid
type GetUserRoleBody struct {
	Token   string `json:"token"`
	Thingid string `json:"thingid"`
}
type RespUserRole struct {
	Data []string
}

func CheckUserRole(token, thingID, command string) bool {
	url := "http://54.254.177.239:1323/api/user/userPushCommand"
	var getUserRoleBody GetUserRoleBody
	getUserRoleBody.Token = token
	getUserRoleBody.Thingid = thingID
	//convert struct to string
	e, err := json.Marshal(getUserRoleBody)
	if err != nil {
		fmt.Println(err)
		return false
	}
	requestTokenStr := string(e)
	//fmt.Println(requestTokenStr)
	//convert struct to string
	resp := Post(url, requestTokenStr)
	//fmt.Println(resp)
	var respUserRole RespUserRole
	json.Unmarshal([]byte(resp), &respUserRole)
	checkResp := contains(respUserRole.Data, command)
	return checkResp

}

//check element exits in array
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

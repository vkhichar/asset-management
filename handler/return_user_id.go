package handler

import (
	"fmt"
	"strconv"
	"strings"
)

var UseId string

func ReadUserId() {

	tokenParts := strings.Split(Token, ";")

	userIDParts := strings.Split(tokenParts[0], ":")

	userId, err := strconv.Atoi(userIDParts[1])
	if err != nil {
		fmt.Printf("handler: error in return userid converting userid in string: %s", err.Error())
	}
	UseId = strconv.Itoa(userId)
}

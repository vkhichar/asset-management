package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vkhichar/asset-management/contract"
	"github.com/vkhichar/asset-management/service"
)

func ListUsersHandler(userService service.UserService) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type","application/json")

		if(r.Method == "GET"){
			user,err := userService.ListUser(r.Context())

			if err!=nil{
				fmt.Printf(err)
				return 
			}
			//write a loop to convert domain object to contract object

			w.WriteHeader(http.StatusOK)
			responsebytes,_:= json.Marshal(contract.ListUserResponse{user.})
			w.Write(responsebytes)
		}
	}
}
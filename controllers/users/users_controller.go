package users

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/smbirch/bookstore_users-api/domain/users"
)

func CreateUser(c *gin.Context) {
	var user users.User
	fmt.Println("user ID:", user)
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println("err:", err.Error())
		//todo: Handle error!
		return
	}
	if err := json.Unmarshal(bytes, &user); err != nil {
		fmt.Println("err:", err.Error())
		//todo: handle json error
		return
	}

	fmt.Println(string(bytes))
	c.String(http.StatusNotImplemented, "implement me")
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")

}

package helper

import (
	"io/ioutil"
	"log"
	"encoding/json"
	"github.com/gin-gonic/gin"
)


func  ValidUser() gin.Accounts {

	var users gin.Accounts

	// Read file
	users_list, err := ioutil.ReadFile("./helper/users.json")

	// Validate that there is no error when reading the file
	if err != nil {
		log.Fatal(err)
	}

	// JSON is converted
	json.Unmarshal(users_list, &users)

	return users
}

package auth

import (
	"fmt"
	"github.com/blhack/goGoApi/database"
	"github.com/blhack/goGoApi/users"
	"github.com/blhack/goGoApi/utils"
	"golang.org/x/crypto/bcrypt"
)

func Register(username string, pass1 string, pass2 string) utils.Success {

	this := utils.Success{false, "null"}

	if pass1 == pass2 {
		userStatus := users.CheckUser(username)
		fmt.Println("Does user exist?")
		fmt.Println(userStatus.Success)
		if userStatus.Success == false {
			password := []byte(pass1)
			hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)

			utils.CheckErr(err)

			uuid := utils.GenerateUuid()

			stmt, err := database.DBCon.Prepare("insert into users(username,password,uuid) values(?,?,?)")
			utils.CheckErr(err)

			_, err = stmt.Exec(username, hashedPassword, uuid)
			utils.CheckErr(err)
			this.Success = true
			this.Status = "user created"
		} else {
			this.Success = false
			this.Status = "user exists"
		}
	}
	return (this)
}

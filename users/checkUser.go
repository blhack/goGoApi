package users

import (
	"fmt"
	"github.com/blhack/goGoApi/database"
	"github.com/blhack/goGoApi/utils"
)

func CheckUser(username string) utils.Success {
	fmt.Println("Checking for user:", username)
	rows, err := database.DBCon.Query("select count(*) from users where username = ?", username)
	utils.CheckErr(err)

	this := utils.Success{false, "null"}
	count := 0
	for rows.Next() {
		err = rows.Scan(&count)
		utils.CheckErr(err)
	}

	if count > 0 {
		fmt.Printf("%v exists\n", username)
		this.Success = true
		this.Status = "user exists"
	} else {
		fmt.Printf("%v does not exist\n", username)
	}

	return (this)
}

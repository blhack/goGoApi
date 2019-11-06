package auth

import (
	"github.com/blhack/goGoApi/database"
	"github.com/blhack/goGoApi/utils"
	"golang.org/x/crypto/bcrypt"
)

func Auth(username string, password string) utils.Success {
	hashedPassword := []byte("")
	bytedPassword := []byte(password)
	this := utils.Success{false, "invalid username/password combination"}
	rows, err := database.DBCon.Query("select password from users where username = ?", username)
	utils.CheckErr(err)

	for rows.Next() {
		err = rows.Scan(&hashedPassword)
		utils.CheckErr(err)
	}

	err = bcrypt.CompareHashAndPassword(hashedPassword, bytedPassword)
	if err == nil {
		this.Success = true
		this.Status = "user authenticated"
	}

	return (this)
}

func DeAuth(sessionId string) utils.Success {
	stmt, err := database.DBCon.Prepare("delete from sessions where sessionId = ? limit 1")
	utils.CheckErr(err)
	_, err = stmt.Exec(sessionId)
	utils.CheckErr(err)
	success := utils.Success{true, "session ended"}
	return (success)
}

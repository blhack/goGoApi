package auth

import (
	"fmt"
	"github.com/blhack/goGoApi/database"
	"github.com/blhack/goGoApi/utils"
	"net/http"
)

func WhoAmI(r *http.Request) utils.Success {
	var userName string

	this := utils.Success{false, ""}

	sessionId, err := r.Cookie("sessionId")
	fmt.Println(r.Header.Get("Origin"))

	utils.CheckErr(err)

	if err == nil {
		rows, err := database.DBCon.Query("select username from sessions where sessionId = ?", sessionId.Value)
		utils.CheckErr(err)

		for rows.Next() {
			err = rows.Scan(&userName)
			this.Success = true
			utils.CheckErr(err)
		}
		this.Status = userName

	}

	return this
}

package things

import (
	"github.com/blhack/goGoApi/database"
	"github.com/blhack/goGoApi/utils"
	"time"
)

func AddThing(title string, url string, text string, username string, fileUuid string) utils.Success {

	uuid := utils.GenerateUuid()
	this := utils.Success{false, "failed to add comment"}


	if (title != "") {

		stmt, err := database.DBCon.Prepare("INSERT INTO things(uuid,username,title,url,text,creationDate,fileUuid) values(?,?,?,?,?,?,?)")
		utils.CheckErr(err)

		_, err = stmt.Exec(uuid, username, title, url, text, time.Now(), fileUuid)
		utils.CheckErr(err)

		this.Success = true
		this.Status = uuid
	}
	return this
}

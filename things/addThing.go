package things

import (
	"github.com/blhack/flexIms-go/database"
	"github.com/blhack/flexIms-go/utils"
	"time"
	"fmt"
)

func AddThing(title string, username string, fileUuid string, attributes []Attribute) utils.Success {

	uuid := utils.GenerateUuid()
	this := utils.Success{false, "failed to add comment"}
	fmt.Println("Adding an element.  Here are the attributes:")
	for _, element := range attributes {
		fmt.Println(element.AttributeName)
		fmt.Println(element.AttributeValue)
		stmt, err := database.DBCon.Prepare("INSERT INTO attributes(username,uuid,attributeName,attributeValue) values(?,?,?,?)")
		utils.CheckErr(err)
		_, err = stmt.Exec(username,uuid,element.AttributeName,element.AttributeValue)
		this.Success = true
		this.Status = uuid
	}

	if (title != "") {

		stmt, err := database.DBCon.Prepare("INSERT INTO attributes(username,uuid,attributeName,attributeValue) values(?,?,?,?)")
		utils.CheckErr(err)
		_, err = stmt.Exec(username,uuid,"title",title)


		stmt, err = database.DBCon.Prepare("INSERT INTO things(uuid,username,title,creationDate,fileUuid) values(?,?,?,?,?)")
		utils.CheckErr(err)

		_, err = stmt.Exec(uuid, username, title, time.Now(), fileUuid)
		utils.CheckErr(err)

		this.Success = true
		this.Status = uuid
	}
	return this
}

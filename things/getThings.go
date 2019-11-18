package things

import (
	"fmt"
	"github.com/blhack/goGoApi/database"
	"github.com/blhack/goGoApi/utils"
	"database/sql"
)

func GetThings(userName string) []Thing {

	var things []Thing

	rows, err := database.DBCon.Query("select * from things where username = ?", userName)
	utils.CheckErr(err)
	for rows.Next() {
		var thing Thing
		var id int
		var title string
		var url sql.NullString
		var text sql.NullString
		var _username string
		var creationDate string
		var uuid string
		var fileUuid string
		err = rows.Scan(&id, &title, &url, &text, &_username, &creationDate, &uuid, &fileUuid)
		utils.CheckErr(err)
		thing.Id = id
		thing.Title = title
		thing.Url = url.String
		thing.Text = text.String
		thing.Username = _username
		thing.CreationDate = creationDate
		thing.Uuid = uuid
		thing.FileUuid = fileUuid

		if fileUuid != "" {
			thing.FilePath = fmt.Sprintf("/uploads/photos/%s.jpg", fileUuid)
		} else {
			thing.FilePath = ""
		}
		fmt.Println("UUID")
		fmt.Println(uuid)
		attrs, err := database.DBCon.Query("select attributeName,attributeValue from attributes where uuid = ?", uuid)
		utils.CheckErr(err)
		for attrs.Next() {
			var _Attribute Attribute
			var attributeName sql.NullString
			var attributeValue sql.NullString
			err = attrs.Scan(&attributeName,&attributeValue)
			utils.CheckErr(err)
			_Attribute.AttributeName = attributeName.String
			_Attribute.AttributeValue = attributeValue.String
			fmt.Println(attributeName)
			thing.Attributes = append(thing.Attributes, _Attribute)
		}

		things = append(things, thing)
	}

	return things
}

package things

import (
	"fmt"
	"github.com/blhack/goGoApi/database"
	"github.com/blhack/goGoApi/utils"
)

func GetThings(listId string) []Thing {

	var things []Thing

	rows, err := database.DBCon.Query("select * from things where listId = ?", listId)
	utils.CheckErr(err)
	for rows.Next() {
		var thing Thing
		var id int
		var title string
		var url string
		var text string
		var _username string
		var creationDate string
		var _listId string
		var uuid string
		var fileUuid string
		err = rows.Scan(&id, &title, &url, &text, &_username, &creationDate, &_listId, &uuid, &fileUuid)
		utils.CheckErr(err)
		thing.Id = id
		thing.Title = title
		thing.Url = url
		thing.Text = text
		thing.Username = _username
		thing.CreationDate = creationDate
		thing.ListId = _listId
		thing.Uuid = uuid
		thing.FileUuid = fileUuid

		if fileUuid != "" {
			thing.FilePath = fmt.Sprintf("/uploads/photos/%s.jpg", fileUuid)
		} else {
			thing.FilePath = ""
		}

		things = append(things, thing)
	}

	return things
}

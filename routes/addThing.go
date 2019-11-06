package routes

import (
	"encoding/json"
	"fmt"
	"github.com/blhack/goGoApi/auth"
	"github.com/blhack/goGoApi/things"
	"github.com/blhack/goGoApi/utils"
	"io"
	"net/http"
	"os"
)

func AddThing(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	userName := auth.WhoAmI(r).Status

	this := utils.Success{false, ""}

	if auth.WhoAmI(r).Success == true {
		r.Body = http.MaxBytesReader(w, r.Body, 1024*1024*1024)
		r.ParseMultipartForm(500 << 20)
		file, handler, err := r.FormFile("file")
		fileUuid := ""
		if err == nil {
			fmt.Println("There was a FILE FLIGHT!")
			fileName := handler.Filename
			fmt.Printf("Filename: %s\n", fileName)
			fileUuid = utils.GenerateUuid()
			targetName := "./static/uploads/photos/" + fileUuid + ".jpg"
			f, err := os.OpenFile(targetName, os.O_WRONLY|os.O_CREATE, 0666)
			utils.CheckErr(err)
			defer f.Close()
			io.Copy(f, file)
			f.Close()

		} else {
			fmt.Println("There was decidedly not a file flight")
			fmt.Printf("Error: %s\n", err)
		}

		title := r.FormValue("title")
		url := r.FormValue("url")
		text := r.FormValue("text")

		fmt.Printf("You %v told me to add the following title %v", userName, title)

		this = things.AddThing(title, url, text, userName, fileUuid)
	} else {
		this.Success = false
		this.Status = "not authenticated"
	}
	_this, _ := json.Marshal(this)
	fmt.Fprintf(w, string(_this))
}

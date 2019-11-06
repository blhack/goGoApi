package routes

import (
	"encoding/json"
	"fmt"
	"github.com/blhack/goGoApi/auth"
	"github.com/blhack/goGoApi/utils"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	username := r.FormValue("username")
	pass1 := r.FormValue("pass1")
	pass2 := r.FormValue("pass2")

	this := auth.Register(username, pass1, pass2)

	registerSuccess, err := json.Marshal(this)
	utils.CheckErr(err)

	fmt.Fprintf(w, string(registerSuccess))

}

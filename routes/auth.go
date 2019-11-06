package routes

import (
	"encoding/json"
	"fmt"
	"github.com/blhack/goGoApi/auth"
	"github.com/blhack/goGoApi/config"
	"github.com/blhack/goGoApi/database"
	"github.com/blhack/goGoApi/utils"
	"net/http"
	"time"
)

func Auth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")

	username := r.FormValue("username")
	password := r.FormValue("password")

	this := auth.Auth(username, password)

	if this.Success == true {
		sessionId, err := auth.GenerateRandomString(64)
		utils.CheckErr(err)
		expiration := time.Now().Add(365 * 24 * time.Hour)
		stmt, err := database.DBCon.Prepare("insert into sessions(username,sessionId,time) values(?,?,?)")
		utils.CheckErr(err)
		_, err = stmt.Exec(username, sessionId, time.Now())
		utils.CheckErr(err)
		cookie := http.Cookie{Name: "sessionId", Value: sessionId, Expires: expiration, Domain: config.CookieDomain, Path: "/"}
		http.SetCookie(w, &cookie)
	}

	authSuccess, err := json.Marshal(this)
	utils.CheckErr(err)
	fmt.Fprintf(w, string(authSuccess))

}

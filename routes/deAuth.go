package routes

import (
	"encoding/json"
	"fmt"
	"github.com/blhack/goGoApi/auth"
	"github.com/blhack/goGoApi/utils"
	"net/http"
)

func DeAuth(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", 302)
	success := utils.Success{false, ""}

	sessionIdCookie, err := r.Cookie("sessionId")
	if err == nil {
		sessionId := sessionIdCookie.Value
		success = auth.DeAuth(sessionId)

	}

	successJson, err := json.Marshal(success)
	utils.CheckErr(err)
	fmt.Fprintf(w, string(successJson))
}
